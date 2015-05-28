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

// pure function
func parseJsonBytes(bytes []byte) (*ConfigVars, error) {

	var err error
	var config ConfigVars

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// pure function
func isConfigPath(path string) bool {
	return strings.HasSuffix(path, ".json")
}

// pure function
func validConfigPaths(paths []string) []string {
	var configPaths = []string{}

	for _, f := range paths {
		if isConfigPath(f) {
			configPaths = append(configPaths, f)
		}
	}

	return configPaths
}

// impure; file I/O
func readVars(path string) (*ConfigVars, error) {

	var bytes []byte

	bytes, _ = ioutil.ReadFile(path)
	vars, err := parseJsonBytes(bytes)
	if err != nil {
		log.Printf("Can't read config %s", path)
	}

	return vars, err
}

// impure; file I/O
func allConfigPathsInDirectory(dirPath string) []string {
	return validConfigPaths(allFilePathsInDirectory(dirPath))
}

// impure; file I/O
func allFilePathsInDirectory(dirPath string) []string {
	var paths = []string{}

	files, _ := ioutil.ReadDir(dirPath)
	for _, f := range files {
		name := f.Name()
		paths = append(paths, name)
	}

	return paths
}

// impure; file I/O
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
