package term

import (
        "strings"
        "fmt"
        "log"
)

var okayResponses = map[string]string{
        "y":   "yes",
        "Y":   "yes",
        "yes": "yes",
        "Yes": "yes",
        "YES": "yes",
}

var nokayResponses = map[string]string{
        "n":  "no",
        "N":  "no",
        "no": "no",
        "No": "no",
        "NO": "no",
}

func yesNoParser(response string) string {
        if answer, known := okayResponses[response]; known != false {
                return answer
        } else if answer, known := nokayResponses[response]; known != false {
                return answer
        }
        return ""
}

func AskForConfirmation(question string, askForInput InputAsker) (bool, error) {
        answer, err := askForInput(question)
        answer = yesNoParser(answer)
        if err == nil {
                if answer == "yes" {
                        return true, nil
                } else if answer == "no" {
                        return false, nil
                }
                return AskForConfirmation(question, askForInput)
        }

        return false, err
}

func AskForAnswer(question string, defaultAnswer string, askForInputDefaultAnswer InputAskerWithDefault) (string, error) {
        response, err := askForInputDefaultAnswer(question, defaultAnswer)

        if err != nil {
                return "", err
        }

        return response, nil
}

func GatherVars(requiredVars []string, args []string) map[string]string {
        suppliedVars := parseCmdLineVariables(args)
        complementedVars := complementWithRequiredVariables(suppliedVars, requiredVars)
        return complementedVars
}

func parseCmdLineVariables(args []string) (map[string]string) {
        vars := map[string]string{}
        for _, arg := range args {
                if varPart, kvPart := splitToKeyValue(arg, " "); varPart == "-var" {
                        key, value := splitToKeyValue(kvPart, "=")
                        vars[key] = value
                }
        }

        return vars
}

func complementWithRequiredVariables(suppliedVars map[string]string, requiredVars []string) map[string]string {
        vars := suppliedVars
        for _, requiredVar := range requiredVars {
                vars[requiredVar] = extractValue(suppliedVars, requiredVar, askUserToComplementVar)
        }
        return vars
}

func extractValue(suppliedVars map[string]string, requiredVar string, askFn func(string) string) string {
        if suppliedValue, known := suppliedVars[requiredVar]; !known || len(suppliedValue) == 0 {
                return askFn(requiredVar)
        } else {
                return suppliedValue
        }
}

func askUserToComplementVar(requiredVar string) string {
        answer, err := AskForInput(fmt.Sprintf("You must supply a value for variable '%s'", requiredVar))
        if err != nil || len(answer) == 0 {
                log.Fatalf("Could not get answer for required value '%s' due to error '%v'", requiredVar, err)
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
