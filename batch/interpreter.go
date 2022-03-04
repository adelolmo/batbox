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

func (b *Batch) ExecuteCommand(line string) error {
	if b.continueTo != "" {
		if b.continueTo != line {
			return nil
		}
		b.continueTo = ""
		return nil
	}

	filename, isBat := b.checkBatchFile(line)
	if isBat {
		b.currentFile = filename
		b.currentFileLine = 0
		return b.processFile(filepath.Join(b.directory, filename))
	}
	if strings.HasPrefix(line, "SET ") {
		b.handleSet(line)
	}
	if strings.HasPrefix(line, "ECHO.") {
		b.handleEchoDot()
	}
	if strings.HasPrefix(line, "ECHO ") {
		b.handleEcho(line)
	}
	if strings.HasPrefix(line, "GOTO ") {
		b.handleGoto(line)
	}
	if strings.HasPrefix(line, "IF ") {
		commandToExecute := b.handleIf(line)
		return b.ExecuteCommand(commandToExecute)
	}
	return nil
}

func (b *Batch) SetArguments(args []string) {
	for i := range args {
		argumentName := args[i]
		if strings.HasPrefix(argumentName, "%") {
			name := "%" + strconv.Itoa(i+1)
			value := strings.ReplaceAll(args[i], "\n", "")
			b.arguments[name] = b.arguments[value]
			continue
		}

		argumentName = "%" + strconv.Itoa(i+1)
		if args[i] == argumentName {
			continue
		}
		value := strings.ReplaceAll(args[i], "\n", "")
		b.arguments[argumentName] = value
	}
}

func (b *Batch) processFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		b.currentFileLine++

		if strings.HasPrefix(line, "@") {
			continue
		}
		err = b.ExecuteCommand(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (b *Batch) checkBatchFile(line string) (string, bool) {
	parts := strings.Split(line, " ")
	for i := range b.batchFiles {
		if b.batchFiles[i] == strings.ToUpper(parts[0])+".BAT" {
			b.SetArguments(parts[1:])
			return b.batchFiles[i], true
		}
	}
	return "", false
}
