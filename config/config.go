package config

import (
        "cisco/micro/logger"
        "encoding/json"
        "io/ioutil"
        "log"
        "os"
        "path/filepath"
        "strings"
        "github.com/jessevdk/go-flags"
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

type CliOptions struct {
        ConfigDir string `short:"d" long:"config-dir" default:"./" description:"Search for configuration files in this directory: defaults to ./"`
        ClusterIds []string `short:"i" long:"id" description:"Comma separated list (no space) of cluster ids"`
        AcceptAllClusters bool `short:"a" long:"all"  description:"Applies command to all cluster found in config-dir"`

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
        var pred func(Config) bool

        //
        //  Parse command line arguments
        //
        cli := &CliOptions{}
        args, _ = flags.NewParser(cli,  flags.HelpFlag | flags.PassDoubleDash | flags.IgnoreUnknown).ParseArgs(args)

        //
        //  Determine which predicate to use
        //
        if cli.AcceptAllClusters {
                pred = func(config Config) bool {
                        return true
                }
        } else {
                if len(cli.ClusterIds) == 0 {
                        logger.Errorf("Missing cluster id(s) or all flag. See help for more infos.")
                } else {
                        pred = func(config Config) bool {

                                for _, clusterId := range cli.ClusterIds {
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
                configs = ReadConfigs(cli.ConfigDir)
                configs = filterConfigs(pred, configs)
        }

        return configs, args
}
