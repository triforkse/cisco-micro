package provider

import (
        "fmt"
        "cisco/micro/config"
)

func GceBuild(cfg config.Config, args []string) int{
        fmt.Println("GCE build")
        return 0
}



func getGceDispatchTable() map[string] CommandFunction{
        return map[string] CommandFunction{
                "build": GceBuild,
        }
}
