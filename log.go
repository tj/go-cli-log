package log

import "fmt"

//
// Toggle verbose logging.
//

var Verbose = false

//
// Key width.
//

var Width = "14"

//
// Info.
//

func Info(key string, format string, args ...interface{}) {
	p := fmt.Sprintf("\033[36m%"+Width+"s \033[90m:\033[0m ", key)
	fmt.Printf(p+format+"\n", args...)
}

//
// Debug.
//

func Debug(key string, format string, args ...interface{}) {
	if Verbose {
		p := fmt.Sprintf("\033[96m%"+Width+"s \033[90m:\033[0m ", key)
		fmt.Printf(p+format+"\n", args...)
	}
}

//
// Warning.
//

func Warn(format string, args ...interface{}) {
	p := fmt.Sprintf("\033[33m%"+Width+"s \033[90m:\033[0m ", "warn")
	fmt.Printf(p+format+"\n", args...)
}

//
// Error.
//

func Error(err error) {
	fmt.Printf("\033[31m%"+Width+"s \033[90m:\033[0m %s\n", "error", err)
}
