package term

import (
        "testing"
        "errors"
)

func TestParseCommandLineVariables(t *testing.T) {
        variableParser := VariableParser{askFunction: func(question string) (string, error) {
                return "complementedValue", nil
        }}
        result := variableParser.parseCmdLineVariables([]string{"-var k1=v1", "-var k2=v2", "-notvar k3=v3", "-var -var"})

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
        variableParser := VariableParser{askFunction: func(question string) (string, error) {
                return "complementedValue", nil
        }}
        suppliedVars := map[string]string{"key1":"value1", "key2": "value2"}

        result := variableParser.extractValue(suppliedVars, "key1")
        if result != "value1" {
                t.Errorf("Expected key1 to be value1 , got %v", result)
        }

        result = variableParser.extractValue(suppliedVars, "key2")
        if result != "value2" {
                t.Errorf("Expected key2 to be value2 , got %v", result)
        }

        result = variableParser.extractValue(suppliedVars, "key3")
        if result != "complementedValue" {
                t.Errorf("Expected key3 to be complementedValue , got %v", result)
        }
}

func TestGatherVariablesWithRequiredSupplied(t *testing.T) {
        variableParser := VariableParser{askFunction: func(question string) (string, error) {
                return "complementedValue", nil
        }}

        result := variableParser.GatherVariables([]string{"key1"}, []string{"-var key1=value1"})

        if value, known := result["key1"]; !known || value != "value1" {
                t.Errorf("Expected value1 got '%v'", value)
        }
}

func TestGatherVariablesWithRequiredAndAdditionalSupplied(t *testing.T) {
        variableParser := VariableParser{askFunction: func(question string) (string, error) {
                return "complementedValue", nil
        }}

        result := variableParser.GatherVariables([]string{"key1"}, []string{"-var key1=value1", "-var key2=value2"})

        if value, known := result["key1"]; !known || value != "value1" {
                t.Errorf("Expected value1 got '%v'", value)
        }

        if value, known := result["key2"]; !known || value != "value2" {
                t.Errorf("Expected value2 got '%v'", value)
        }
}

func TestGatherVariablesWithRequiredMissing(t *testing.T) {
        variableParser := VariableParser{askFunction: func(question string) (string, error) {
                return "complementedValue", nil
        }}

        result := variableParser.GatherVariables([]string{"key1"}, []string{"-var key2=value2"})

        if value, known := result["key1"]; !known || value != "complementedValue" {
                t.Errorf("Expected complementedValue got '%v'", value)
        }

        if value, known := result["key2"]; !known || value != "value2" {
                t.Errorf("Expected value2 got '%v'", value)
        }
}


func TestGatherVariablesWithNonKnownOptions(t *testing.T) {
        variableParser := VariableParser{askFunction: func(question string) (string, error) {
                return "complementedValue", nil
        }}

        result := variableParser.GatherVariables([]string{"key1"}, []string{"-var key1=value1", "-unknown test:test"})

        if value, known := result["key1"]; !known || value != "value1" {
                t.Errorf("Expected value1 got '%v'", value)
        }
}

func TestGatherVariablesWithWrongFormat(t *testing.T) {
        variableParser := VariableParser{askFunction: func(question string) (string, error) {
                return "complementedValue", nil
        }}

        result := variableParser.GatherVariables([]string{}, []string{"-var key1value1"})

        if _, known := result["key1valu1"]; known {
                t.Errorf("Expected to be unknown")
        }
}

func TestGatherVariablesAskRetriesUntilValidAnswer(t *testing.T) {
        callCount := 0
        askFunction := func(question string) (string, error) {
                switch callCount {
                case 0:
                        callCount++
                        return "", errors.New("Error")
                case 1:
                        callCount++
                        return "", nil
                default:
                        callCount++
                        return "complementedValue", nil
                }
        }
        variableParser := VariableParser{askFunction: askFunction}

        result := variableParser.GatherVariables([]string{"key1"}, []string{})

        if value, _ := result["key1"]; value != "complementedValue" {
                t.Errorf("Expected 'complementedValue', got '%s'", value)
        }

        if callCount != 3 {
                t.Errorf("Expected askFunctionMock to have been called 3 times but was called %d times", callCount)
        }
}

