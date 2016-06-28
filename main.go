package main

import (
	"fmt"
	"github.com/DaemonGearIT/go-sample-rest/config"
	"github.com/DaemonGearIT/go-sample-rest/infraestructure"
	"github.com/fatih/color"
)

func main() {
	color.Green("Welcome to Go-Sample-Rest by DaemonGear")
	color.Green("*Read Config File {")
	conf := config.Configure()
	postgres := infraestructure.Postgres{}
	err := postgres.Connect(conf)

	if err != nil {
		color.Red(fmt.Sprintf("Error: %v", err.Error()))
	}

	color.Green("} Read Config File")
	color.Green("")
	color.Green("Run Server {")
	server := infraestructure.Server{postgres}
	server.Run()
	color.Green("} Run Server ")
}
