package provider

import (
        "fmt"
        "os"
        "cisco/micro/config"
)


type CommandFunction func(config.Config, []string) int

type Provider struct {
        dispatchTable map[string]CommandFunction
}

func NewProvider(dispatchTable map[string]CommandFunction) *Provider {
        return &Provider{
                dispatchTable: dispatchTable}
}

var providers = map[string]*Provider{
        "aws": NewProvider(getAwsDispatchTable()),
        "gce": NewProvider(getGceDispatchTable()),

}

func Dispatch(cfg config.Config, args[]string) int {
        if provider, ok := providers[cfg.Config.Provider]; ok {
                return provider.dispatch(cfg, args)
        } else {
                fmt.Fprintf(os.Stderr, `No arguments given.`)
                return 1
        }

}

func (provider *Provider) dispatch(cfg config.Config, args []string) int {

        dispatchTable := provider.dispatchTable

        if (len(args) == 0){
                fmt.Fprintf(os.Stderr, `No arguments given.`)
                return 1
        }

        id := args[0]

        if commandFn, ok := dispatchTable[id]; ok {
                return commandFn(cfg, args)
        }

        fmt.Fprintf(os.Stderr, `Unknown command.`)
        return 1
}



