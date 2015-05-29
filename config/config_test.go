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

func TestParseMatchConfigsForAllCluster(t *testing.T) {
        args := []string{"-all", "someCmd"}

        configs, remainingArgs := MatchConfigs(args)

        if len(remainingArgs) != 1 {
                t.Error("Command line argument for select clusters should be removed")
        }

        if remainingArgs[0] != "someCmd" {
                t.Error("Wrong command line argument removed")
        }

        if len(configs) != 4 {
                t.Error(`Expected four configs`)
        }
}

func TestParseMatchConfigsOneCluster(t *testing.T) {
        args := []string{`-id=aws-2`}

        configs, remainingArgs := MatchConfigs(args)

        if len(remainingArgs) != 0 {
                t.Error("Command line argument should be removed")
        }

        if len(configs) != 1 {
                t.Error(`Expected four configs`)
        }

        if configs[0].Config.Id != "aws-2" {
                t.Error("Expected another config id")
        }

        if configs[0].Config.Provider != "aws" {
                t.Error("Expected another config provider")
        }

        if configs[0].Path != "testdata/aws-cluster2.json" {
                t.Error("Expected another config path")
        }
}

func TestParseMatchConfigsServeralClusters(t *testing.T) {
        args := []string{`-id=aws-2,gce-1`}

        configs, remainingArgs := MatchConfigs(args)

        if len(remainingArgs) != 0 {
                t.Error("Command line argument should be removed")
        }

        if len(configs) != 2 {
                t.Error(`Expected two configs`)
        }

        if !(configs[0].Config.Id == "aws-2" || configs[0].Config.Id == "gce-1") {
                t.Error("Expected another config id")
        }
        if !(configs[1].Config.Id == "aws-2" || configs[1].Config.Id == "gce-1") {
                t.Error("Expected another config id")
        }
}

func TestParseMatchConfigsNoMatchingClusters(t *testing.T) {
        args := []string{`-id=aws-not-found`}

        configs, remainingArgs := MatchConfigs(args)

        if len(remainingArgs) != 0 {
                t.Error("Command line argument should be removed")
        }

        if len(configs) != 0 {
                t.Error(`Expected no configs`)
        }

}
