package provider

import (
        "fmt"
        "flag"
        "cisco/micro/config"
        "cisco/micro/term"
        "cisco/micro/util/json"
        "cisco/micro/terraform"
        "cisco/micro/util/executil"
        "cisco/micro/logger"
)

var requiredVars []string = []string{"deployment_id", "region", "secret_key", "access_key"}

func AwsBuild(cfg config.Config, args []string) int{
        fmt.Println("AWS build")
        return 0
}

func findVar(varname string, vars map[string]string) string {
        value, known := vars[varname]
        if !known {
                for {
                        answer, err := term.AskForInput(fmt.Sprintf("The variable %s is required. Please suply a value: ", varname))
                        if err != nil {
                                continue
                        }
                        return answer
                }
        }

        return value
}

func createAskForInputValueAndLookInConfigClosureFactory(vars map[string]string) (func(string) string) {
        return func(v string) {
                return findVar(v, vars)
        }
}

func AwsApply(cfg config.Config, args []string) int {


        parsedConfig := map[string]string{} // TODO parse it

        variableProvider := term.VariableParser{AskFunction: createAskForInputValueAndLookInConfigClosureFactory(parsedConfig), WriteToFile: json.WriteToFile}
        gatheredVars, remainingArgs := variableProvider.GatherVariablesFromCommandLineArgs(requiredVars, args)
        variableProvider.WriteVariablesToFile(gatheredVars, cfg.Path)

        //Todo cmd flags
        flags := map[string]string{}
        flagSet := flag.NewFlagSet("Apply command", flag.ContinueOnError)
        // TODO add flags
        flagSet.Parse(remainingArgs)

        if len(flagSet.Args()) > 0 {
                // TODO crash and burn
        }


        cmdCtx := terraform.CommandContext{"apply", gatheredVars, flags}
        terraformCtx := terraform.TerraformContext{cmdCtx, cfg}

        commandList := terraformCtx.ToList()

        logger.Debugf("%s\n", commandList)
        executil.CommandList(commandList...).Run()

        return  0
}


func getAwsDispatchTable() map[string] CommandFunction{
        return map[string] CommandFunction{
                "build": AwsBuild,
                "apply": AwsApply,
        }
}
