package term
import "fmt"

var okayResponses = map[string]string{
  "y":"yes",
  "Y":"yes",
  "yes":"yes",
  "Yes":"yes",
  "YES":"yes",
}

type answerParser func(string) string

func AskForConfirmation(question string) (bool, error) {

  answerParser := func(response string) string {
    if answer, known := okayResponses[response]; known != false {
      return answer
    }
    return "no"
  }

  answer, err := askForInput(question, answerParser)
  if err == nil {
    if answer == "yes" {
      return true, nil
    }
    return false, nil
  }

  return false, err
}

func askForInput(question string, answerParserFn answerParser) (string, error) {

  fmt.Println(question)
  var response string
  _, err := fmt.Scanln(&response)
  if err != nil {
    return "", err
  }

  return answerParserFn(response), nil
}

