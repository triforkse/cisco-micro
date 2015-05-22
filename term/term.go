package term
import "fmt"

var okayResponses = map[string]string{
  "y":"yes",
  "Y":"yes",
  "yes":"yes",
  "Yes":"yes",
  "YES":"yes",
}

var nokayResponses = map[string]string{
  "n": "no",
  "N": "no",
  "no":"no",
  "No":"no",
  "NO":"no",
}

type answerParser func(string) string

func answerParserFn(response string) string {
  if answer, known := okayResponses[response]; known != false {
    return answer
  } else if answer, known := nokayResponses[response]; known != false {
    return answer
  }
  return ""
}

func AskForConfirmation(question string) (bool, error) {
  answer, err := askForInput(question, answerParserFn)
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

func askForInput(question string, answerParserFn answerParser) (string, error) {
  fmt.Println(question)
  var response string
  _, err := fmt.Scanln(&response)
  if err != nil {
    return "", err
  }

  return answerParserFn(response), nil
}

