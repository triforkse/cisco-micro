package logger

import (
  "fmt"
)

var isDebugging bool = false

func EnableDebug(debug bool) {
  isDebugging = debug
}

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


func Debugf(format string, args... interface{}) {
  if isDebugging {
    fmt.Print(purple)
    fmt.Printf(format, args...)
    fmt.Println(noColor)
  }
}
