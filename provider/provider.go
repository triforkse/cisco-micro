package provider

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"cisco/micro/provider/aws"
	"cisco/micro/provider/gce"

	"log"
        "reflect"
        "fmt"
        "errors"
)

type Provider interface {
	Populate()
	ConfigId() string
	ProviderId() string
	Run(action func() error) error
	TerraformVars() map[string]string
	PackerVars() map[string]string
}

func New(providerId string) Provider {
	providers := map[string]Provider{
		"aws": new(aws.Config),
		"gce": new(gce.Config),
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

func ComplementVars(provider Provider, fieldName string, question string, complement func(string, string) (string, error)) (string, error) {
        config := reflect.TypeOf(provider).Elem()
        field, known := config.FieldByName(fieldName)

        if !known {
                return "", errors.New(fmt.Sprintf("The struct has no field with name '%s'", fieldName))
        }

        configValue := reflect.ValueOf(provider)
        configField := configValue.Elem().FieldByName(fieldName)
        defaultValue := configField.String()

        if field.Tag.Get("complement") == "true" {
                complementedValue, err := complement(question, defaultValue)
                if err != nil {
                        return "", err
                }
                return complementedValue, nil
        }

        return defaultValue, nil
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

func VarList(vars map[string]string) []string {
	args := []string{}
	for k, v := range vars {
		args = append(args, "-var", k+"="+v)
	}
	return args
}
