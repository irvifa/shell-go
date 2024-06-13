package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Shell struct {
	commands map[string]Command
}

func New() *Shell {
	commands := map[string]Command{
		"exit": &ExitCommand{},
		"echo": &EchoCommand{},
		"cd":   &CdCommand{},
		"pwd":  &PwdCommand{},
	}
	commands["type"] = NewTypeCommand(commands)
	return &Shell{commands: commands}
}

func (s *Shell) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$ ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if input == "" {
			continue
		}

		command := s.tokenize(input)
		if !s.execute(command) {
			s.runExternalCommand(command)
		}
	}
}

type CommandStruct struct {
	name string
	args []string
}

func (s *Shell) tokenize(input string) CommandStruct {
	parts := strings.Fields(input)
	return CommandStruct{
		name: parts[0],
		args: parts[1:],
	}
}

func (s *Shell) execute(command CommandStruct) bool {
	if cmd, ok := s.commands[command.name]; ok {
		return cmd.Execute(command.args)
	}
	return false
}

func (s *Shell) runExternalCommand(command CommandStruct) {
	cmd := exec.Command(command.name, command.args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("%s: command not found\n", command.name)
	}
}
