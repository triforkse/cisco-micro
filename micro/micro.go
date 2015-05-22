package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/triforkse/cisco-micro/logger"
	"github.com/triforkse/cisco-micro/provider"
)

func main() {
	filePath := flag.String("config", "infrastructure.json", "the configuration file")
	isDebugging := flag.Bool("debug", false, "show debug info")
	gitRepo := flag.String("gitrepo", "https://github.com/CiscoCloud/microservices-infrastructure.git", "the reopostory for the infrastructure project")

	flag.Parse()

	logger.EnableDebug(*isDebugging)

	var command string
	cmdArgs := flag.Args()
	if len(cmdArgs) > 0 {
		command = cmdArgs[0]
	} else {
		command = "apply"
	}

	logger.Debugf("Git repo: %s", *gitRepo)
	logger.Debugf("Command: %s", command)
	logger.Debugf("Config File: %s", *filePath)

	installMsInfra(*gitRepo)

	switch command {
	case "init":
		providerId := cmdArgs[1]
		installMsInfra(*gitRepo)
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

func installMsInfra(gitRepo string) {

	cmd := exec.Command("git", []string{"clone", "--depth=1", gitRepo, ".micro/src"}...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal("Error cloning from git. ", err.Error())
	}
}
