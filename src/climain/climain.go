package main

import "flag"

func main() {
	inputPtr := flag.String("input", "../input.txt", "default input for battleship programm")
	flag.Parse()
	f := factories{}
	battleShip := f.GetBattleShipService(f, *inputPtr)
	battleShip.Start()
}
