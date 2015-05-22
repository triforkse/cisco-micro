package gce

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestGCEPrepare(t *testing.T) {
	target := Config{
		Id:           "foo",
		Provider:     "gce",
		PrivateKeyId: "test_private_key_id",
		PrivateKey:   "test_private_key",
		ClientEmail:  "test_client_email",
		ClientId:     "test_client_id",
	}
	original := accountFileFromConfig(target)

	// Check that the account file can be found.

	target.Run(func() error {
		bytes, _ := ioutil.ReadFile(target.AccountFile)

		var accountFile AccountFile
		json.Unmarshal(bytes, &accountFile)

		if !reflect.DeepEqual(accountFile, original) {
			t.Errorf("Expected %#v, got %#v", original, accountFile)
		}

		return nil
	})
}

func TestTerraformVarsGCE(t *testing.T) {
	provider := new(Config)
	provider.Region = "MY_REGION"
	provider.Project = "MY_PROJECT"
	provider.AccountFile = "account.json"

	vars := provider.TerraformVars()

	expected := map[string]string{
		"region":       "MY_REGION",
		"project":      "MY_PROJECT",
		"nodes":        "1",
		"account_file": "account.json",
	}

	if !reflect.DeepEqual(expected, vars) {
		t.Errorf("Expected %v, got %v", expected, vars)
	}
}
