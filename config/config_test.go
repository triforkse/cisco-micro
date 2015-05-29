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
        args := []string{"--all", "someCmd"}

        configs, remainingArgs := MatchConfigs(args)

        if len(remainingArgs) != 1 {
                t.Error("Command line argument for select clusters should be removed")
        }

        if remainingArgs[0] != "someCmd" {
                t.Error("Wrong command line argument removed")
        }

        if len(configs["aws"]) != 2 {
                t.Errorf("Expected two config for aws. But got %d", len(configs["aws"]))
        }

        if len(configs["gce"]) != 2 {
                t.Errorf("Expected two config for gce. But got %d", len(configs["aws"]))
        }

}

func TestParseMatchConfigsOneCluster(t *testing.T) {
        args := []string{`--id=aws-2`}

        configs, remainingArgs := MatchConfigs(args)

        if len(remainingArgs) != 0 {
                t.Error("Command line argument should be removed")
        }


        if _, ok := configs["aws"]; !ok {
                t.Errorf("Expected aws key", configs["aws"])
        }

        if len(configs["aws"]) != 1 {
                t.Errorf("Expected one config for aws. But got %d", len(configs["aws"]))
        }

        if configs["aws"][0].Config.Id != "aws-2" {
                t.Error("Expected another config id")
        }

        if configs["aws"][0].Config.Provider != "aws" {
                t.Error("Expected another config provider")
        }

        if configs["aws"][0].Path != "testdata/aws-cluster2.json" {
                t.Error("Expected another config path")
        }
}

func TestParseMatchConfigsServeralClusters(t *testing.T) {
        args := []string{`--id=aws-2`,`--id=gce-1`}

        configs, remainingArgs := MatchConfigs(args)

        if len(remainingArgs) != 0 {
                t.Error("Command line argument should be removed")
        }

        if _, ok :=configs["aws"]; !ok {
                t.Errorf("Expected aws key", configs)
        }

        if _, ok :=configs["gce"]; !ok {
                t.Errorf("Expected gce key", configs)
        }


        if !(configs["aws"][0].Config.Id == "aws-2") {
                t.Error("Expected another aws config id")
        }
        if !(configs["gce"][0].Config.Id == "gce-1") {
                t.Error("Expected another gce config id")
        }
}

func TestParseMatchConfigsNoMatchingClusters(t *testing.T) {
        args := []string{`--id=aws-not-found`}

        configs, remainingArgs := MatchConfigs(args)

        if len(remainingArgs) != 0 {
                t.Error("Command line argument should be removed")
        }

        if len(configs) != 0 {
                t.Error(`Expected no configs`)
        }

}

func TestParseMatchConfigsNoIdOrAllCliParameter(t *testing.T) {
        args := []string{`-debug`}

        configs, remainingArgs := MatchConfigs(args)

        if len(remainingArgs) != 0 {
                t.Error("Command line argument should be removed")
        }

        if len(configs) != 0 {
                t.Error(`Expected no configs`)
        }

}
