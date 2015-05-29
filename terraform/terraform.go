package terraform

import (
        "path/filepath"



        "fmt"
        "cisco/micro/config"
)

type CommandContext struct {
        Command string
        Vars map[string]string
        Flags map[string]string
}

type TerraformContext struct {
        Context CommandContext
        Config config.Config
}

func (cmd *CommandContext) ToList() []string {
        args := []string{cmd.Command}
        args = append(args, createVariableList(cmd.Vars)...)
        args = append(args, createFlagsList(cmd.Flags)...)

        return args
}

func computeTerraformStatePath(config config.Config) string {
        return filepath.Join(".micro", config.Config.Id +".tfstate")
}

func computeTerraformConfigPath(config config.Config) string {
      // return filepath.Join(configFileLocation, config.Config.Id()))
        return "TODO"
}

func (cmd *TerraformContext) ToList() []string {
        cmd.Context.Flags["state"] = computeTerraformStatePath(cmd.Config)
        commandList := append(cmd.Context.ToList(), computeTerraformConfigPath(cmd.Config))

        return append([]string{"echo", "terraform"}, commandList...)
}

func createVariableList(vars map[string]string) []string {
        var addedArgs []string
        for key, value := range vars {
                addedArgs = append(addedArgs, "-var", fmt.Sprintf("%s=%s", key, value))
        }
        return addedArgs
}

func createFlagsList(flags map[string]string) []string {
        var addedFlags []string
        for key, value := range flags {
                if len(value) == 0 {
                        addedFlags = append(addedFlags, fmt.Sprintf("-%s", key))
                } else {
                        addedFlags = append(addedFlags, fmt.Sprintf("-%s=%s", key, value))
                }
        }
        return addedFlags
}

