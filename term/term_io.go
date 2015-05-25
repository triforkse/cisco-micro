package term
import (
        "cisco/micro/logger"
        "fmt"
)

type AnswerParser func(string) string
type InputAsker func(string, AnswerParser) (string, error)
type InputAskerWithDefault func(string, string, AnswerParser) (string, error)

func AskForInput(question string, answerParser AnswerParser) (string, error) {
        fmt.Println(question)
        var response string
        _, err := fmt.Scanln(&response)
        if err != nil {
                return "", err
        }

        return answerParser(response), nil
}

func AskForInputDefaultAnswer(question string, defaultAnswer string, answerParser AnswerParser) (string, error) {
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
