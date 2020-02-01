package logger

import "fmt"

var StdoutLogger = stdoutLogger{}

type stdoutLogger struct{}

func (bl stdoutLogger) Print(args ...interface{}) {
	fmt.Print(args...)
}

func (bl stdoutLogger) Printf(log string, args ...interface{}) {
	fmt.Printf(log, args...)
}

type Stdout interface {
	Print(...interface{})
	Printf(string, ...interface{})
}
