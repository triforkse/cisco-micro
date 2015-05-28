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

func TestParseCommandLineArguments(t *testing.T) {
        result := parseCmdLineVariables([]string{"-var k1=v1", "-var k2=v2", "-notvar k3=v3"})

        if value, _ := result["k1"]; value != "v1" {
                t.Errorf("Expected k1 to have value v1, got %s", value)
        }

        if value, _ := result["k2"]; value != "v2" {
                t.Errorf("Expected k1 to have value v2, got %s", value)
        }

        if _, known := result["k3"]; known {
                t.Error("Did not expect k3 to be known")
        }
}

func TestExtractValue(t *testing.T) {
        suppliedVars := map[string]string{"key1":"value1", "key2": "value2"}

        result := extractValue(suppliedVars, "key1", func(answer string) string { return "complementedValue"})
        if result != "value1" {
                t.Errorf("Expected key1 to be value1 , got %v", result)
        }

        result = extractValue(suppliedVars, "key2", func(answer string) string { return "complementedValue"})
        if result != "value2"{
                t.Errorf("Expected key2 to be value2 , got %v", result)
        }

        result = extractValue(suppliedVars, "key3", func(answer string) string { return "complementedValue"})
        if result != "complementedValue"{
                t.Errorf("Expected key3 to be complementedValue , got %v", result)
        }
}

