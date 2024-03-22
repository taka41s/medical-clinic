package main

//go:generate /go/bin/sqlboiler psql

import (
	"ajp-medical-clinic/api"
	"ajp-medical-clinic/config"
	"ajp-medical-clinic/config/test"
)

func main() {
	config.SetupDatabase()
	api.StartServer()
}

func TestSetup() {
	TestConfig.SetupDatabase()
}


