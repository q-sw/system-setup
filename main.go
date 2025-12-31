package main

import (
	"flag"
	"fmt"
	"time"

	systemsetup "github.com/q-sw/system-setup/cmd/systemSetup"
)

func main() {
	startTime := time.Now()
	var configFilePath string
	flag.StringVar(&configFilePath, "path", "", "Path of the config file to use")
	initConfigFile := flag.Bool("init", false, "Create empty config file")

	flag.Parse()

	switch {
	case *initConfigFile:
		fmt.Println("Init the configuration file")
		systemsetup.InitConfiguration()
	case configFilePath != "":
		fmt.Println("Custom configuration is used")
		systemsetup.Setup(configFilePath)
	default:
		fmt.Println("Start System configuration by default")
		systemsetup.InitConfiguration()
		systemsetup.Setup("./config.yaml")
	}

	duration := time.Since(startTime)
	fmt.Printf("Configuration time: %v\n", duration)

}
