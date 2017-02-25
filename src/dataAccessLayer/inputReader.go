package dal

import (
	"bufio"
	"os"

	"strconv"

	"strings"

	"errors"

	model "github.com/ContinuumLLC/go-jek/src/model"
)

// InputReaderFactoryImpl ...
type InputReaderFactoryImpl struct{}

type inputReaderDalImpl struct {
	inputPath string
}

// GetInputReaderDal ...
func (InputReaderFactoryImpl) GetInputReaderDal(input string) model.InputReaderDal {
	return &inputReaderDalImpl{
		inputPath: input,
	}
}

//ReadConfigFile ...
//Reading file into battleshipConfig object. Logic can be improved
func (cs inputReaderDalImpl) ReadConfigFile() (*model.BattleShipConfig, error) {

	file, err := os.Open(cs.inputPath)

	if err != nil {
		return &model.BattleShipConfig{}, err
	}

	defer file.Close()
	bcs := &model.BattleShipConfig{}
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	var a [7]string

	for i := 0; i < 7; i++ {
		scanner.Scan()
		a[i] = scanner.Text()
	}
	bcs.GridSize, err = strconv.Atoi(a[0])
	if err != nil {
		return &model.BattleShipConfig{}, err
	}
	numOfShips, err := strconv.Atoi(a[1])
	if err != nil {
		return &model.BattleShipConfig{}, err
	}
	bcs.Player1.MyShips = numOfShips
	bcs.Player2.MyShips = numOfShips
	bcs.TotMissiles, err = strconv.Atoi(a[4])
	if err != nil {
		return &model.BattleShipConfig{}, err
	}
	bcs.Player1.MyShipsPos = make(map[model.Coordinates]bool)
	ss := strings.Split(a[2], ":")

	for i := 0; i < bcs.Player1.MyShips; i++ {
		ssx := strings.Split(ss[i], ",")
		x, err := strconv.Atoi(ssx[0])
		if err != nil {
			return &model.BattleShipConfig{}, err
		}
		y, err := strconv.Atoi(ssx[1])
		if err != nil {
			return &model.BattleShipConfig{}, err
		}
		if x >= bcs.Player1.MyShips || y >= bcs.Player1.MyShips {
			return &model.BattleShipConfig{}, errors.New("Invalid coordinates")
		}
		bcs.Player1.MyShipsPos[model.Coordinates{X: x, Y: y}] = true
	}

	bcs.Player1.EnemyTarget = make([]model.Coordinates, 0)
	ena := strings.Split(a[5], ":")

	for i := 0; i < bcs.TotMissiles; i++ {
		ssx := strings.Split(ena[i], ",")
		x, err := strconv.Atoi(ssx[0])
		if err != nil {
			return &model.BattleShipConfig{}, err
		}
		y, err := strconv.Atoi(ssx[1])
		if err != nil {
			return &model.BattleShipConfig{}, err
		}
		bcs.Player1.EnemyTarget = append(bcs.Player1.EnemyTarget, model.Coordinates{X: x, Y: y})
	}

	bcs.Player2.MyShipsPos = make(map[model.Coordinates]bool)
	ssb := strings.Split(a[3], ":")

	for i := 0; i < bcs.Player2.MyShips; i++ {
		ssx := strings.Split(ssb[i], ",")
		x, err := strconv.Atoi(ssx[0])
		if err != nil {
			return &model.BattleShipConfig{}, err
		}
		y, err := strconv.Atoi(ssx[1])
		if err != nil {
			return &model.BattleShipConfig{}, err
		}
		if x >= bcs.Player2.MyShips || y >= bcs.Player2.MyShips {
			return &model.BattleShipConfig{}, errors.New("Invalid coordinates")
		}
		bcs.Player2.MyShipsPos[model.Coordinates{X: x, Y: y}] = true
	}
	bcs.Player2.EnemyTarget = make([]model.Coordinates, 0)
	enb := strings.Split(a[6], ":")

	for i := 0; i < bcs.TotMissiles; i++ {
		ssx := strings.Split(enb[i], ",")
		x, err := strconv.Atoi(ssx[0])
		if err != nil {
			return &model.BattleShipConfig{}, err
		}
		y, err := strconv.Atoi(ssx[1])
		if err != nil {
			return &model.BattleShipConfig{}, err
		}
		bcs.Player2.EnemyTarget = append(bcs.Player2.EnemyTarget, model.Coordinates{X: x, Y: y})
	}
	//bcs.Player1.MissedByOpponent = make([]model.Coordinates, bcs.TotMissiles)
	//bcs.Player2.MissedByOpponent = make([]model.Coordinates, bcs.TotMissiles)
	//make(map[model.Coordinates]bool)
	bcs.Player1.MissedByOpponent = make(map[model.Coordinates]bool)
	bcs.Player2.MissedByOpponent = make(map[model.Coordinates]bool)

	return bcs, nil
}
