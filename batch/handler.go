package batch

import (
	"fmt"
	"strings"
)

func (bat Batch) handleEchoDot() {
	fmt.Println()
}

func (bat Batch) handleEcho(line string) {
	s := line[len("ECHO "):]
	fmt.Println(s)
}

func (bat *Batch) handleSet(line string) {
	key, value := bat.parseVariable(line)
	bat.variables[key] = value
}

func (bat *Batch) handleGoto(line string) {
	bat.continueTo = ":" + line[len("GOTO "):]
}

func (bat *Batch) handleIf(line string) string {
	return bat.parseIfStatement(line)
}

func (bat *Batch) parseVariable(text string) (string, string) {
	parts := strings.Split(text, "=")
	key := (parts[0])[len("SET "):]
	value := bat.resolveVariable(parts[1])
	return key, value
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
		batchArgumentValue := bat.resolveVariable(strings.ReplaceAll(conditionParts[1], "\"", ""))

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
