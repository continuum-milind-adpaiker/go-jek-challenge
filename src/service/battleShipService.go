package service

import (
	"fmt"
	"log"

	model "github.com/ContinuumLLC/go-jek/src/model"
)

// BattleShipServiceFactoryImpl ...
type BattleShipServiceFactoryImpl struct{}

// battleShipServiceImpl ...
type battleShipServiceImpl struct {
	dependency model.BattleShipSvcDep
	inputPath  string
}

// GetBattleShipService ...
func (BattleShipServiceFactoryImpl) GetBattleShipService(dep model.BattleShipSvcDep, inPath string) model.BattleShipService {
	return &battleShipServiceImpl{
		dependency: dep,
		inputPath:  inPath,
	}
}

func (bs battleShipServiceImpl) Start() {
	readerDal := bs.dependency.GetInputReaderDal(bs.inputPath)
	inputData, err := readerDal.ReadConfigFile()
	if err != nil {
		log.Fatal(err)
	}
	bs.play(inputData)
	bs.saveResultToFile(inputData)

}

func (bs battleShipServiceImpl) saveResultToFile(resultData *model.BattleShipConfig) {
	writerDal := bs.dependency.GetOutputSaverDal()
	err := writerDal.SaveOutput(resultData)
	if err != nil {
		log.Fatal(err)
	}
}

func (bs battleShipServiceImpl) play(inputData *model.BattleShipConfig) string {
	//create channel of target coordinates
	//Each payer will target the opponent's ship
	target := make(chan model.Coordinates)
	playerAStatusCh := make(chan int)
	playerBStatusCh := make(chan int)

	go playerOne(inputData.Player1, inputData.TotMissiles, target, playerAStatusCh)
	go playerTwo(inputData.Player2, inputData.TotMissiles, target, playerBStatusCh)
	//<-done
	hitsByB := <-playerAStatusCh
	hitsByA := <-playerBStatusCh
	var result string
	if hitsByB == hitsByA {
		result = "Draw"
	} else if hitsByB > hitsByA {
		result = "B wins"
	} else {
		result = "A wins"
	}
	close(target)

	close(playerAStatusCh)
	close(playerBStatusCh)
	fmt.Printf("Final Score: PlayerA Hits: %d, PlayerB Hits %d\n", hitsByA, hitsByB)
	fmt.Println(result)
	fmt.Printf("Player A survived at length %d %v\n", len(inputData.Player1.MissedByOpponent), inputData.Player1.MissedByOpponent)
	fmt.Printf("Player B survived at %v\n", inputData.Player2.MissedByOpponent)
	return result
}

func playerOne(player model.Player, TotMissiles int, target chan model.Coordinates, playerAStatusCh chan<- int) {

	var pTwoHits int
	var hit bool
	missed := 0
	for i := 0; i < TotMissiles; i++ {
		//I fire first
		target <- player.EnemyTarget[i]

		//Get ready to receive enemy's Incoming missile
		missileIncoming := <-target
		for coordinate := range player.MyShipsPos {
			if missileIncoming.X == coordinate.X && missileIncoming.Y == coordinate.Y {
				player.MyShipsPos[coordinate] = false
				pTwoHits++
				hit = true
				break
			}
		}
		if hit == false {
			//fmt.Printf("index %d \n", missed)
			//player.MissedByOpponent[missed] = model.Coordinates{X: missileIncoming.X, Y: missileIncoming.Y}
			player.MissedByOpponent[model.Coordinates{X: missileIncoming.X, Y: missileIncoming.Y}] = true
			missed++
		}
		hit = false
	}

	//player 1 reports total hits by player 2
	playerAStatusCh <- pTwoHits
}

func playerTwo(player model.Player, TotMissiles int, target chan model.Coordinates, playerBStatusCh chan<- int) {
	var pOneHits int
	var hit bool
	missed := 0
	for i := 0; i < TotMissiles; i++ {
		//Receive enemy's incoming missile
		missileIncoming := <-target
		for coordinate := range player.MyShipsPos {
			if missileIncoming.X == coordinate.X && missileIncoming.Y == coordinate.Y {
				player.MyShipsPos[coordinate] = false
				pOneHits++
				hit = true
				break
			}
		}
		if hit == false {
			//player.MissedByOpponent[missed] = model.Coordinates{X: missileIncoming.X, Y: missileIncoming.Y}
			player.MissedByOpponent[model.Coordinates{X: missileIncoming.X, Y: missileIncoming.Y}] = true

			missed++
		}
		hit = false
		target <- player.EnemyTarget[i]
	}
	//player 2 reports total hits by player 1
	playerBStatusCh <- pOneHits
}
