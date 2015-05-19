package main

import (
  "testing"
  "reflect"
  "io/ioutil"
  "encoding/json"
  "os"
)

func TestReadConfig(t *testing.T) {
  config := ReadConfig("testdata/aws_test.json")

  if config.Provider != "aws" {
    t.Error("expected provider 'aws'")
  }

  if config.Id != "test-aws-123" {
    t.Error("expected another 'id' attribute")
  }
}

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

func TestTerraformVarsGCC(t *testing.T) {
  provider := new(GCCProvider)
  provider.Region = "MY_REGION"
  provider.Project = "MY_PROJECT"
  provider.AccountFile = "account.json"
  //provider.PrivateKeyId = "MY_PRIVATE_KEY_ID"
  //provider.PrivateKey = "MY_PRIVATE_KEY"
  //provider.ClientEmail = "MY_CLIENT_EMAIL"

  vars := provider.terraformVars()

  expected := map[string]string{
    "region": "MY_REGION",
    "project": "MY_PROJECT",
    "nodes": "1",
    "account_file": "account.json",
    //"private_key_id": "MY_PRIVATE_KEY_ID",
    //"private_key": "MY_PRIVATE_KEY",
    //"client_email": "MY_CLIENT_EMAIL",
  }

  if !reflect.DeepEqual(expected, vars) {
    t.Errorf("Expected %v, got %v", expected, vars)
  }
}

func TestGccPrepare(t *testing.T) {
  target := GCCProvider{
    PrivateKeyId: "test_private_key_id",
    PrivateKey: "test_private_key",
    ClientEmail: "test_client_email",
	  ClientId: "test_client_id",
  }

  target.prepare()
  bytes, _ := ioutil.ReadFile(target.AccountFile)

  type AccountStub struct {
    PrivateKeyId  string `json:"private_key_id"`
    PrivateKey 	  string `json:"private_key"`
    ClientEmail   string `json:"client_email"`
    ClientId	  string `json:"client_id"`
  }

  var accountStub AccountStub
  json.Unmarshal(bytes, &accountStub)

  expected := AccountStub{
    PrivateKeyId: "test_private_key_id",
    PrivateKey: "test_private_key",
    ClientEmail: "test_client_email",
	  ClientId: "test_client_id",
  }

  if !reflect.DeepEqual(expected, accountStub) {
    t.Errorf("Expected %v, got %v", expected, accountStub)
  }

  defer os.Remove(target.AccountFile)
}
