package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/adelolmo/batbox/batch"
	"os"
	"strings"
)

var (
	Version string = strings.TrimSpace(version)
	//go:embed VERSION
	version string
)

func main() {
	fmt.Printf("version: %s\n", Version)
	directory := "."
	if len(os.Args) > 1 {
		directory = os.Args[1]
	}
	fmt.Printf("Using directory: %s\n", directory)

	bat := batch.NewInterpreter(directory)
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		userInput, err := reader.ReadString('\n')
		userArguments := strings.Split(userInput, " ")
		bat.SetArguments(userArguments[1:])
		err = bat.ExecuteCommand(strings.ReplaceAll(userInput, "\n", ""))
		if err != nil {
			panic(err)
		}
	}
}
