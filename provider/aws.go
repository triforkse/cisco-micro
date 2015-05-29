package provider

import (
        "fmt"
        "cisco/micro/config"
)

func AwsBuild(cfg []config.Config, args []string) int{
        fmt.Println("AWS build")
        return 0
}



func getAwsDispatchTable() map[string] CommandFunction{
        return map[string] CommandFunction{
                "build": AwsBuild,
        }
}
