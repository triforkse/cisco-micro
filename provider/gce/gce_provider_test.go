package gce

import (
  "testing"
  "reflect"
  "io/ioutil"
  "encoding/json"
  "os"
)

func TestGccPrepare(t *testing.T) {
  target := Provider{
    PrivateKeyId: "test_private_key_id",
    PrivateKey: "test_private_key",
    ClientEmail: "test_client_email",
	  ClientId: "test_client_id",
  }

  target.Prepare()
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

func TestTerraformVarsGCC(t *testing.T) {
  provider := new(Provider)
  provider.Region = "MY_REGION"
  provider.Project = "MY_PROJECT"
  provider.AccountFile = "account.json"

  vars := provider.TerraformVars()

  expected := map[string]string{
    "region": "MY_REGION",
    "project": "MY_PROJECT",
    "nodes": "1",
    "account_file": "account.json",
  }

  if !reflect.DeepEqual(expected, vars) {
    t.Errorf("Expected %v, got %v", expected, vars)
  }
}
