package term

import (
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

func freeParser(response string) string {
	return response
}

func AskForConfirmation(question string, askForInput InputAsker) (bool, error) {
	answer, err := askForInput(question, yesNoParser)
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
	return askForInputDefaultAnswer(question, defaultAnswer, freeParser)
}
