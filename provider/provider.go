package provider

import (
        "os"
        "fmt"
)



type CommandFunction func(string,[]string)int

type Provider struct {
        dispatchTable map[string]CommandFunction
}

func NewProvider(dispatchTable map[string]CommandFunction) *Provider {
        return &Provider{
                dispatchTable: dispatchTable}
}

func (provider *Provider) dispatch(cluster string, args []string) int {

        dispatchTable := provider.dispatchTable
        id := args[0]

        if commandFn, ok := dispatchTable[id]; ok {
                return commandFn(cluster, args)
        }

        fmt.Fprintf(os.Stderr, `Unknown command.\n`)
        return 1
}

