package term
import (
        "fmt"
        "strings"
        "cisco/micro/util/config"
)


type VariableParser struct {
        askFunction func(string) (string, error)
        writeToFile func(string, []byte) error
}

func (self *VariableParser) GatherVariablesFromFile(requiredVars []string, filePath string) map[string]string {
        suppliedVars, err := config.ParseJsonToMap(filePath)
        if err != nil {
                panic(fmt.Sprintf("Could not parse json file due to: %v", err))
        }
        return self.complementWithRequiredVariables(suppliedVars, requiredVars)
}

func (self *VariableParser) GatherVariablesFromCommandLineArgs(requiredVars []string, commandLineArgs []string) map[string]string {
        suppliedVars := self.parseCmdLineVariables(commandLineArgs)
        return self.complementWithRequiredVariables(suppliedVars, requiredVars)
}

func (self *VariableParser) WriteVariablesToFile(variables map[string]string, filePath string) {
        if jsonData, err := config.WriteToJson(variables); err == nil {
                err = self.writeToFile(filePath, jsonData)
                if err != nil {
                        panic(fmt.Sprintf("Could not write json to file due to: %v", err))
                }
        } else {
                panic(fmt.Sprintf("Could not write json to file due to: %v", err))
        }
}

func (self *VariableParser) parseCmdLineVariables(args []string) map[string]string {
        vars := map[string]string{}
        for _, arg := range args {
                if varPart, kvPart := self.splitToKeyValue(arg, " "); varPart == "-var" {
                        key, value := self.splitToKeyValue(kvPart, "=")
                        vars[key] = value
                }
        }

        return vars
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
        answer, err := self.askFunction(fmt.Sprintf("You must supply a value for variable '%s'", requiredVar))
        if err != nil || len(answer) == 0 {
                return self.askUserToComplementVar(requiredVar)
        }
        return answer
}

func (self *VariableParser) splitToKeyValue(kv string, splitter string) (string, string) {
        parts := strings.Split(kv, splitter)
        if len(parts) == 2 {
                return parts[0], parts[1]
        }
        return "", ""
}
