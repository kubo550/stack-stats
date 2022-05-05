package consoleLog

import (
	"fmt"
	"strings"
)

const (
	cInfo  = "\x1b[36m%s\x1b[0m"
	cWarn  = "\x1b[33m%s\x1b[0m"
	cError = "\x1b[31m%s\x1b[0m"
)

func printEvent(level string, color, message string) {
	fmt.Println()
	fmt.Printf(color, fmt.Sprintf("[%s] %s", strings.ToUpper(level), message))
	fmt.Println()
}

func Info(message string) {
	printEvent("info", cInfo, message)
}

func Warning(message string) {
	printEvent("warning", cWarn, message)
}

func Error(message error) {
	printEvent("error", cError, message.Error())
}
