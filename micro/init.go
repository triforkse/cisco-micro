package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/triforkse/cisco-micro/logger"
	"github.com/triforkse/cisco-micro/provider"
)

func initCmd(providerId string, filePath string) {
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
