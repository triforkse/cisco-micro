package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type ConfigVars struct {
	id       string
	provider string
}

type Config struct {
	config ConfigVars
	path   string
}

func readVars(path string) (*ConfigVars, error) {

	var bytes []byte
	var err error

	bytes, _ = ioutil.ReadFile(path)

	var config ConfigVars
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Printf("Can't read config %s", path)
		return nil, err
	}

	return &config, nil
}

func isConfigPath(path string) bool {
	return strings.HasSuffix(path, ".json")
}

func allConfigPathsInDirectory(dirPath string) []string {

	var configPaths = []string{}

	files, _ := ioutil.ReadDir(dirPath)
	for _, f := range files {
		name := f.Name()
		if isConfigPath(name) {
			configPaths = append(configPaths, name)
		}
	}

	return configPaths
}

func readConfigs(directory string) []Config {
	var configs = []Config{}
	var paths = []string{}

	paths = allConfigPathsInDirectory(directory)
	for _, path := range paths {
		if vars, err := readVars(path); err == nil {
			configs = append(configs, Config{*vars, path})
		}
	}

	return configs
}
