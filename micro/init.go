package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	"cisco/micro/logger"
	"cisco/micro/provider"
  "flag"
  "os"
  "os/exec"
)

const defaultRepo string = "https://github.com/triforkse/microservices-infrastructure.git"

func initCmd(providerId string, filePath string) {

  // Download the Infrastructure files

  downloadRepo := flag.Bool("clone", true, "should a packer project be downloaded")
  gitRepo := flag.String("repo", defaultRepo, "the reopostory for the infrastructure project")

  logger.Debugf("Git repo: %s", *gitRepo)

  if *downloadRepo == true {
    clonePackerConfigProject(*gitRepo)
  }

  // Write Configuration

  if _, fileErr := os.Stat(filePath); fileErr != nil {
      
  }

	config := provider.New(providerId)
	config.Populate()

	logger.Debugf("Generating Config: %+v", config)

	data, err := json.Marshal(config)

	if err != nil {
		log.Fatal("Could not write configuration." + err.Error())
	}

	var out bytes.Buffer
	json.Indent(&out, data, "", "  ")

	err = ioutil.WriteFile(filePath, out.Bytes(), 0644)

	if err != nil {
		log.Fatal("Could not write configuration. " + err.Error())
	}
}


func clonePackerConfigProject(gitRepo string) {

  defaultLocation := ".micro/src"

  _, statErr := os.Stat(defaultLocation)

  if statErr != nil {
    cmd := exec.Command("git", "clone", "--depth=1", gitRepo, defaultLocation)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
      log.Fatal("Error cloning from git. ", err.Error())
    }
  }
}
