package main

import (
	dal "github.com/ContinuumLLC/go-jek/src/dataAccessLayer"
	service "github.com/ContinuumLLC/go-jek/src/service"
)

type factories struct {
	dal.InputReaderFactoryImpl
	dal.OutputSaverFactoryImpl
	service.BattleShipServiceFactoryImpl
}
