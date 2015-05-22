package main

import (
	"flag"
	"log"
	"cisco/micro/logger"
	"cisco/micro/provider"
)

func main() {
	filePath := flag.String("config", "infrastructure.json", "the configuration file")
	isDebugging := flag.Bool("debug", false, "show debug info")

	flag.Parse()

	logger.EnableDebug(*isDebugging)

	var command string
	cmdArgs := flag.Args()
	if len(cmdArgs) > 0 {
		command = cmdArgs[0]
	} else {
		command = "apply"
	}

	logger.Debugf("Command: %s", command)
	logger.Debugf("Config File: %s", *filePath)

	switch command {
	case "init":
		providerId := cmdArgs[1]
		initCmd(providerId, *filePath)
	case "apply", "destroy", "plan":
		config := provider.FromFile(*filePath)
		// TODO: handle read error here, not in the lib
		terraformCmd(command, config)
	case "build":
		err := packerCmd("build", provider.FromFile(*filePath))
		if err != nil {
			log.Fatal("Could not run packer. " + err.Error())
		}
	}
}
