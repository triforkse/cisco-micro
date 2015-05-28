package term

import (
        "testing"
        "errors"
)

func TestYesNoParserExpectingYesForAllYesLikeAnswers(t *testing.T) {
        expectedAnswer := "yes"
        for _, element := range []string{"YES", "yes", "Yes", "y", "Y"} {
                if answer := yesNoParser(element); answer != expectedAnswer {
                        t.Errorf("Expected 'yes', got '%s'", answer)
                }
        }
}

func TestYesNoParserExpectingNoForAllNoLikeAnswers(t *testing.T) {
        expectedAnswer := "no"
        for _, element := range []string{"NO", "no", "No", "n", "N"} {
                if answer := yesNoParser(element); answer != expectedAnswer {
                        t.Errorf("Expected 'no', got '%s'", answer)
                }
        }
}

func TestAskForConfirmationExpectingTrueWhenAnswerIsYes(t *testing.T) {
        answer, err := AskForConfirmation("Test question", func(question string) (string, error) {
                return "yes", nil
        })

        if err != nil {
                t.Errorf("Expected true, got %v", err)
        }

        if !answer {
                t.Errorf("Expected true, got %v", answer)
        }
}

func TestAskForConfirmationExpectingFalseWhenAnswerIsNo(t *testing.T) {
        answer, err := AskForConfirmation("Test question", func(question string) (string, error) {
                return "no", nil
        })

        if err != nil {
                t.Errorf("Expected false, got %v", err)
        }

        if answer {
                t.Errorf("Expected false, got %v", answer)
        }
}

func TestAskForConfirmationExpectingFalseOnError(t *testing.T) {
        answer, err := AskForConfirmation("Test question", func(question string) (string, error) {
                return "", errors.New("")
        })

        if err == nil {
                t.Error("Expected false and error")
        }

        if answer {
                t.Errorf("Expected false and error, got %v and %v", answer, err)
        }
}


