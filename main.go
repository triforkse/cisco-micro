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
	Id         string
	Provider   string
	Properties json.RawMessage // Delay parsing until we know the provider
}

type Provider interface {
	cmdArgs() []string
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

	parsers := map[string]Provider{
		"aws": new(AWSProvider),
		"gcc": new(GCCProvider),
	}

	provider, known := parsers[config.Provider]

	if !known {
		log.Fatal("Unknown provider: '" + config.Provider + "'")
	}

	err := json.Unmarshal(config.Properties, provider)
	if err != nil {
		log.Fatal("Invalid configuration. " + err.Error())
	}

	// TODO: call provider.prepare()

	args := provider.cmdArgs()

	// Determine if we have an old tfstate file we need to load.

	stateFileArg := "-state=" + filepath.Join(".micro", config.Id + ".tfstate")
	args = append([]string{"apply", stateFileArg}, args...)

	fmt.Printf("terraform %+v", args)
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

type AWSProvider struct {
	SecretKey string `json:"secret_key"`
	AccessKey string `json:"access_key"`
	Region    string
}

func (p *AWSProvider) cmdArgs() []string {
	return []string{
		"-var", "secret_key=" + p.SecretKey,
		"-var", "access_key=" + p.AccessKey,
		"-var", "region=" + p.Region,
		"templates/aws",
	}
}


type GCCProvider struct {
	Project       string
	Region				string
}

func (p *GCCProvider) cmdArgs() []string {
	return []string{
		"-var", "project=" + p.Project,
		"-var", "region=" + p.Region,
		"-var", "account_file=account.json",
		"templates/gcc",
	}
}
