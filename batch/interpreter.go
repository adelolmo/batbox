package batch

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var continueTo = ""
var batchFiles []string

type Batch struct {
	directory       string
	batchFiles      []string
	variables       map[string]string
	currentFile     string
	currentFileLine int
	arguments       map[string]string
}

func NewInterpreter(directory string) Batch {
	err := filepath.Walk(directory, walkFn)
	if err != nil {
		panic(err)
	}
	return Batch{
		directory:       directory,
		variables:       map[string]string{},
		currentFileLine: 0,
		arguments:       map[string]string{},
		batchFiles:      batchFiles,
	}
}

func walkFn(path string, info fs.FileInfo, err error) error {
	if filepath.Ext(info.Name()) != ".BAT" {
		return nil
	}
	batchFiles = append(batchFiles, info.Name())
	return nil
}

func (bat *Batch) ExecuteCommand(line string) error {
	if continueTo != "" {
		if continueTo != line {
			return nil
		}
		continueTo = ""
		return nil
	}

	if strings.HasPrefix(line, "SET ") {
		key, value := bat.parseVariable(line)
		bat.variables[key] = value
	}
	filename, isBat := bat.checkBatchFile(line)
	if isBat {
		bat.currentFile = filename
		bat.currentFileLine = 0
		return bat.processFile(filepath.Join(bat.directory, filename))
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
		commandToExecute := bat.parseIfStatement(line)
		return bat.ExecuteCommand(commandToExecute)
	}
	return nil
}

func (bat *Batch) SetArguments(args []string) {
	for i := range args {
		argumentName := "%" + strconv.Itoa(i+1)
		if args[i] == argumentName {
			continue
		}
		bat.arguments[argumentName] = args[i]
	}
}

func (bat *Batch) processFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		bat.currentFileLine++

		if strings.HasPrefix(line, "@") {
			continue
		}
		err = bat.ExecuteCommand(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
func (bat *Batch) parseIfStatement(line string) string {
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
			if bat.arguments[argumentName] != batchArgumentValue {
				return commandToExecute
			}
		} else {
			if bat.arguments[argumentName] == batchArgumentValue {
				return commandToExecute
			}
		}
		return ""
	}
	if strings.HasPrefix(conditionParts[0], "\"%") {
		firstPartVariableValue := bat.resolveVariable(conditionParts[0])
		secondPartVariable := bat.resolveVariable(conditionParts[1])

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
func (bat *Batch) resolveVariable(variable string) string {
	variableName := strings.ReplaceAll(variable, "\"", "")
	if strings.HasPrefix(variableName, "%") {
		variableName = strings.ReplaceAll(variableName, "%", "")
		variableValue := bat.variables[variableName]
		return variableValue
	}
	return variableName
}

func (bat *Batch) parseVariable(text string) (string, string) {
	parts := strings.Split(text, "=")
	key := (parts[0])[len("SET "):]
	value := bat.resolveVariable(parts[1])
	return key, value
}

func (bat *Batch) checkBatchFile(line string) (string, bool) {
	parts := strings.Split(line, " ")
	for i := range bat.batchFiles {
		if bat.batchFiles[i] == strings.ToUpper(parts[0])+".BAT" {
			bat.SetArguments(parts[1:])
			return bat.batchFiles[i], true
		}
	}
	return "", false
}
