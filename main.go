package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	Version string = strings.TrimSpace(version)
	//go:embed VERSION
	version string
)

var directory = ""
var _batchFiles []string
var variables = map[string]string{}
var arguments = map[string]string{}
var continueTo = ""
var currentFile string
var currentFileLine = 0
var userInput string

func main() {
	fmt.Printf("version: %s\n", Version)
	if len(os.Args) > 1 {
		directory = os.Args[1]
	}
	fmt.Printf("Using directory: %s\n", directory)

	err := filepath.Walk(directory, walkFn)
	if err != nil {
		panic(err)
	}

	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")

		userInput, err = reader.ReadString('\n')
		userArguments := strings.Split(userInput, " ")
		arguments = make(map[string]string)
		for i := range userArguments {
			if i == 0 {
				continue
			}
			arguments["%"+strconv.Itoa(i)] = strings.ReplaceAll(userArguments[i], "\n", "")
		}
		err = executeCommand(strings.ReplaceAll(userInput, "\n", ""))
		if err != nil {
			panic(err)
		}
	}
}

func executeCommand(line string) error {
	if continueTo != "" {
		if continueTo != line {
			return nil
		}
		continueTo = ""
		return nil
	}

	if strings.HasPrefix(line, "SET ") {
		key, value := parseVariable(line)
		variables[key] = value
	}
	filename, isBat := checkBatchFile(line)
	if isBat {
		currentFile = filename
		currentFileLine = 0
		return processFile(filepath.Join(directory, filename))
	}
	if strings.HasPrefix(line, "ECHO.") {
		fmt.Println()
	}
	if strings.HasPrefix(line, "ECHO ") {
		s := line[len("ECHO "):]
		fmt.Println(s)
	}
	if strings.HasPrefix(line, "GOTO ") {
		continueTo = ":" + line[len("GOTO "):]
	}
	if strings.HasPrefix(line, "IF ") {
		commandToExecute := parseIfStatement(line)
		return executeCommand(commandToExecute)
	}
	return nil
}

func processFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		currentFileLine++

		if strings.HasPrefix(line, "@") {
			continue
		}
		err = executeCommand(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func parseIfStatement(line string) string {
	var statement string
	isNegation := false
	statement = line[len("IF "):]
	if strings.HasPrefix(line, "IF NOT ") {
		statement = line[len("IF NOT "):]
		isNegation = true
	}

	statementParts := strings.Split(statement, " ")[0]
	conditionParts := strings.Split(statementParts, "==")
	index := strings.Index(statement, " ")
	commandToExecute := statement[index+1:]
	if len(conditionParts[0]) == 4 { // "%1"
		argumentName := strings.ReplaceAll(conditionParts[0], "\"", "")
		batchArgumentValue := strings.ReplaceAll(conditionParts[1], "\"", "")

		if isNegation {
			if arguments[argumentName] != batchArgumentValue {
				return commandToExecute
			}
		} else {
			if arguments[argumentName] == batchArgumentValue {
				return commandToExecute
			}
		}
		return ""
	}
	if strings.HasPrefix(conditionParts[0], "\"%") {
		firstPartVariableValue := resolveVariable(conditionParts[0])
		secondPartVariable := resolveVariable(conditionParts[1])

		if isNegation {
			if firstPartVariableValue != secondPartVariable {
				return commandToExecute
			}
		} else {
			if firstPartVariableValue == secondPartVariable {
				return commandToExecute
			}
		}
	}
	return ""
}

func resolveVariable(variable string) string {
	variableName := strings.ReplaceAll(variable, "\"", "")
	if strings.HasPrefix(variableName, "%") {
		variableName = strings.ReplaceAll(variableName, "%", "")
		variableValue := variables[variableName]
		return variableValue
	}
	return variableName
}

func checkBatchFile(line string) (string, bool) {
	parts := strings.Split(line, " ")
	for i := range _batchFiles {
		if _batchFiles[i] == strings.ToUpper(parts[0])+".BAT" {
			setArguments(parts[1:])
			return _batchFiles[i], true
		}
	}
	return "", false
}

func setArguments(args []string) {
	for i := range args {
		value := strings.ReplaceAll(args[i], "\n", "")
		argumentName := "%" + strconv.Itoa(i+1)
		if value == argumentName {
			continue
		}
		arguments[argumentName] = value
	}
}

func parseVariable(text string) (string, string) {
	parts := strings.Split(text, "=")
	key := (parts[0])[len("SET "):]
	value := resolveVariable(parts[1])
	return key, value
}

func walkFn(path string, info fs.FileInfo, err error) error {
	if filepath.Ext(info.Name()) != ".BAT" {
		return nil
	}
	_batchFiles = append(_batchFiles, info.Name())
	return nil
}
