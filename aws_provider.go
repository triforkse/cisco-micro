package main


type AWSProvider struct {
	Id        string `json:"id"`
	Provider  string `json:"provider"`
	SecretKey string `json:"secret_key"`
	AccessKey string `json:"access_key"`
	Region    string `json:"region"`
}

func (p *AWSProvider) terraformVars() map[string]string {
	return map[string]string{
		"secret_key": p.SecretKey,
		"access_key": p.AccessKey,
		"region": p.Region,
	}
}

func (p *AWSProvider) prepare() { }

func (p *AWSProvider) cleanup() { }

func (p *AWSProvider) configId() string {
	return p.Id
}

func (p *AWSProvider) providerId() string {
	return "aws"
}

func (p *AWSProvider) populate() {
	p.Id = randStr(16)
	p.Provider = "aws"
	p.SecretKey = "REPLACE WITH YOUR SECRET KEY"
  p.AccessKey = "REPLACE WITH YOUR ACCESS KEY"
  p.Region = "eu-west-1"
}
