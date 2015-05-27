package term
import (
        "cisco/micro/logger"
        "fmt"
        "bufio"
        "os"
        "strings"
)

type AnswerParser func(string) string
type InputAsker func(string) (string, error)
type InputAskerWithDefault func(string, string) (string, error)

func AskForInput(question string) (string, error) {
        fmt.Println(question)
        response := ""

        reader := bufio.NewReader(os.Stdin)
        line, err := reader.ReadString('\n')

        if err != nil {
                return "", err
        }

        response = strings.TrimRight(line, "\r\n")
        return response, nil
}

func AskForInputDefaultAnswer(question string, defaultAnswer string) (string, error) {
        fmt.Println(question)
        fmt.Printf("Default value is '%s': ", defaultAnswer)
        response := defaultAnswer

        reader := bufio.NewReader(os.Stdin)
        line, err := reader.ReadString('\n')

        if err != nil {
                return defaultAnswer, err
        }

        response = strings.TrimRight(line, "\r\n")
        logger.Debugf("Response %v", response)

        return response, nil
}
