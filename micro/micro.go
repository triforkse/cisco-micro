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
                provider.Dispatch(cfg, args)
        }
}
