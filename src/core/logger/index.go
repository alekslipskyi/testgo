package logger

import (
	"fmt"
	"os"
	"strings"
)

const (
	WARN  string = "WARN"
	LOG   string = "LOG"
	ERROR string = "ERROR"
	ALL   string = "DEBUG"
)

type Logger struct {
	Context string
}

func (logger *Logger) isModeAllowed(mode string) bool {
	currentMode := os.Getenv("LOGGER_LEVEL")

	if currentMode == ALL {
		return true
	}

	return strings.Contains(currentMode, mode)
}

func (logger *Logger) Warn(data ...interface{}) {
	if logger.isModeAllowed(WARN) {
		fmt.Println(fmt.Sprintf("[%s] warn:", logger.Context), data)
	}
}

func (logger *Logger) Error(data ...interface{}) {
	if logger.isModeAllowed(ERROR) {
		fmt.Println(fmt.Errorf("[%s] error:", logger.Context), data)
	}
}

func (logger *Logger) Debug(data ...interface{}) {
	if logger.isModeAllowed(ALL) {
		fmt.Println(fmt.Errorf("[%s] debug:", logger.Context), data)
	}
}

func (logger *Logger) Log(data ...interface{}) {
	if logger.isModeAllowed(LOG) {
		fmt.Println(fmt.Sprintf("[%s] log:", logger.Context), data)
	}
}
