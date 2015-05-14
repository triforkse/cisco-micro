package main

import "testing"
import "reflect"

func TestReadConfig(t *testing.T) {
  config := ReadConfig("testdata/aws_test.json")
  if config.Provider != "aws" {
    t.Error("expected provider 'aws'")
  }
}

func TestParseAWS(t *testing.T) {
  jsonData := []byte(`{
    "access_key": "MY_ACCESS",
    "secret_key": "MY_SECRET",
    "region": "MY_REGION"
  }`)
  cmdArgs, err := ParseAWS(jsonData)

  expected := []string{
    "apply",
    "-var", "secret_key=MY_SECRET",
    "-var", "access_key=MY_ACCESS",
    "-var", "region=MY_REGION",
    "templates/aws",
  }

  if !reflect.DeepEqual(expected, cmdArgs) {
    t.Errorf("Expected %v, got %v", expected, cmdArgs)
  }

  if err != nil {
    t.Error("The json should be parsable.")
  }
}
