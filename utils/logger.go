package utils

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type Logger struct {
	Verbose bool
}

var logger *Logger

func NewLogger(verbose bool) *Logger {
	logger = &Logger{
		Verbose: verbose,
	}
	return logger
}

func GetLogger() *Logger {
	if logger == nil {
		logger = &Logger{
			Verbose: true,
		}
	}
	return logger
}

func (l *Logger) Info(message string) {
	if l.Verbose {
		now := time.Now()
		formattedTime := now.Format(time.RFC3339)
		color.Green("%s [INFO] %s", formattedTime, message)
	} else {
		fmt.Printf("%s\n", message)
	}
}

func (l *Logger) Error(message string) {
	now := time.Now()
	formattedTime := now.Format(time.RFC3339)
	color.Red("%s [ERROR] %s", formattedTime, message)
}

func (l *Logger) Warning(message string) {
	now := time.Now()
	formattedTime := now.Format(time.RFC3339)
	color.Yellow("%s [WARN] %s", formattedTime, message)
}

func (l *Logger) Debug(message string) {
	now := time.Now()
	formattedTime := now.Format(time.RFC3339)
	if l.Verbose {
		fmt.Printf("%s [DEBUG] %s\n", formattedTime, message)
	}
}
