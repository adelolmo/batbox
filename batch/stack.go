package batch

type executionStack struct {
	fileStack []executingBatchFile
}

type executingBatchFile struct {
	name       string
	lineNumber int
	arguments  map[string]string
}

func NewExecutionStack() *executionStack {
	return &executionStack{
		fileStack: []executingBatchFile{},
	}
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

func (f *executionStack) top() (*executingBatchFile, bool) {
	if len(f.fileStack) == 0 {
		return nil, true
	}
	topElementIndex := len(f.fileStack) - 1
	return &f.fileStack[topElementIndex], false
}
