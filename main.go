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

type Provider interface {
	populate()
	configId() string
	providerId() string
	prepare()
	cleanup()
	terraformVars() map[string]string
}

func main() {
	filePath := flag.String("config", "infrastructure.json", "the configuration file")
	debugging := flag.Bool("debug", false, "show debug info")

	flag.Parse()

	isDebugging = *debugging

	var command string
	cmdArgs := flag.Args()
	if len(cmdArgs) > 0 {
		command = cmdArgs[0]
	}	else {
		command = "apply"
	}

	debugf("Command: %s", command)
	debugf("Config File: %s", *filePath)

	switch command {
		case "init":
			providerId := cmdArgs[1]
			initCmd(providerId, *filePath)
		case "apply", "destroy", "plan":
			config := provider(*filePath)
			terraformCmd(command, config)
	}
}

func readProviderFile(filePath string) (providerId string, bytes []byte, err error) {

	// Determine what provider we are using,
	// and parse the configuration accordingly.

	bytes, _ = ioutil.ReadFile(filePath)

	var config struct {
		Provider string
	}
	err = json.Unmarshal(bytes, &config)

	if err == nil {
		providerId = config.Provider
	}

	return
}


func provider(filePath string) Provider {

	providerId, bytes, err := readProviderFile(filePath)
	if err != nil {
		absPath, _ := filepath.Abs(filePath)
		log.Fatal("Failed to read configuration file: " + absPath)
	}

	provider := newProvider(providerId)

	err = json.Unmarshal(bytes, provider)
	if err != nil {
		log.Fatal("Invalid configuration. " + err.Error())
	}

	return provider
}


func terraformCmd(command string, provider Provider) {

	provider.prepare()
	defer provider.cleanup()

	args := []string{command}

	// Determine if we have an old tfstate file we need to load.
	args = append(args, "-state=" + filepath.Join(".micro", provider.configId() + ".tfstate"))

	// Pass in the arguments
	for k, v := range provider.terraformVars() {
		args = append(args, "-var", k + "=" + v)
	}
	args = append(args, "-var", "deployment_id=" + provider.configId())

	// Tell it what template to use based on the provider.
	args = append(args, filepath.Join("templates", provider.providerId()))

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
		"Type": provider.providerId(),
		"ID": provider.configId(),
	})
}
