package main

import (
        "os"
        "cisco/micro/logger"
        "cisco/micro/config"
        "cisco/micro/provider"
)

const defaultLocation string = ".micro/src"

func main() {

        configs, args := config.MatchConfigs(os.Args[2:])

        for _, cfg := range configs {
                logger.Debugf("Apply to cluster %s using configuration file: %s", cfg.Config.Id,cfg.Path)
                provider.Dispatch(cfg, args)
        }
}

