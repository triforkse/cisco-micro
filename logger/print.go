package logger

import (
	"fmt"
	"os"
)

var isDebugging bool = true

func EnableDebug(debug bool) {
	isDebugging = debug
}

var red string = "\033[0;31m"
var green string = "\033[0;32m"
var blue string = "\033[0;34m"
var purple string = "\033[0;35m"
var noColor string = "\033[0m"

func PrintTable(title string, tbl map[string]string) {

	fmt.Printf("%s\n%s:%s\n\n", blue, title, noColor)
	for k, v := range tbl {
		fmt.Printf("  %s%s%s = %s%s%s\n", green, k, noColor, blue, v, noColor)
	}

	fmt.Println("")
}

func Debugf(format string, args ...interface{}) {
	if isDebugging {
		fmt.Print(purple)
		fmt.Printf(format, args...)
		fmt.Println(noColor)
	}
}

func Messagef(format string, args ...interface{}) {
	fmt.Print(green)
	fmt.Printf(format, args...)
	fmt.Println(noColor)
}

func Errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, red)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr, noColor)
}
