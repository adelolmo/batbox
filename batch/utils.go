package batch

import "fmt"

func (b *Batch) logCommand(line string) {
	if len(b.executionStack.fileStack) == 0 {
		return
	}

	batchFile := b.executionStack.fileStack[len(b.executionStack.fileStack)-1]
	b.log("\t%s %d\t%s\n", batchFile.name, batchFile.lineNumber, line)
}

func (b *Batch) log(format string, a ...interface{}) {
	if b.debug {
		fmt.Printf(format, a...)
	}
}
