package provider

import (
        "reflect"
        "testing"

        "cisco/micro/provider/aws"
        "cisco/micro/provider/gce"
)

func TestReadConfig(t *testing.T) {
        config := FromFile("testdata/aws_test.json")

        if config.ProviderId() != "aws" {
                t.Error("expected provider 'aws'")
        }

        if config.ConfigId() != "test-aws-123" {
                t.Error("expected another 'id' attribute")
        }

        // Make sure we produce the correct type.

        var expectedType aws.Config
        if reflect.TypeOf(config) != reflect.TypeOf(&expectedType) {
                t.Error("the wrong config type was produced")
        }
}

func TestComplementVars(t *testing.T) {
        config := &gce.Config{Id: "test-id", ControlCount: "1", Datacenter: "gce"}

        //Return existing value
        complementedValue, err := ComplementVars(config, "Id", "question", func(q string, d string) (string, error) { return "another-id", nil})
        if err != nil {
                t.Errorf("Expected test-id got %v", err)
        }
        if complementedValue != "test-id" {
                t.Errorf("Expected test-id got %v", complementedValue)
        }

        //Return default value
        complementedValue, err = ComplementVars(config, "ControlCount", "question", func(q string, d string) (string, error) { return d, nil})
        if err != nil {
                t.Errorf("Expected 1 got %v", err)
        }
        if complementedValue != "1" {
                t.Errorf("Expected 1 got %v", complementedValue)
        }

        //Return changed value
        complementedValue, err = ComplementVars(config, "Datacenter", "question", func(q string, d string) (string, error) { return "aws", nil})
        if err != nil {
                t.Errorf("Expected aws got %v", err)
        }
        if complementedValue != "aws" {
                t.Errorf("Expected aws got %v", complementedValue)
        }

        //Return supplied value
        complementedValue, err = ComplementVars(config, "PrivateKey", "question", func(q string, d string) (string, error) { return "test-key", nil})
        if err != nil {
                t.Errorf("Expected test-key got %v", err)
        }
        if complementedValue != "test-key" {
                t.Errorf("Expected test-key got %v", complementedValue)
        }

        //Should return error for non existing fields
        complementedValue, err = ComplementVars(config, "non_existing", "question", func(q string, d string) (string, error) { return "test-key", nil})
        if err == nil {
                t.Errorf("Expected error got %v", complementedValue)
        }
}
