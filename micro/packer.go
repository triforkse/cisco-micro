package main

import (
//	"cisco/micro/logger"
//	"cisco/micro/provider"
//	"cisco/micro/util/executil"
//	"os"
)

//func packerCmd(config provider.Provider) error {
//
//	return config.Run(func() error {
//
//		// Temporarily change to packer directory,
//		// since the packer files use relative paths.
//
//		pwd, _ := os.Getwd()
//		cwdErr := os.Chdir(".micro/src/img-build/packer")
//		defer os.Chdir(pwd)
//
//		if cwdErr != nil {
//			return cwdErr
//		}
//
//		args := []string{"build", "-only=" + config.ProviderId()}
//		args = append(args, provider.VarList(config.PackerVars())...)
//		args = append(args, "all-in-one.json")
//
//		logger.Debugf("packer", args)
//		cmd := executil.Command("packer", args...)
//
//		return cmd.Run()
//	})
//}
