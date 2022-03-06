package batch

type executionStack struct {
	fileStack []executingBatchFile
}

type executingBatchFile struct {
	name       string
	lineNumber int
}

func (f *executionStack) increaseExecutionLine() {
	topFile := &f.fileStack[len(f.fileStack)-1]
	topFile.lineNumber++
}

func (f *executionStack) drop() {
	topElementIndex := len(f.fileStack) - 1
	f.fileStack = f.fileStack[:topElementIndex]
}

func (f *executionStack) push(file executingBatchFile) {
	f.fileStack = append(f.fileStack, file)
}
