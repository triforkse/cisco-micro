package config
import (
        "io/ioutil"
        "errors"
        "fmt"
        "path/filepath"
        "encoding/json"
)

func ParseJsonToMap(pathToJsonFile string) (map[string]string, error) {
        bytes, err := ioutil.ReadFile(pathToJsonFile)
        if err != nil {
                return nil, createError(pathToJsonFile)
        } else {
                return unmarshalFileContents(bytes)
        }
}

func WriteToJson(data map[string]string) ([]byte, error) {
        jsonData, err := json.MarshalIndent(data, "", "  ")
        if err != nil {
                return nil, errors.New(fmt.Sprintf("Could not write map to json. %v", err))
        }

        return jsonData, nil
}

func WriteToFile(filePath string, contents []byte) error {
        err := ioutil.WriteFile(filePath, contents, 0644)
        if err != nil {
                return errors.New(fmt.Sprintf("Could not write to %s due to error %v", filePath, err))
        }

        return nil
}

func createError(pathToJsonFile string) error {
        fullPath, _ := filepath.Abs(pathToJsonFile)
        return errors.New(fmt.Sprintf("Failed to load file %s", fullPath))
}

func unmarshalFileContents(contents []byte) (map[string]string, error) {
        var vars map[string]string
        err := json.Unmarshal(contents, &vars)
        if err != nil {
                return vars, err
        }
        return vars, nil
}
