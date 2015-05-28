package main

import (
        "testing"
)

const SUCCESSFUL_EXIT_CODE int = 0
const FAILURE_EXIT_CODE int = 1


func TestDispatchKnownCommand(t *testing.T) {
        args := []string{"cmd-2"}

        command1 := NewMockedCommand("cmd-1", SUCCESSFUL_EXIT_CODE)
        command2 := NewMockedCommand("cmd-2", SUCCESSFUL_EXIT_CODE)

        exitValue := mainDispatch([]*Command{command1.cmd, command2.cmd}, args)


        if exitValue != SUCCESSFUL_EXIT_CODE {
                t.Error("unexpected return value of our command ")
        }

        if command1.called {
                t.Error("expected command function not to be called")
        }

        if !command2.called {
                t.Error("expected command function not to be called")
        }
}


func TestDispatchUnknownCommand(t *testing.T) {
        args := []string{"unknown-command"}

        command1 := NewMockedCommand("cmd-1", SUCCESSFUL_EXIT_CODE)

        exitValue := mainDispatch([]*Command{command1.cmd}, args)


        if exitValue != FAILURE_EXIT_CODE {
                t.Error("unexpected return value of our command ")
        }
}

type MockCommand struct {
        called bool
        args []string
        mockedReturnValue int
        cmd *Command
}


func (self *MockCommand) mockedFunction(args []string) int {
        self.called = true
        self.args = args
        return self.mockedReturnValue
}


func NewMockedCommand(id string, exitCode int) *MockCommand{
        mockCommandFunction := &MockCommand{mockedReturnValue: exitCode}
        cmd := NewCommand(id, mockCommandFunction.mockedFunction)
        mockCommandFunction.cmd = cmd
        return mockCommandFunction
}

