package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os/exec"
	"os"
	"path/filepath"
)

var isDebugging bool

type Config struct {
	Id         string
	Provider   string
	Properties json.RawMessage // Delay parsing until we know the provider
}

type Provider interface {
	prepare()
	cleanup()
	terraformVars() map[string]string
}

func main() {
	filePath := flag.String("config", "micro.json", "the configuration file")
	debugging := flag.Bool("debug", false, "show debug info")

	flag.Parse()

	isDebugging = *debugging
	config := ReadConfig(*filePath)

	var action string
	cmdArgs := flag.Args()
	if len(cmdArgs) > 0 {
		action = cmdArgs[len(cmdArgs) - 1]
	}	else {
		action = "apply"
	}

	debugf("Action: %s", action)

	runTerraform(action, config)
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


func provider(config Config) Provider {
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

	return provider
}


func runTerraform(action string, config Config) {

	provider := provider(config)

	provider.prepare()
	defer provider.cleanup()

	args := []string{action}

	// Determine if we have an old tfstate file we need to load.
	args = append(args, "-state=" + filepath.Join(".micro", config.Id + ".tfstate"))

	// Pass in the arguments
	for k, v := range provider.terraformVars() {
		args = append(args, "-var", k + "=" + v)
	}
	args = append(args, "-var", "deployment_id=" + config.Id)

	// Tell it what template to use based on the provider.
	args = append(args, filepath.Join("templates", config.Provider))

	debugf("terraform %+v", args)

	// Run Terraform
	cmd := exec.Command("terraform", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if cmd.Run() != nil {
		os.Exit(1)
	}

	printTable("Cluster Properties", map[string]string{
		"ID": config.Id,
	})
}
