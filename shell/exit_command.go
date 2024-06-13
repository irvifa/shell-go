package shell

import (
	"fmt"
	"os"
	"strconv"
)

type ExitCommand struct{}

func (c *ExitCommand) Execute(args []string) bool {
	if len(args) == 1 {
		if code, err := strconv.Atoi(args[0]); err == nil {
			os.Exit(code)
		} else {
			fmt.Println("exit: invalid status code")
		}
	} else {
		fmt.Println("exit: invalid number of arguments")
	}
	return true
}
