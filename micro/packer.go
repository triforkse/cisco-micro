package main

import (
	"cisco/micro/provider"
	"os"
	"os/exec"
)

func packerCmd(command string, provider provider.Provider) error {

	args := []string{command}
	args = append(args, "all-in-one.json")

	//Temporarily change to packer directory
	pwd, _ := os.Getwd()
	cwd_err := os.Chdir(".micro/src/img-build/packer")
	defer os.Chdir(pwd)

	if cwd_err != nil {
		return cwd_err
	}

	cmd := exec.Command("packer", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
