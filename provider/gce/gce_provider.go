package gce

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/triforkse/cisco-micro/logger"
	"github.com/triforkse/cisco-micro/util/strutil"
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

func (p *Config) Prepare() {
	accountJson, _ := json.Marshal(map[string]string{
		"private_key_id": p.PrivateKeyId,
		"private_key":    p.PrivateKey,
		"client_email":   p.ClientEmail,
		"client_id":      p.ClientId,
	})

	accountFileName := filepath.Join(os.TempDir(), "account.json")
	err := ioutil.WriteFile(accountFileName, accountJson, 0600)
	if err != nil {
		log.Fatal("Could not write account file at: " + accountFileName)
	}

	logger.Debugf("Google Account file: %s", accountFileName)

	p.AccountFile = accountFileName
}

func (p *Config) Cleanup() {
	logger.Debugf("Removing file %v", p.AccountFile)
	os.Remove(p.AccountFile)
	// TODO Report error
}

func (p *Config) ConfigId() string {
	return p.Id
}

func (p *Config) ProviderId() string {
	return "gcc"
}

func (p *Config) Populate() {
	p.Id = strutil.Random(16)
	p.Provider = "gcc"
	p.Region = "eu-west-1"
	p.Project = "REPLACE WITH YOUR ACCESS KEY"
	p.Region = "eu"
	p.PrivateKeyId = "REPLACE WITH YOUR PRIVATE KEY ID FROM YOUR ACCOUNT FILE"
	p.PrivateKey = "REPLACE WITH YOUR PRIVATE KEY FROM YOUR ACCOUNT FILE"
	p.ClientEmail = "REPLACE WITH YOUR CLIENT EMAIL FROM YOUR ACCOUNT FILE"
	p.ClientId = "REPLACE WITH YOUR CLIENT ID FROM YOUR ACCOUNT FILE"
}
