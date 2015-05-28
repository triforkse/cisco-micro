package provider

import (
        "fmt"
        "os"
        "cisco/micro/config"
)



type CommandFunction func(config.Config,[]string)int

type Provider struct {
        dispatchTable map[string]CommandFunction
}

func NewProvider(dispatchTable map[string]CommandFunction) *Provider {
        return &Provider{
                dispatchTable: dispatchTable}
}

var providers = map[string]*Provider {
        "aws": NewProvider(getAwsDispatchTable()),
        "gce": NewProvider(getGceDispatchTable()),

}

func Dispatch(cfg config.Config, args[]string) int{
        return providers[cfg.Config.Provider].dispatch(cfg, args)
}

func (provider *Provider) dispatch(cfg config.Config, args []string) int {

        dispatchTable := provider.dispatchTable
        id := args[0]

        if commandFn, ok := dispatchTable[id]; ok {
                return commandFn(cfg, args)
        }

        fmt.Fprintf(os.Stderr, `Unknown command.\n`)
        return 1
}



