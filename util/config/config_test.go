package config
import (
        "testing"
        "os"
)

const outputTestFile string = "testdata/output-config.json"

func tearDown() {
        if _, err := os.Stat(outputTestFile); err == nil {
                os.Remove(outputTestFile)
        }
}

func TestMain(m *testing.M) {
        status := m.Run()
        tearDown()
        os.Exit(status)
}

func TestParseJsonToMap(t *testing.T) {
        filePath := "testdata/test-config.json"

        result, _ := ParseJsonToMap(filePath)

        if val, _ := result["id"]; val != "test-aws-123" {
                t.Errorf("Expected id to have value test-aws-123, got %s", val)
        }

        if val, _ := result["provider"]; val != "aws" {
                t.Errorf("Expected id to have value aws, got %s", val)
        }

        if val, _ := result["secret_key"]; val != "TEST SECRET KEY" {
                t.Errorf("Expected id to have value TEST SECRET KEY, got %s", val)
        }

        if val, _ := result["access_key"]; val != "TEST ACCESS KEY" {
                t.Errorf("Expected id to have value TEST ACCESS KEY, got %s", val)
        }

        if val, _ := result["region"]; val != "eu-west-1" {
                t.Errorf("Expected id to have value eu-west-1, got %s", val)
        }
}

func TestWriteToJson(t *testing.T) {
        result, _ :=  WriteToJson(map[string]string{"key1":"value1", "key2":"value2"})

        jsonMap, _ := unmarshalFileContents(result)

        if value, known := jsonMap["key1"]; !known || value != "value1" {
                t.Errorf("Expected valid json")
        }
}

func TestWriteToFile(t *testing.T) {
        if err := WriteToFile(outputTestFile, []byte{65}); err != nil {
                t.Errorf("Expected file to created, got error %v", err)
        }

        if _, err := os.Stat(outputTestFile); err != nil {
                t.Errorf("Expected file to created, got error %v", err)
        }
}
