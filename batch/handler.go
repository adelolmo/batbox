package batch

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func (b Batch) handleEchoDot() {
	fmt.Fprintln(out)
}

func (b Batch) handleEcho(line string) {
	fmt.Fprintln(out, line[len("ECHO "):])
}

func (b *Batch) handleSet(line string) {
	key, value := b.parseVariable(line)
	b.variables[key] = value
}

func (b *Batch) handleGoto(line string) {
	b.continueTo = ":" + line[len("GOTO "):]
}

func (b *Batch) handleIf(line string) string {
	return b.parseIfStatement(line)
}

func (b *Batch) parseVariable(text string) (string, string) {
	parts := strings.Split(text, "=")
	key := (parts[0])[len("SET "):]
	value := b.resolveVariable(parts[1])
	return key, value
}

func (b *Batch) parseIfStatement(line string) string {
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

	var firstPartValue string
	if len(conditionParts[0]) == 4 {
		firstPartValue = b.resolveArgument(conditionParts[0])
	} else if strings.HasPrefix(conditionParts[0], "\"%") {
		firstPartValue = b.resolveVariable(conditionParts[0])
	} else {
		firstPartValue = conditionParts[0]
	}
	secondPartVariable := b.resolveVariable(conditionParts[1])

	if isNegation {
		if firstPartValue != secondPartVariable {
			return commandToExecute
		}
	} else {
		if firstPartValue == secondPartVariable {
			return commandToExecute
		}
	}
	return ""
}

func (b *Batch) resolveArgument(argument string) string {
	argumentName := sanitizeVariableName(argument)
	return b.arguments[argumentName]
}

func (b *Batch) resolveVariable(variable string) string {
	variableName := sanitizeVariableName(variable)
	if strings.HasPrefix(variableName, "%") {
		variableName = strings.ReplaceAll(variableName, "%", "")
		variableValue := b.variables[variableName]
		return variableValue
	}
	return variableName
}

func sanitizeVariableName(variable string) string {
	return strings.ReplaceAll(variable, "\"", "")
}
