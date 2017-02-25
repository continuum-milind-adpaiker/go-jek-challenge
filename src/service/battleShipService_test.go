package service

import (
	"testing"

	"strings"

	dal "github.com/ContinuumLLC/go-jek/src/dataAccessLayer"
	"github.com/ContinuumLLC/go-jek/src/model"
)

type factories struct {
	dal.InputReaderFactoryImpl
	dal.OutputSaverFactoryImpl
	BattleShipServiceFactoryImpl
}

func TestPlayDraw(t *testing.T) {
	inputData := &model.BattleShipConfig{}
	inputData.TotMissiles = 5
	inputData.GridSize = 5
	inputData.Player1.MyShips = 5
	inputData.Player1.MyShipsPos = make(map[model.Coordinates]bool)
	inputData.Player1.EnemyTarget = make([]model.Coordinates, 0)
	inputData.Player1.MissedByOpponent = make(map[model.Coordinates]bool)

	inputData.Player1.MyShipsPos[model.Coordinates{X: 1, Y: 1}] = true
	inputData.Player1.MyShipsPos[model.Coordinates{X: 2, Y: 0}] = true
	inputData.Player1.MyShipsPos[model.Coordinates{X: 2, Y: 3}] = true
	inputData.Player1.MyShipsPos[model.Coordinates{X: 3, Y: 4}] = true
	inputData.Player1.MyShipsPos[model.Coordinates{X: 4, Y: 3}] = true

	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 0, Y: 1})
	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 4, Y: 3})
	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 2, Y: 3})
	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 3, Y: 1})
	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 2, Y: 1})

	inputData.Player2.MyShips = 5
	inputData.Player2.MyShipsPos = make(map[model.Coordinates]bool)
	inputData.Player2.EnemyTarget = make([]model.Coordinates, 0)
	inputData.Player2.MissedByOpponent = make(map[model.Coordinates]bool)

	inputData.Player2.MyShipsPos[model.Coordinates{X: 0, Y: 1}] = true
	inputData.Player2.MyShipsPos[model.Coordinates{X: 2, Y: 3}] = true
	inputData.Player2.MyShipsPos[model.Coordinates{X: 3, Y: 0}] = true
	inputData.Player2.MyShipsPos[model.Coordinates{X: 3, Y: 4}] = true
	inputData.Player2.MyShipsPos[model.Coordinates{X: 4, Y: 1}] = true

	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 4, Y: 1})
	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 2, Y: 0})
	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 1, Y: 2})
	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 2, Y: 3})
	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 4, Y: 4})

	expectedRes := "Draw"
	b := battleShipServiceImpl{}
	res := b.play(inputData)
	if strings.Compare(expectedRes, res) != 0 {
		t.Errorf("Expected result %s, got %s", expectedRes, res)
	}
}

func TestPlayBWins(t *testing.T) {
	inputData := &model.BattleShipConfig{}
	inputData.TotMissiles = 5
	inputData.GridSize = 5
	inputData.Player1.MyShips = 5
	inputData.Player1.MyShipsPos = make(map[model.Coordinates]bool)
	inputData.Player1.EnemyTarget = make([]model.Coordinates, 0)
	inputData.Player1.MissedByOpponent = make(map[model.Coordinates]bool)

	inputData.Player1.MyShipsPos[model.Coordinates{X: 1, Y: 1}] = true
	inputData.Player1.MyShipsPos[model.Coordinates{X: 2, Y: 0}] = true
	inputData.Player1.MyShipsPos[model.Coordinates{X: 2, Y: 3}] = true
	inputData.Player1.MyShipsPos[model.Coordinates{X: 3, Y: 4}] = true
	inputData.Player1.MyShipsPos[model.Coordinates{X: 4, Y: 3}] = true

	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 0, Y: 1})
	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 4, Y: 3})
	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 2, Y: 3})
	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 3, Y: 1})
	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 2, Y: 1})

	inputData.Player2.MyShips = 5
	inputData.Player2.MyShipsPos = make(map[model.Coordinates]bool)
	inputData.Player2.EnemyTarget = make([]model.Coordinates, 0)
	inputData.Player2.MissedByOpponent = make(map[model.Coordinates]bool)

	inputData.Player2.MyShipsPos[model.Coordinates{X: 0, Y: 1}] = true
	inputData.Player2.MyShipsPos[model.Coordinates{X: 2, Y: 3}] = true
	inputData.Player2.MyShipsPos[model.Coordinates{X: 3, Y: 0}] = true
	inputData.Player2.MyShipsPos[model.Coordinates{X: 3, Y: 4}] = true
	inputData.Player2.MyShipsPos[model.Coordinates{X: 4, Y: 1}] = true

	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 4, Y: 1})
	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 2, Y: 0})
	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 1, Y: 2})
	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 2, Y: 3})
	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 4, Y: 3})

	expectedRes := "B wins"
	b := battleShipServiceImpl{}
	res := b.play(inputData)
	if strings.Compare(expectedRes, res) != 0 {
		t.Errorf("Expected result %s, got %s", expectedRes, res)
	}
}

func TestPlayAWins(t *testing.T) {
	inputData := &model.BattleShipConfig{}
	inputData.TotMissiles = 5
	inputData.GridSize = 5
	inputData.Player1.MyShips = 5
	inputData.Player1.MyShipsPos = make(map[model.Coordinates]bool)
	inputData.Player1.EnemyTarget = make([]model.Coordinates, 0)
	inputData.Player1.MissedByOpponent = make(map[model.Coordinates]bool)

	inputData.Player1.MyShipsPos[model.Coordinates{X: 1, Y: 1}] = true
	inputData.Player1.MyShipsPos[model.Coordinates{X: 2, Y: 0}] = true
	inputData.Player1.MyShipsPos[model.Coordinates{X: 2, Y: 3}] = true
	inputData.Player1.MyShipsPos[model.Coordinates{X: 3, Y: 4}] = true
	inputData.Player1.MyShipsPos[model.Coordinates{X: 4, Y: 3}] = true

	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 0, Y: 1})
	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 4, Y: 3})
	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 2, Y: 3})
	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 3, Y: 1})
	inputData.Player1.EnemyTarget = append(inputData.Player1.EnemyTarget, model.Coordinates{X: 2, Y: 1})

	inputData.Player2.MyShips = 5
	inputData.Player2.MyShipsPos = make(map[model.Coordinates]bool)
	inputData.Player2.EnemyTarget = make([]model.Coordinates, 0)
	inputData.Player2.MissedByOpponent = make(map[model.Coordinates]bool)

	inputData.Player2.MyShipsPos[model.Coordinates{X: 0, Y: 1}] = true
	inputData.Player2.MyShipsPos[model.Coordinates{X: 2, Y: 3}] = true
	inputData.Player2.MyShipsPos[model.Coordinates{X: 3, Y: 0}] = true
	inputData.Player2.MyShipsPos[model.Coordinates{X: 3, Y: 4}] = true
	inputData.Player2.MyShipsPos[model.Coordinates{X: 4, Y: 1}] = true

	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 4, Y: 1})
	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 2, Y: 0})
	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 0, Y: 0})
	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 0, Y: 0})
	inputData.Player2.EnemyTarget = append(inputData.Player2.EnemyTarget, model.Coordinates{X: 0, Y: 0})

	expectedRes := "A wins"
	b := battleShipServiceImpl{}
	res := b.play(inputData)
	if strings.Compare(expectedRes, res) != 0 {
		t.Errorf("Expected result %s, got %s", expectedRes, res)
	}
}
