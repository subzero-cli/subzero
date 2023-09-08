package utils

import (
	"bytes"
	"encoding/json"
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
	if l.Verbose {
		now := time.Now()
		formattedTime := now.Format(time.RFC3339)
		color.Red("%s [INFO] %s", formattedTime, message)
	} else {
		color.Red("%s\n", message)
	}
}

func (l *Logger) Warning(message string) {
	if l.Verbose {
		now := time.Now()
		formattedTime := now.Format(time.RFC3339)
		color.Yellow("%s [INFO] %s", formattedTime, message)
	} else {
		color.Yellow("%s\n", message)
	}
}

func (l *Logger) Debug(message string) {
	now := time.Now()
	formattedTime := now.Format(time.RFC3339)
	if l.Verbose {
		fmt.Printf("%s [DEBUG] %s\n", formattedTime, message)
	}
}

func (l *Logger) JsonDebug(message []byte) {
	if l.Verbose {
		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, message, "", "  ")
		if err != nil {
			l.Error(err.Error())
			return
		}
		fmt.Println(string(prettyJSON.String()))
	}
}
