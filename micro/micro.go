package main

import (
        "os"
)


// commands
var BuildCommand = NewCommand("build", buildCommand)
var HelpCommand = NewCommand( "help", helpCommand)

var AllCommands = []*Command{BuildCommand, HelpCommand}

// special command which is only implicit available when unknown input
var UnknownInput = &Command{commandFn: unknownCommand}


type CommandFunction func([]string) int

type Executable interface {
        Execute() CommandFunction
}

type Command struct {
        id        string
        commandFn CommandFunction
}

func NewCommand(id string, commandFunction CommandFunction) *Command {
        return &Command{
                id: id,
                commandFn: commandFunction}
}

func (self *Command) Execute(args []string) int {
        return self.commandFn(args)
}



func unknownCommand(args []string) int {
        println("unknown command")
        return 1
}
func helpCommand(args []string) int {
        println("help command")
        return 0
}

func buildCommand(args []string) int {
        println("build command")
        return 0
}


func determineCommand(commands []*Command, args []string) *Command {
        id := args[0]

        for _, cmd := range commands {
                if cmd.id == id {
                        return cmd
                }
        }

        return UnknownInput
}

func mainDispatch(commands []*Command, args []string) int {
        cmd := determineCommand(commands, args)
        return cmd.Execute(args)
}



func main() {
        args := os.Args[1:]

        exitCode := mainDispatch(AllCommands, args)

        os.Exit(exitCode)
}


//func main() {
//
//	filePath := flag.String("config", "infrastructure.json", "the configuration file")
//	isDebugging := flag.Bool("debug", false, "show debug info")
//
//	flag.Parse()
//
//	logger.EnableDebug(*isDebugging)
//
//	var command string
//	cmdArgs := flag.Args()
//	if len(cmdArgs) > 0 {
//		command = cmdArgs[0]
//	} else {
//		command = "apply"
//	}
//
//	logger.Debugf("Command: %s", command)
//	logger.Debugf("Config File: %s", *filePath)
//
//	switch command {
//	case "init":
//		logger.Debugf("Is file:")
//		providerId := cmdArgs[1]
//		initCmd(providerId, *filePath, defaultLocation)
//	case "apply", "destroy", "plan":
//		config := provider.FromFile(*filePath)
//		 TODO: handle read error here, not in the lib
//		terraform.TerraformCmd(command, config, filepath.Join(defaultLocation, "terraform"))
//	case "build":
//		config := provider.FromFile(*filePath)
//		err := packerCmd(config)
//		if err != nil {
//			log.Fatal("Could not run packer. " + err.Error())
//		}
//	}
//}
//
//
//
