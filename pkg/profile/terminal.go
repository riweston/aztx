package profile

import "fmt"

// The terminal file contains helper functions primarily for sending messages to the terminal.

const (
	InfoColor    = "\033[0;32m%s\033[0m"
	NoticeColor  = "\033[0;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

// PrintInfo prints an info message to the terminal.
func PrintInfo(message string) {
	fmt.Println(fmt.Sprintf(InfoColor, message))
}

// PrintNotice prints a notice message to the terminal.
func PrintNotice(message string) {
	fmt.Println(fmt.Sprintf(NoticeColor, message))
}
