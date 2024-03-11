package main

//go:generate /go/bin/sqlboiler psql

import (
	"ajp-medical-clinic/api"
	"ajp-medical-clinic/config"
)

func main() {
	config.SetupDatabase()
	api.StartServer()
}


