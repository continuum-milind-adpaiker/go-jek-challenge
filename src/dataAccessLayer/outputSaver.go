package dal

import (
	"os"

	model "github.com/ContinuumLLC/go-jek/src/model"
)

// ConfigServiceFactoryImpl ...
type OutputSaverFactoryImpl struct{}

type outputSaverDalImpl struct{}

// GetOutputSaverDal ...
func (OutputSaverFactoryImpl) GetOutputSaverDal() model.OutputSaverDal {
	return &outputSaverDalImpl{}
}

//SaveOutput saves the result to output.txt in the form of grid
func (cs outputSaverDalImpl) SaveOutput(resultData *model.BattleShipConfig) error {
	f, err := os.Create("../playerAoutput.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString("Player A\n")
	save(f, resultData.GridSize, resultData.Player1)
	f.WriteString("\n\nPlayer B\n")
	save(f, resultData.GridSize, resultData.Player2)

	return nil
}

func save(f *os.File, gridSize int, player model.Player) {
	//Save player1 output
	for i := 0; i < gridSize; i++ {
		for k := 0; k < gridSize; k++ {
			value, ok := player.MyShipsPos[model.Coordinates{X: i, Y: k}]
			if ok == true {
				if value == false {
					f.WriteString("X ")
					continue
				} else {
					f.WriteString("B ")
					continue
				}

			}
			value, ok = player.MissedByOpponent[model.Coordinates{X: i, Y: k}]
			if ok == true {
				f.WriteString("O ")
				continue
			}

			f.WriteString("_ ")
		}
		f.WriteString("\n")
	}

}
