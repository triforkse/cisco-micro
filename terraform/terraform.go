package terraform

import (
        "path/filepath"
        "cisco/micro/logger"
        "cisco/micro/provider"
        "cisco/micro/util/executil"
        "reflect"
        "cisco/micro/term"
        "log"
        "fmt"
)

func TerraformCmd(command string, config provider.Provider, configFileLocation string) {

        config.Run(func() error {
                args := []string{command}

                // Determine if we have an old tfstate file we need to load.
                args = append(args, "-state="+filepath.Join(".micro", config.ConfigId()+".tfstate"))

                // Pass in the arguments
                terraformVars := gatherVars(config, term.AskForInputDefaultAnswer)
                args = append(args, provider.VarList(terraformVars)...)
                args = append(args, "-var", "deployment_id="+config.ConfigId())

                // Tell it what template to use based on the provider.
                args = append(args, filepath.Join(configFileLocation, config.ProviderId()))

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

func gatherVars(config provider.Provider, askFn term.InputAskerWithDefault) map[string]string {
        vars := config.TerraformVars()
        numberOfElements := reflect.TypeOf(config).Elem().NumField()
        for i := 0; i < numberOfElements; i++ {

                field := reflect.TypeOf(config).Elem().FieldByIndex([]int{i})
                configName := field.Tag.Get("json")

                question := fmt.Sprintf("Enter a new value for property '%s' to override the default", configName)
                value, err := provider.ComplementVars(config, field.Name, question, askFn)

                if err != nil {
                        log.Fatalf("Could not retrive value for '%s' due to error '%v'", field.Name, err)
                }

                vars[configName] = value
        }

        return vars
}

