package aws

import (
	"cisco/micro/util/strutil"
)

type Config struct {
	Id          string `json:"id"`
	Provider    string `json:"provider"`
	SecretKey   string `json:"secret_key"`
	AccessKey   string `json:"access_key"`
	Region      string `json:"region"`
	SourceAMI   string `json:"source_ami"`
	SSHUsername string `json:"ssh_username"`
}

func (p *Config) TerraformVars() map[string]string {
	return map[string]string{
		"secret_key": p.SecretKey,
		"access_key": p.AccessKey,
		"region":     p.Region,
	}
}

// Terraform

func (p *Config) Run(action func() error) error {
	return action()
}

func (p *Config) ConfigId() string {
	return p.Id
}

func (p *Config) ProviderId() string {
	return "aws"
}

func (p *Config) Populate() {
	p.Id = strutil.Random(16)
	p.Provider = "aws"
	p.Region = "eu-west-1"
	p.SourceAMI = "ami-10e14667"
	p.SSHUsername = "centos"
	// TODO: Ask the user for these on micro init
	p.SecretKey = "REPLACE WITH YOUR SECRET KEY"
	p.AccessKey = "REPLACE WITH YOUR ACCESS KEY"
}

// Packer

func (p *Config) PackerVars() map[string]string {
	return map[string]string{
		"aws_secret_key":   p.SecretKey,
		"aws_access_key":   p.AccessKey,
		"aws_source_ami":   p.SourceAMI,
		"aws_region":       p.Region,
		"aws_ssh_username": p.SSHUsername,
	}
}
