package log

import (
	"fmt"
	"io"
	"os"
)

var (
	// Verbose enables debug.
	Verbose = false
	// Width of column.
	Width = "14"
	// Writer defines the io.Writer for all non-error output (Debug, Info, Warn)
	Writer io.Writer = os.Stdout
	// ErrorWriter defines the io.Writer for error output (Error, Fatal)
	ErrorWriter io.Writer = os.Stderr
)

// Info.
func Info(key string, format string, args ...interface{}) {
	p := fmt.Sprintf("\033[36m%"+Width+"s \033[90m:\033[0m ", key)
	fmt.Fprintf(Writer, p+format+"\n", args...)
}

// Debug.
func Debug(key string, format string, args ...interface{}) {
	if Verbose {
		p := fmt.Sprintf("\033[96m%"+Width+"s \033[90m:\033[0m ", key)
		fmt.Fprintf(Writer, p+format+"\n", args...)
	}
}

// Warning.
func Warn(format string, args ...interface{}) {
	p := fmt.Sprintf("\033[33m%"+Width+"s \033[90m:\033[0m ", "warn")
	fmt.Fprintf(Writer, p+format+"\n", args...)
}

// Error.
func Error(err error) {
	p := fmt.Sprintf("\033[31m%"+Width+"s \033[90m:\033[0m ", "error")
	fmt.Fprintf(ErrorWriter, p+"%s\n", err.Error())
}

// Fatal.
func Fatal(err error) {
	Error(err)
	os.Exit(1)
}
