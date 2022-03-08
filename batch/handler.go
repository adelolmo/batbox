package batch

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func (b *Batch) handlePause() {
	fmt.Print("Press any key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Text()
}

func (b *Batch) handleClearScreen() {
	fmt.Print("\033[H\033[2J")
}

/*func (b *Batch) handleCall(line string) {
	parts := strings.Split(line, " ")

	err := b.processFile(parts[0])
}
*/
func (b Batch) handleEchoDot() {
	fmt.Fprintln(out)
}

func (b Batch) handleEcho(line string) {
	fmt.Fprintln(out, line[len("ECHO "):])
}

func (b *Batch) handleSet(line string) {
	parts := strings.Split(line, "=")
	key := (parts[0])[len("SET "):]
	value := b.resolveVariable(parts[1])
	b.variables[key] = value
}

func (b *Batch) handleGoto(line string) {
	b.continueTo = ":" + line[len("GOTO "):]
}

func (b *Batch) handleIf(line string) string {
	line = strings.ReplaceAll(line, "if not ", "IF NOT ")
	line = strings.ReplaceAll(line, "if ", "IF ")
	line = strings.ReplaceAll(line, " == ", "==")

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

	b.log("\t%s=%s\n", conditionParts[0], firstPartValue)
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
	argumentName := removeQuotes(argument)
	return b.arguments[argumentName]
}

func (b *Batch) resolveVariable(variable string) string {
	variableName := removeQuotes(variable)
	if strings.HasPrefix(variableName, "%") {
		variableName = strings.ReplaceAll(variableName, "%", "")
		variableValue := b.variables[variableName]
		return variableValue
	}
	return variableName
}

func removeQuotes(variable string) string {
	return strings.ReplaceAll(variable, "\"", "")
}
