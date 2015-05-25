package term

import (
  "cisco/micro/logger"
  "fmt"
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

type AnswerParser func(string) string

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

func AskForConfirmation(question string) (bool, error) {
  answer, err := askForInput(question, yesNoParser)
  if err == nil {
    if answer == "yes" {
      return true, nil
    } else if answer == "no" {
      return false, nil
    }
    return AskForConfirmation(question)
  }

  return false, err
}

func AskForAnswer(question string, defaultAnswer string) (string, error) {
  return askForInputDefaultAnswer(question, defaultAnswer, freeParser)
}

func askForInput(question string, answerParser AnswerParser) (string, error) {
  fmt.Println(question)
  var response string
  _, err := fmt.Scanln(&response)
  if err != nil {
    return "", err
  }

  return answerParser(response), nil
}

func askForInputDefaultAnswer(question string, defaultAnswer string, answerParser AnswerParser) (string, error) {
  fmt.Println(question)
  fmt.Printf("Default value is '%s': ", defaultAnswer)
  response := defaultAnswer
  _, err := fmt.Scanln(&response)
  if err != nil {
    return defaultAnswer, err
  }

  logger.Debugf("REsponse %v", response)

  return answerParser(response), nil
}
