package term
import (
        "fmt"
        "strings"
        "cisco/micro/util/json"
        "flag"
        "errors"
)

//
//  Command line parsing
//


// https://lawlessguy.wordpress.com/2013/07/23/filling-a-slice-using-command-line-flags-in-go-golang/
type varmap map[string]string

func (self *varmap) Set(value string) error {
        k, v := splitToKeyValue(value, "=")
        if len(k) == 0 {
                return errors.New("Couldn't parse variable " + value)
        }
        (*self)[k] = v
        return nil
}

func (self *varmap) String() string {

        return "TODO varmap"
}

func parseCmdLineVariables2(args []string) (varmap, []string) {
        var myvars varmap = varmap{}
        flagSet := flag.NewFlagSet("Command line variables", flag.ContinueOnError)
        flagSet.Var(&myvars, "var", "Specify a variable")
        flagSet.Parse(args)

        return myvars, flagSet.Args()
}



//
// VariableParser
//
type VariableParser struct {
        AskFunction func(string) (string, error)
        WriteToFile func(string, []byte) error
}

func (self *VariableParser) GatherVariablesFromFile(requiredVars []string, filePath string) map[string]string {
        suppliedVars, err := json.ParseJsonToMap(filePath)
        if err != nil {
                panic(fmt.Sprintf("Could not parse json file due to: %v", err))
        }
        return self.complementWithRequiredVariables(suppliedVars, requiredVars)
}

func (self *VariableParser) GatherVariablesFromCommandLineArgs(requiredVars []string, commandLineArgs []string) (varmap, []string) {
        suppliedVars, remaining := parseCmdLineVariables2(commandLineArgs)
        return self.complementWithRequiredVariables(suppliedVars, requiredVars), remaining
}


func (self *VariableParser) WriteVariablesToFile(variables map[string]string, filePath string) {
        if jsonData, err := json.WriteToJson(variables); err == nil {
                err = self.WriteToFile(filePath, jsonData)
                if err != nil {
                        panic(fmt.Sprintf("Could not write json to file due to: %v", err))
                }
        } else {
                panic(fmt.Sprintf("Could not write json to file due to: %v", err))
        }
}

func (self *VariableParser) complementWithRequiredVariables(suppliedVars map[string]string, requiredVars []string) map[string]string {
        vars := suppliedVars
        for _, requiredVar := range requiredVars {
                vars[requiredVar] = self.extractValue(suppliedVars, requiredVar)
        }
        return vars
}

func (self *VariableParser) extractValue(suppliedVars map[string]string, requiredVar string) string {
        if suppliedValue, known := suppliedVars[requiredVar]; !known || len(suppliedValue) == 0 {
                return self.askUserToComplementVar(requiredVar)
        } else {
                return suppliedValue
        }
}

func (self *VariableParser) askUserToComplementVar(requiredVar string) string {
        answer, err := self.AskFunction(fmt.Sprintf("You must supply a value for variable '%s'", requiredVar))
        if err != nil || len(answer) == 0 {
                return self.askUserToComplementVar(requiredVar)
        }
        return answer
}

func splitToKeyValue(kv string, splitter string) (string, string) {
        parts := strings.Split(kv, splitter)
        if len(parts) == 2 {
                return parts[0], parts[1]
        }
        return "", ""
}
