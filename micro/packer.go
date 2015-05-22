package main

import (
	"os"
	"cisco/micro/util/executil"

	"cisco/micro/provider"
)

func packerCmd(command string, provider provider.Provider) error {

	args := []string{command, "all-in-one.json"}

	//Temporarily change to packer directory
	pwd, _ := os.Getwd()
	cwdErr := os.Chdir(".micro/src/img-build/packer")
	defer os.Chdir(pwd)

	if cwdErr != nil {
		return cwdErr
	}

	cmd := executil.Command("packer", args...)

	return cmd.Run()
}
