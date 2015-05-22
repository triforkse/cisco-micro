package gce

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"cisco/micro/logger"
	"cisco/micro/util/strutil"
)

type Config struct {
	// JSON Fields
	Id           string `json:"id"`
	Provider     string `json:"provider"`
	Project      string `json:"project"`
	Region       string `json:"region"`
	PrivateKeyId string `json:"private_key_id"`
	PrivateKey   string `json:"private_key"`
	ClientEmail  string `json:"client_email"`
	ClientId     string `json:"client_id"`
	Zone         string `json:"zone"`

	AccountFile string `json:"-"` // Path to the temp file needed by terraform
}

func (p *Config) TerraformVars() map[string]string {
	return map[string]string{
		"project":      p.Project,
		"region":       p.Region,
		"nodes":        "1",
		"account_file": p.AccountFile,
	}
}

func (p *Config) Run(action func() error) error {
	// TODO: handle error
	accountFile := accountFileFromConfig(*p)
	accountJson, _ := json.Marshal(accountFile)

	accountFileName := filepath.Join(os.TempDir(), "account.json")
	err := ioutil.WriteFile(accountFileName, accountJson, 0600)
	if err != nil {
		log.Fatal("Could not write account file at: " + accountFileName)
	}

	logger.Debugf("Google Account File: %s", accountFileName)

	p.AccountFile = accountFileName

	err = action()

	logger.Debugf("Removing file %v", p.AccountFile)
	os.Remove(p.AccountFile)
	// TODO Report any error from remove

	return err
}

func (p *Config) ConfigId() string {
	return p.Id
}

func (p *Config) ProviderId() string {
	return "gce"
}

func (p *Config) Populate() {
	p.Id = strutil.Random(16)
	p.Provider = "gce"
	p.Region = "eu-west-1"
	p.Region = "eu"
	p.Zone = "us-central1-a"
	// TODO: Ask for user input for these.
	p.Project = "REPLACE WITH YOUR ACCESS KEY"
	p.PrivateKeyId = "REPLACE WITH YOUR PRIVATE KEY ID FROM YOUR ACCOUNT FILE"
	p.PrivateKey = "REPLACE WITH YOUR PRIVATE KEY FROM YOUR ACCOUNT FILE"
	p.ClientEmail = "REPLACE WITH YOUR CLIENT EMAIL FROM YOUR ACCOUNT FILE"
	p.ClientId = "REPLACE WITH YOUR CLIENT ID FROM YOUR ACCOUNT FILE"
}

func (p *Config) PackerVars() map[string]string {
	return map[string]string{
		"gce_creds_file": p.AccountFile,
		"gce_project_id": p.Project,
		"gce_zone":       p.Zone,
	}
}
