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
        provider.Populate()

        provider.Region = "MY_REGION"
        provider.Project = "MY_PROJECT"
        provider.AccountFile = "account.json"

        vars := provider.TerraformVars()

        expected := map[string]string{
                "region":       "MY_REGION",
                "project":      "MY_PROJECT",
                "nodes":        "1",
                "account_file": "account.json",

                "control_count": "3",
                "control_type" : "n1-standard-1",
                "datacenter"   : "gce",
                "long_name"    : "microservices-infrastructure",
                "network_ipv4" : "10.0.0.0/16",
                "short_name"   : "mi",
                "worker_count" : "1",
                "worker_type"  : "n1-highcpu-2",
                "ssh.username" : "REPLACE WITH THE PUBLIC SSH KEY YOU WANT TO USE",
                "ssh.key"      : "REPLACE WITH THE SSH USERNAME YOU WANT TO USE",
        }

        if expected["region"]  != vars["region"] &&
           expected["project"] != vars["project"] &&
           expected["nodes"]   != vars["nodes"] &&
           expected["account_file"] != vars["account_file"] &&
           expected["datacenter"] != vars["datacenter"] &&
           expected["ssh.key"] != vars["ssh.key"] &&
           expected["ssh.username"] != vars["ssh.username"] {
                t.Errorf("Expected %v, got %v", expected, vars)
        }

}
