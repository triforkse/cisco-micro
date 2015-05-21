package aws

import (
  "cisco/micro/util/strutil"
)

type Provider struct {
	Id        string `json:"id"`
	Provider  string `json:"provider"`
	SecretKey string `json:"secret_key"`
	AccessKey string `json:"access_key"`
	Region    string `json:"region"`
}

func (p *Provider) TerraformVars() map[string]string {
	return map[string]string{
		"secret_key": p.SecretKey,
		"access_key": p.AccessKey,
		"region": p.Region,
	}
}

func (p *Provider) Prepare() { }

func (p *Provider) Cleanup() { }

func (p *Provider) ConfigId() string {
	return p.Id
}

func (p *Provider) ProviderId() string {
	return "aws"
}

func (p *Provider) Populate() {
	p.Id = strutil.Random(16)
	p.Provider = "aws"
	p.SecretKey = "REPLACE WITH YOUR SECRET KEY"
    p.AccessKey = "REPLACE WITH YOUR ACCESS KEY"
	p.Region = "eu-west-1"
}
