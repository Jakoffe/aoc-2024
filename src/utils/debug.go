package utils

import (
	"fmt"
)

func DebugPrint(str string, debugMode bool) {
	if debugMode {
		fmt.Print(str)
	}
}

func DebugPrintln(str string, debugMode bool) {
	if debugMode {
		fmt.Println(str)
	}
}

func DebugPrintF(format string, debugMode bool, args ...any) {
	if debugMode {
		fmt.Printf(format, args...)
	}
}
