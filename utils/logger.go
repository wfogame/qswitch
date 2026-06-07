package utils

import "fmt"

const (
	colorPurple = "\033[35m"
	colorReset  = "\033[0m"
)

func Debug(enabled bool, format string, args ...interface{}) {
	if !enabled {
		return
	}
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s[DEBUG]%s %s\n", colorPurple, colorReset, msg)
}
