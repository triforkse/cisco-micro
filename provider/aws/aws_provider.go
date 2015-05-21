package aws

import (
	"cisco/micro/util/strutil"
)

type Config struct {
	Id        string `json:"id"`
	Provider  string `json:"provider"`
	SecretKey string `json:"secret_key"`
	AccessKey string `json:"access_key"`
	Region    string `json:"region"`
}

func (p *Config) TerraformVars() map[string]string {
	return map[string]string{
		"secret_key": p.SecretKey,
		"access_key": p.AccessKey,
		"region":     p.Region,
	}
}

func (p *Config) Prepare() {}

func (p *Config) Cleanup() {}

func (p *Config) ConfigId() string {
	return p.Id
}

func (p *Config) ProviderId() string {
	return "aws"
}

func (p *Config) Populate() {
	p.Id = strutil.Random(16)
	p.Provider = "aws"
	p.SecretKey = "REPLACE WITH YOUR SECRET KEY"
	p.AccessKey = "REPLACE WITH YOUR ACCESS KEY"
	p.Region = "eu-west-1"
}
