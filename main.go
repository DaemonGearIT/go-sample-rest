package main

import (
	"fmt"
	"github.com/daemongear/go-sample-rest/config"
	"github.com/fatih/color"
)

func main() {
	color.Green("Welcome to Go-Sample-Rest by DaemonGear")
	color.Green("Step 1: Read Config File")
	config := &config.Config{}
	config.Configure()
}
