package shell

import (
	"fmt"
	"os"
	"path/filepath"
)

type CdCommand struct{}

func (c *CdCommand) Execute(args []string) bool {
	if len(args) != 1 {
		fmt.Println("cd: invalid number of arguments")
		return true
	}

	path := args[0]
	targetPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("cd: %s: %s\n", path, err.Error())
		return true
	}

	if err := os.Chdir(targetPath); err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", path)
	}
	return true
}
