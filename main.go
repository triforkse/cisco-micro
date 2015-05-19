package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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
	prepare()
	cleanup()
	terraformVars() map[string]string
}

func main() {
	filePath := flag.String("config", "micro.json", "the configuration file")

	flag.Parse()

	config := ReadConfig(*filePath)
	runTerraform(config)
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


func runTerraform(config Config) {

	provider := provider(config)

	provider.prepare()
	defer provider.cleanup()

	args := []string{"apply"}

	// Determine if we have an old tfstate file we need to load.
	args = append(args, "-state=" + filepath.Join(".micro", config.Id + ".tfstate"))

	// Pass in the arguments
	for k, v := range provider.terraformVars() {
		args = append(args, "-var", k + "=" + v)
	}
	args = append(args, "-var", "deployment_id=" + config.Id)

	// Tell it what template to use based on the provider.
	args = append(args, filepath.Join("templates", config.Provider))

	fmt.Printf("terraform %+v", args)

	// Run Terraform
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

func (p *AWSProvider) terraformVars() map[string]string {
	return map[string]string{
		"secret_key": p.SecretKey,
		"access_key": p.AccessKey,
		"region": p.Region,
	}
}
func (p *AWSProvider) prepare() { }
func (p *AWSProvider) cleanup() { }


type GCCProvider struct {
	// JSON Fields
	Project       string
	Region				string
	PrivateKeyId  string `json:"private_key_id"`
	PrivateKey 	  string `json:"private_key"`
	ClientEmail   string `json:"client_email"`
	ClientId	  	string `json:"client_id"`

	AccountFile   string `json:"-"` // Path to the temp file needed by terraform
}

func (p *GCCProvider) terraformVars() map[string]string {
	return map[string]string{
		"project":	p.Project,
		"region": 	p.Region,
		"nodes" : "1",
		"account_file": p.AccountFile,
	}
}

func (p *GCCProvider) prepare() {
	accountJson, _ := json.Marshal(map[string]string{
		"private_key_id": p.PrivateKeyId,
		"private_key":    p.PrivateKey,
		"client_email":   p.ClientEmail,
		"client_id":      p.ClientId,
	})

	accountFileName := filepath.Join(os.TempDir(), "account.json")
	err := ioutil.WriteFile(accountFileName, accountJson, 0600)
	if err != nil {
		log.Fatal("Could not write account file at: " + accountFileName)
	}

	p.AccountFile = accountFileName
}

func (p *GCCProvider) cleanup() {
	log.Printf("Removing file %v", p.AccountFile)
	os.Remove(p.AccountFile)
	// TODO Report error
}
