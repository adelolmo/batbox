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
	err := filepath.Walk(directory, collectBatchFiles)
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

func collectBatchFiles(path string, info fs.FileInfo, err error) error {
	if filepath.Ext(info.Name()) != ".BAT" {
		return nil
	}
	batchFiles = append(batchFiles, info.Name())
	return nil
}

func (b *Batch) SetArguments(args []string) {
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

	if strings.HasPrefix(line, "@") {
		return nil
	}
	filename, isBat := b.checkBatchFile(line)
	if isBat {
		return b.processFile(filename)
	}
	if "PAUSE" == line {
		fmt.Print("Press any key to continue.")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		scanner.Text()
	}
	if "CLS" == line {
		fmt.Print("\033[H\033[2J")
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
	b.currentFile = filename
	b.currentFileLine = 0

	file, err := os.Open(filepath.Join(b.directory, filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		b.currentFileLine++
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
