package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"log"
	"path/filepath"
)

type GCCProvider struct {
	// JSON Fields
	Project       string
	Region				string
	PrivateKeyId  string `json:"private_key_id"`
	PrivateKey 	  string `json:"private_key"`
	ClientEmail   string `json:"client_email"`
	ClientId	  	string `json:"client_id"`

	AccountFile   string `json:"-"` // Path to the temp file needed by terraform
}

func (p *GCCProvider) terraformVars() map[string]string {
	return map[string]string{
		"project":	p.Project,
		"region": 	p.Region,
		"nodes" : "1",
		"account_file": p.AccountFile,
	}
}

func (p *GCCProvider) prepare() {
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

  debugf("Google Account file: %s", accountFileName)

	p.AccountFile = accountFileName
}

func (p *GCCProvider) cleanup() {
	debugf("Removing file %v", p.AccountFile)
	os.Remove(p.AccountFile)
	// TODO Report error
}
