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
	fmt.Printf("%+v", args);
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

/*
module "gcc" {
                   source = "git::https://gitlab.trifork.se/flg/ms-infra-terraform-ccp.git//gcc?ref=master"
                   account_file="{{account_file}}"
                   gce_ssh_user="554985525398-p9se88l5e3fupvj1v8t6tujq5qsumh1q.apps.googleusercontent.com"
                   gce_ssh_private_key_file="pkey"
                   region ="europe-west1"
                   zone="europe-west1-d"
                   project="cs-cisco"
                   image="ubuntu-os-cloud/ubuntu-1404-trusty-v20150128"
                   master_machine_type="n1-standard-2"
                   slave_machine_type= "n1-standard-4"
                   network= "10.20.30.0/24"
                   localaddress="92.111.228.8/32"
                   domain="gcc.trifork.se"
                   name="mymesoscluster"
                   masters= "1"
              }
*/


type AWSProvider struct {
	SecretKey string `json:"secret_key"`
	AccessKey string `json:"access_key"`
	Region    string
}

func (p *AWSProvider) cmdArgs() []string {
	return []string{
		"apply",
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
		"apply",
		"-var", "project=" + p.Project,
		"-var", "region=" + p.Region,
		"-var", "account_file=account.json",
		"templates/gcc",
	}
}
