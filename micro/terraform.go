package main

import (
	"os"
	"os/exec"
	"path/filepath"

	"cisco/micro/logger"
	"cisco/micro/provider"
)

func terraformCmd(command string, provider provider.Provider) {

	provider.Prepare()
	defer provider.Cleanup()

	args := []string{command}

	// Determine if we have an old tfstate file we need to load.
	args = append(args, "-state="+filepath.Join(".micro", provider.ConfigId()+".tfstate"))

	// Pass in the arguments
	for k, v := range provider.TerraformVars() {
		args = append(args, "-var", k+"="+v)
	}
	args = append(args, "-var", "deployment_id="+provider.ConfigId())

	// Tell it what template to use based on the provider.
	args = append(args, filepath.Join("templates", provider.ProviderId()))

	logger.Debugf("terraform %+v", args)

	// Run Terraform
	cmd := exec.Command("terraform", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if cmd.Run() != nil {
		os.Exit(1)
	}

	logger.PrintTable("Cluster Properties", map[string]string{
		"Type": provider.ProviderId(),
		"ID":   provider.ConfigId(),
	})
}
