package gce

type AccountFile struct {
	PrivateKeyId string `json:"private_key_id"`
	PrivateKey   string `json:"private_key"`
	ClientEmail  string `json:"client_email"`
	ClientId     string `json:"client_id"`
}

func accountFileFromConfig(config Config) AccountFile {
	return AccountFile{
		PrivateKeyId: config.PrivateKeyId,
		PrivateKey:   config.PrivateKey,
		ClientEmail:  config.ClientEmail,
		ClientId:     config.ClientId,
	}
}
