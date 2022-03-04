package batch

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var batchFiles []string

type Batch struct {
	directory       string
	batchFiles      []string
	variables       map[string]string
	currentFile     string
	currentFileLine int
	arguments       map[string]string
	continueTo      string
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
		continueTo:      "",
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
	if bat.continueTo != "" {
		if bat.continueTo != line {
			return nil
		}
		bat.continueTo = ""
		return nil
	}

	filename, isBat := bat.checkBatchFile(line)
	if isBat {
		bat.currentFile = filename
		bat.currentFileLine = 0
		return bat.processFile(filepath.Join(bat.directory, filename))
	}
	if strings.HasPrefix(line, "SET ") {
		bat.handleSet(line)
	}
	if strings.HasPrefix(line, "ECHO.") {
		bat.handleEchoDot()
	}
	if strings.HasPrefix(line, "ECHO ") {
		bat.handleEcho(line)
	}
	if strings.HasPrefix(line, "GOTO ") {
		bat.handleGoto(line)
	}
	if strings.HasPrefix(line, "IF ") {
		commandToExecute := bat.handleIf(line)
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
