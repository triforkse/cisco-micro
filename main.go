package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

type Config struct {
	Provider   string
	Properties json.RawMessage // Delay parsing until we know the provider
}

func main() {
	filePath := flag.String("config", "micro.json", "the configuration file")

	flag.Parse()

	config := ReadConfig(*filePath)
	RunTerraform(config)
}

func ReadConfig(filePath string) Config {

	// Determine what provider we are using,
	// and parse the configuration accordingly.

	bytes, _ := ioutil.ReadFile(filePath)

	var config Config
	parseErr := json.Unmarshal(bytes, &config)

	if parseErr != nil {
		absPath, _ := filepath.Abs(filePath)
		log.Fatal("Failed to read configuration file: " + absPath)
	}

	return config
}

func RunTerraform(config Config) {

	type Parser func(json.RawMessage) ([]string, error)

	parsers := map[string]Parser{
		"aws": ParseAWS,
	}

	parser, known := parsers[config.Provider]

	if !known {
		log.Fatal("Unknown provider: '" + config.Provider + "'")
	}

	args, err := parser(config.Properties)

	if err != nil {
		log.Fatal("Invalid configuration. " + err.Error())
	}

	cmd := exec.Command("terraform", args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	var outErr bytes.Buffer
	cmd.Stderr = &outErr

	if cmd.Run() != nil {
		log.Fatal(outErr.String())
	}

	fmt.Printf("%s", out.String())
}

func ParseAWS(jsonData json.RawMessage) ([]string, error) {

	var props struct {
		SecretKey string `json:"secret_key"`
		AccessKey string `json:"access_key"`
		Region    string
	}

	parseErr := json.Unmarshal(jsonData, &props)

	if parseErr != nil {
		return nil, parseErr
	} else {
		return []string{
			"apply",
			"-var", "secret_key=" + props.SecretKey,
			"-var", "access_key=" + props.AccessKey,
			"-var", "region=" + props.Region,
			"templates/aws",
		}, nil
	}
}
