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

type Batch struct {
	directory      string
	batchFiles     []string
	variables      map[string]string
	arguments      map[string]string
	continueTo     string
	executionStack executionStack
	debug          bool
}

func NewInterpreter(directory string, debug bool) *Batch {
	batchFiles, err := collectBatchFiles(directory)
	if err != nil {
		panic(err)
	}
	return &Batch{
		directory:      directory,
		variables:      map[string]string{},
		arguments:      map[string]string{},
		batchFiles:     batchFiles,
		continueTo:     "",
		executionStack: executionStack{fileStack: []executingBatchFile{}},
		debug:          debug,
	}
}

func collectBatchFiles(directory string) ([]string, error) {
	var batchFiles []string
	err := filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(info.Name()) != ".BAT" {
			return nil
		}
		batchFiles = append(batchFiles, info.Name())
		return nil
	})
	return batchFiles, err
}

func (b *Batch) SetArguments(args []string) {
	if len(args) == 0 {
		b.arguments = make(map[string]string)
	}
	for i := range args {
		argumentName := args[i]
		if strings.HasPrefix(argumentName, "%") {
			name := "%" + strconv.Itoa(i+1)
			b.arguments[name] = b.arguments[args[i]]
			continue
		}

		argumentName = "%" + strconv.Itoa(i+1)
		if args[i] == argumentName {
			continue
		}
		b.arguments[argumentName] = args[i]
	}
}

func (b *Batch) ExecuteCommand(line string) error {
	if b.continueTo != "" {
		if b.continueTo != line {
			return nil
		}
		b.continueTo = ""
		return nil
	}

	if len(line) == 0 {
		return nil
	}
	if strings.HasPrefix(line, "@") {
		return nil
	}
	if strings.HasPrefix(line, "REM ") {
		return nil
	}
	b.logCommand(line)
	filename, isBat := b.checkBatchFile(line)
	if isBat {
		return b.processFile(filename)
	}
	if "PAUSE" == line {
		b.handlePause()
	}
	if "CLS" == line {
		b.handleClearScreen()
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

func (b *Batch) processFile(filename string) error {
	b.setCurrentFile(filename)

	file, err := os.Open(filepath.Join(b.directory, filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		b.executionStack.increaseExecutionLine()
		err = b.ExecuteCommand(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	b.executionStack.drop()

	b.arguments = make(map[string]string)
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

func (b *Batch) setCurrentFile(filename string) {
	batchFile := executingBatchFile{
		name:       filename,
		lineNumber: 0,
	}
	b.executionStack.push(batchFile)
}
