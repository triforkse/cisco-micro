package main

import (
        "cisco/micro/config"
        "cisco/micro/logger"
        "cisco/micro/provider"
        "os"
)

const defaultLocation string = ".micro/src"

func main() {
        configs, args := config.MatchConfigs(os.Args[1:])

        if len(configs) == 0 {
                logger.Errorf("No matching configurations found.")
        }

        for _, cfg := range configs {
                logger.Debugf("Apply to cluster %s using configuration file: %s", cfg.Config.Id, cfg.Path)
                provider.Dispatch(cfg, args)
        }
}
