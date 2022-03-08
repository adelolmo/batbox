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
	Version = strings.TrimSpace(version)
	//go:embed VERSION
	version string
)

func main() {
	fmt.Printf("version: %s\n", Version)
	directory := "."
	debug := false
	for i := range os.Args[1:] {
		if os.Args[i+1] == "-d" {
			debug = true
		} else {
			directory = os.Args[i+1]
		}
	}
	fmt.Printf("Using directory: %s\n", directory)

	bat := batch.NewInterpreter(directory, debug)
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
