package shell

import (
	"fmt"
	"os/exec"
)

type TypeCommand struct {
	builtins map[string]Command
}

func NewTypeCommand(builtins map[string]Command) *TypeCommand {
	return &TypeCommand{builtins: builtins}
}

func (c *TypeCommand) Execute(args []string) bool {
	if len(args) != 1 {
		fmt.Println("type: invalid number of arguments")
		return true
	}

	commandName := args[0]

	if _, ok := c.builtins[commandName]; ok {
		fmt.Printf("%s is a shell builtin\n", commandName)
		return true
	}

	if path, err := exec.LookPath(commandName); err == nil {
		fmt.Printf("%s is %s\n", commandName, path)
		return true
	}

	fmt.Printf("%s: not found\n", commandName)
	return true
}
