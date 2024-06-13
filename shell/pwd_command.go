package shell

import (
	"fmt"
	"os"
)

type PwdCommand struct{}

func (c *PwdCommand) Execute(args []string) bool {
	if len(args) > 0 {
		fmt.Println("pwd: too many arguments")
		return true
	}

	if dir, err := os.Getwd(); err == nil {
		fmt.Println(dir)
	} else {
		fmt.Println("pwd: could not determine current directory")
	}
	return true
}
