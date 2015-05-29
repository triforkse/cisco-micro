package config

import (
	"reflect"
	"testing"
)

func TestFilterOutNonJsonPaths(t *testing.T) {
	paths := []string{"abc.json", "bad.file", "my.json"}

	valid := validConfigPaths(paths)

	if len(valid) != 2 {
		t.Error("expected 2 valid paths")
	}

	if !reflect.DeepEqual(valid, []string{"abc.json", "my.json"}) {
		t.Error("expected abc.json and my.json")
	}
}

func TestParseSampleConfig(t *testing.T) {
	sample := []byte(`{ "id": "123", "provider": "aws" }`)

	config, err := parseJsonBytes(sample)
	if err != nil {
		t.Error("Failed to parse JSON")
	}

	if config.Id != "123" {
		t.Error(`Expected id = "123"`)
	}

	if config.Provider != "aws" {
		t.Error(`Expected provider = "aws"`)
	}
}
