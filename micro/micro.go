package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"cisco/micro/logger"
	"cisco/micro/provider"
)

func main() {
	filePath := flag.String("config", "infrastructure.json", "the configuration file")
	downloadRepo := flag.Bool("clone", false, "should a packer project be downloaded")
	gitRepo := flag.String("repo", "https://github.com/CiscoCloud/microservices-infrastructure.git", "the reopostory for the infrastructure project")
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

	logger.Debugf("Git repo: %s", *gitRepo)
	logger.Debugf("Command: %s", command)
	logger.Debugf("Config File: %s", *filePath)

	switch command {
	case "init":
		providerId := cmdArgs[1]
		if *downloadRepo == true {
			clonePackerConfigProject(*gitRepo)
		}
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

func clonePackerConfigProject(gitRepo string) {

	defaultLocation := ".micro/src"

	dir, stat_err := os.Stat(defaultLocation)
	if stat_err == nil {
		logger.Debugf("Name %v", dir.Name())
		err := os.RemoveAll(defaultLocation)
		if err != nil {
			log.Fatal("Error removing " + dir.Name() + " " + err.Error())
		}
	}

	cmd := exec.Command("git", []string{"clone", "--depth=1", gitRepo, defaultLocation}...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal("Error cloning from git. ", err.Error())
	}
}
