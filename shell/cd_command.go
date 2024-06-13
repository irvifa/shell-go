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
	var targetPath string

	if path == "~" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("cd: could not determine home directory")
			return true
		}
		targetPath = homeDir
	} else {
		targetPath, _ = filepath.Abs(path)
	}

	if err := os.Chdir(targetPath); err != nil {
		fmt.Printf("cd: %s: %s\n", path, err.Error())
	}
	return true
}
