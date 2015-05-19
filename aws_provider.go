package main


type AWSProvider struct {
	SecretKey string `json:"secret_key"`
	AccessKey string `json:"access_key"`
	Region    string
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
