package main

import (
  "testing"
  "reflect"
)


func TestTerraformVarsAWS(t *testing.T) {
  provider := new(AWSProvider)
  provider.AccessKey = "MY_ACCESS"
  provider.SecretKey = "MY_SECRET"
  provider.Region = "MY_REGION"

  vars := provider.terraformVars()

  expected := map[string]string{
    "secret_key": "MY_SECRET",
    "access_key": "MY_ACCESS",
    "region": "MY_REGION",
  }

  if !reflect.DeepEqual(expected, vars) {
    t.Errorf("Expected %v, got %v", expected, vars)
  }
}
