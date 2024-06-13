package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.

	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Printf("%s: command not found\n", strings.TrimRight(command, "\n"))
}
