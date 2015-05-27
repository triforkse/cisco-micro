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
	Region       string `json:"region" complement:"true"`
	PrivateKeyId string `json:"private_key_id" complement:"true"`
	PrivateKey   string `json:"private_key" complement:"true"`
	ClientEmail  string `json:"client_email" complement:"true"`
	ClientId     string `json:"client_id" complement:"true"`
	Zone         string `json:"zone"`

        ControlCount string `json:"control_count" complement:"true"`
        ControlType  string `json:"control_type" complement:"true"`
        Datacenter   string `json:"datacenter" complement:"true"`
        LongName     string `json:"long_name" complement:"true"`
        NetworkIpv4  string `json:"network_ipv4" complement:"true"`
        ShortName    string `json:"short_name" complement:"true"`
        WorkerCount  string `json:"worker_count" complement:"true"`
        WorkerType   string `json:"worker_type" complement:"true"`
        SshUserName  string `json:"ssh_username" complement:"true"`
        SshKey       string `json:"ssh_key" complement:"true"`

	AccountFile string `json:"-"` // Path to the temp file needed by terraform
}

func (p *Config) TerraformVars() map[string]string {
	return map[string]string{
		"project":      p.Project,
		"region":       p.Region,
		"nodes":        "1",
		"account_file": p.AccountFile,

                "control_count": p.ControlCount,
                "control_type" : p.ControlType,
                "datacenter"   : p.Datacenter,
                "long_name"    : p.LongName,
                "network_ipv4" : p.NetworkIpv4,
                "short_name"   : p.ShortName,
                "worker_count" : p.WorkerCount,
                "worker_type"  : p.WorkerType,
                "ssh.username" : p.SshUserName,
                "ssh.key"      : p.SshKey,

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
	p.Zone = "us-central1-a"
	// TODO: Ask for user input for these.
	p.Project = "REPLACE WITH YOUR ACCESS KEY"
	p.PrivateKeyId = "REPLACE WITH YOUR PRIVATE KEY ID FROM YOUR ACCOUNT FILE"
	p.PrivateKey = "REPLACE WITH YOUR PRIVATE KEY FROM YOUR ACCOUNT FILE"
	p.ClientEmail = "REPLACE WITH YOUR CLIENT EMAIL FROM YOUR ACCOUNT FILE"
	p.ClientId = "REPLACE WITH YOUR CLIENT ID FROM YOUR ACCOUNT FILE"

        p.ControlCount = "3"
        p.ControlType = "n1-standard-1"
        p.Datacenter = "gce"
        p.LongName = "microservices-infrastructure"
        p.NetworkIpv4 = "10.0.0.0/16"
        p.ShortName = "mi"
        p.WorkerCount = "1"
        p.WorkerType = "n1-highcpu-2"
        p.SshKey = "REPLACE WITH THE PUBLIC SSH KEY YOU WANT TO USE"
        p.SshUserName = "REPLACE WITH THE SSH USERNAME YOU WANT TO USE"

}

func (p *Config) PackerVars() map[string]string {
	return map[string]string{
		"gce_creds_file": p.AccountFile,
		"gce_project_id": p.Project,
		"gce_zone":       p.Zone,
	}
}
