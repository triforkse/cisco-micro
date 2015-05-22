package main

import (
	"os"
	"cisco/micro/util/executil"
	"cisco/micro/provider"
  "cisco/micro/logger"
)

func packerCmd(config provider.Provider) error {

	// Temporarily change to packer directory
	pwd, _ := os.Getwd()
	cwdErr := os.Chdir(".micro/src/img-build/packer")
	defer os.Chdir(pwd)

	if cwdErr != nil {
		return cwdErr
	}

  args := []string{"build", "-only=" + config.ProviderId()}
  args = append(args, provider.VarList(config.PackerVars())...)
  args = append(args, "all-in-one.json")

  logger.Debugf("packer", args)
	cmd := executil.Command("packer", args...)

	return cmd.Run()
}
