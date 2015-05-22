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
  "cisco/micro/util/executil"
  "cisco/micro/term"
  "fmt"
  "errors"
)

const defaultRepo string = "https://github.com/triforkse/microservices-infrastructure.git"

func initCmd(providerId string, filePath string) error {

  // Download the Infrastructure files

  downloadRepo := flag.Bool("clone", true, "should a packer project be downloaded")
  gitRepo := flag.String("repo", defaultRepo, "the reopostory for the infrastructure project")

  logger.Debugf("Git repo: %s", *gitRepo)

  if *downloadRepo == true {
    clonePackerConfigProject(*gitRepo)
  }

  // Write Configuration

  if _, fileErr := os.Stat(filePath); fileErr != nil {
    return generateConfig(filePath, providerId)
  } else {
    logger.Debugf("Is file:")
    overwrite, err := term.AskForConfirmation(fmt.Sprintf("Are you sure you want to replace %s?, [yes/any]", filePath))
    if err != nil {
      return err
    }
    if overwrite == true {
      return generateConfig(filePath, providerId)
    }
  }

  return nil
}

func generateConfig(filePath string, providerId string) error {
  config := provider.New(providerId)
  config.Populate()

  logger.Debugf("Generating Config: %+v", config)

  data, err := json.Marshal(config)

  if err != nil {
    return errors.New("Could not write configuration." + err.Error())
  }

  var out bytes.Buffer
  json.Indent(&out, data, "", "  ")

  err = ioutil.WriteFile(filePath, out.Bytes(), 0644)

  if err != nil {
    return errors.New("Could not write configuration. " + err.Error())
  }

  return nil
}


func clonePackerConfigProject(gitRepo string) {

  defaultLocation := ".micro/src"

  _, statErr := os.Stat(defaultLocation)

  if statErr != nil {
    cmd := executil.Command("git", "clone", "--depth=1", gitRepo, defaultLocation)

    err := cmd.Run()
    if err != nil {
      log.Fatal("Error cloning from git. ", err.Error())
    }
  }
}
