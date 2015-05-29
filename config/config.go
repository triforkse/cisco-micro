package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ConfigVars struct {
	Id       string // the cluster id that this config represents
	Provider string
}

type Config struct {
	Config ConfigVars
	Path   string
}

// pure function
func parseJsonBytes(bytes []byte) (*ConfigVars, error) {

	var err error

	config := ConfigVars{}
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

func readConfigsWithPaths(paths []string) []Config {
	var configs = []Config{}

	for _, path := range paths {
		if vars, err := readVars(path); err == nil {
			configs = append(configs, Config{*vars, path})
		}
	}

	return configs
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

	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		paths = append(paths, path)
		return nil
	})

	return paths
}

// impure; file I/O
//
// Read all configs inside given directory path
func ReadConfigs(directory string) []Config {
	paths := allConfigPathsInDirectory(directory)
	return readConfigsWithPaths(paths)
}

func filterConfigs(filterFn func(Config) bool, configs []Config) []Config {
	filtered := []Config{}
	for _, config := range configs {
		if filterFn(config) {
			filtered = append(filtered, config)
		}
	}
	return filtered
}

// Determines which configs the user wants loaded depending on the command line arguments passed in
// Returns the list of matching configurations and the remaining list of command line arguments)
func MatchConfigs(args []string) (matchingConfigs []Config, remainingArgs []string) {
	var clusterId string
	var allClusterIds bool
	var configDir string
	var pred func(Config) bool

	//
	//  Parse command line arguments
	//
	flagSet := flag.NewFlagSet("Config Loader", flag.ContinueOnError)
	flagSet.StringVar(&configDir, "config-dir", "./", `The root directory to look for config files in. Defaults to "./"`)
	flagSet.StringVar(&clusterId, "id", "", "The cluster id to apply the command on")
	flagSet.BoolVar(&allClusterIds, "all", false, "Apply command to all cluster ids")
	flagSet.Parse(args[1:])

	//
	//  Determine which predicate to use
	//
	if !allClusterIds {
		if len(clusterId) == 0 {
			log.Fatal("Missing cluster id")
		} else {
			pred = func(config Config) bool {
				return config.Config.Id == clusterId
			}
		}
	} else {
		pred = nil
	}

	//
	//  Read configs
	//
	configs := ReadConfigs(configDir)
	if pred != nil {
		configs = filterConfigs(pred, configs)
	}
	if len(configs) == 0 {
		log.Fatal("No matching configurations")
	}

	return configs, append([]string{args[0]}, flagSet.Args()...)
}
