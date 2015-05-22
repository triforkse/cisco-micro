package provider

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/triforkse/cisco-micro/provider/aws"
	"github.com/triforkse/cisco-micro/provider/gce"
)

import (
	"log"
)

type Provider interface {
	Populate()
	ConfigId() string
	ProviderId() string
	Prepare()
	Cleanup()
	TerraformVars() map[string]string
}

func New(providerId string) Provider {
	providers := map[string]Provider{
		"aws": new(aws.Config),
		"gcc": new(gce.Config),
	}

	provider, known := providers[providerId]

	if !known {
		log.Fatal("Unknown provider: '" + providerId + "'")
	}

	return provider
}

func readFile(filePath string) (providerId string, bytes []byte, err error) {

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

func FromFile(filePath string) Provider {

	providerId, bytes, err := readFile(filePath)
	if err != nil {
		absPath, _ := filepath.Abs(filePath)
		log.Fatal("Failed to read configuration file: " + absPath)
	}

	provider := New(providerId)

	err = json.Unmarshal(bytes, provider)
	if err != nil {
		log.Fatal("Invalid configuration. " + err.Error())
	}

	return provider
}
