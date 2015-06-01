package main

import (
        "cisco/micro/config"
        "cisco/micro/logger"
        "cisco/micro/provider"
        "os"
        "fmt"
)

const defaultLocation string = ".micro/src"

func printHelp() {
        fmt.Printf(`
Usage: micro <command> <options>

TODO print usage (e.g. why to use init ...)
        `)
        fmt.Println("")
}

func unsufficientCliOptions(args []string) bool{
        // quick hack to show help without using flag
        if len(args) == 1 {
                option := args[0]
                if (option == "help" || option == "-h" || option == "--help"){
                        return true
                }
        } else if len(args) == 0 {
                return true
        }
        return false
}

func main() {


        if unsufficientCliOptions(os.Args[1:]) {
                printHelp()
                os.Exit(0)
        }

        configs, args := config.MatchConfigs(os.Args[1:])

        if len(configs) == 0 {
                logger.Errorf("No matching configurations found.")
        }

        // TODO collect return values to determine exit code
        for _, cfg := range configs {
                provider.Dispatch(cfg, args)
        }
}
