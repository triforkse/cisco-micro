package main

import (
	"path/filepath"

	"cisco/micro/logger"
	"cisco/micro/provider"
	"cisco/micro/util/executil"
)

func terraformCmd(command string, config provider.Provider) {

	config.Run(func() error {
		args := []string{command}

		// Determine if we have an old tfstate file we need to load.
		args = append(args, "-state="+filepath.Join(".micro", config.ConfigId()+".tfstate"))

		// Pass in the arguments
		args = append(args, provider.VarList(config.TerraformVars())...)
		args = append(args, "-var", "deployment_id="+config.ConfigId())

		// Tell it what template to use based on the provider.
		args = append(args, filepath.Join("templates", config.ProviderId()))

		logger.Debugf("terraform %+v", args)

		// Run Terraform
		cmd := executil.Command("terraform", args...)

		err := cmd.Run()

		logger.PrintTable("Cluster Properties", map[string]string{
			"Type": config.ProviderId(),
			"ID":   config.ConfigId(),
		})

		return err
	})
}
