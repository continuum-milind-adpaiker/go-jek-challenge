package model

//import "github.com/ContinuumLLC/go-jek/src/model"

type Coordinates struct {
	X, Y int
}

type Player struct {
	MyShips int
	//position of my ship and its current status. True for alive, false for destroyed
	MyShipsPos  map[Coordinates]bool
	EnemyTarget []Coordinates
	//MissedByOpponent []Coordinates
	MissedByOpponent map[Coordinates]bool
}

type BattleShipConfig struct {
	Player1     Player
	Player2     Player
	TotMissiles int
	GridSize    int
}

// InputReaderDalFactory interface returns the ConfigDal
type InputReaderDalFactory interface {
	GetInputReaderDal(string) InputReaderDal
}

// InputReaderDal ...
type InputReaderDal interface {
	ReadConfigFile() (*BattleShipConfig, error)
}

// OutputSaverDalFactory ...
type OutputSaverDalFactory interface {
	GetOutputSaverDal() OutputSaverDal
}

// OutputSaverDal ...
type OutputSaverDal interface {
	SaveOutput(resultData *BattleShipConfig) error
}

// BattleShipService interface to handle all dal operations of Config
type BattleShipService interface {
	Start()
}

type BattleShipSvcDep interface {
	InputReaderDalFactory
	OutputSaverDalFactory
}
