package shell

import (
	"fmt"
	"strings"
)

type EchoCommand struct{}

func (c *EchoCommand) Execute(args []string) bool {
	omitNewline := false
	if len(args) > 0 && args[0] == "-n" {
		omitNewline = true
		args = args[1:]
	}

	output := strings.Join(args, " ")
	if omitNewline {
		fmt.Print(output)
	} else {
		fmt.Println(output)
	}
	return true
}
