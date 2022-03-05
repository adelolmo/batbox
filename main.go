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
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		userInput := scanner.Text()
		bat.SetArguments(strings.Split(userInput, " ")[1:])
		err := bat.ExecuteCommand(userInput)
		if err != nil {
			panic(err)
		}
	}
}
