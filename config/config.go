package config

import (
        "cisco/micro/logger"
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


func (v *ConfigVars) AreValid() bool {
        return v.Id != "" &&  v.Provider != ""
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
                if vars, err := readVars(path); err == nil && vars.AreValid() {
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
        var clusterIds string
        var acceptAllClusters bool
        var configDir string
        var pred func(Config) bool

        //
        //  Parse command line arguments
        //

        flagSet := flag.NewFlagSet("Config Loader", flag.ContinueOnError)
        flagSet.SetOutput(ioutil.Discard)
        flagSet.StringVar(&configDir, "config-dir", "./", `The root directory to look for config files in. Defaults to "./"`)
        flagSet.StringVar(&clusterIds, "id", "", "The cluster id to apply the command on")
        flagSet.BoolVar(&acceptAllClusters, "all", false, "Apply command to all cluster ids")
        flagSet.Parse(args)

        //
        //  Determine which predicate to use
        //
        if acceptAllClusters {
                pred = func(config Config) bool {
                        return true
                }
        } else {
                if len(clusterIds) == 0 {
                        logger.Errorf("Missing cluster id(s) or all flag. See help for more infos.")
                } else {
                        clusters := strings.Split(clusterIds, ",")
                        pred = func(config Config) bool {

                                for _, clusterId := range clusters {
                                        if config.Config.Id == clusterId {
                                                return true
                                        }
                                }
                                return false
                        }
                }
        }

        //
        //  Read configs
        //
        configs := make([]Config,0)
        if pred != nil {
                configs = ReadConfigs(configDir)
                configs = filterConfigs(pred, configs)
        }

        return configs, flagSet.Args()
}
