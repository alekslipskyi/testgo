package logger

import (
	"constants/dbError"
	"fmt"
	"os"
	"strings"
)

const YELLOW = "\033[1;33m"
const RED = "\033[1;31m"
const WHITE = "\033[0;39m"
const GREEN = "\033[1;32m"

const (
	WARN  string = "WARN"
	LOG   string = "LOG"
	ERROR string = "ERROR"
	ALL   string = "DEBUG"
)

const (
	WarningColor = YELLOW
	ErrorColor   = RED
	DebugColor   = WHITE
	LogColor     = GREEN
)

type Colors struct {
	Error string
	Debug string
	Info  string
	Warn  string
}

type Logger struct {
	Context string
	Colors  Colors
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
		color := WarningColor

		if len(logger.Colors.Warn) > 0 {
			color = logger.Colors.Warn
		}

		fmt.Println(fmt.Sprintf("%s[%s] warning: %s%s", color, strings.ToUpper(logger.Context), fmt.Sprint(data...), "\033[0m"))
	}
}

func (logger *Logger) Error(data ...interface{}) {
	if logger.isModeAllowed(ERROR) {
		color := ErrorColor

		if len(logger.Colors.Error) > 0 {
			color = logger.Colors.Error
		}
		fmt.Println(fmt.Sprintf("%s[%s] error:%s%s", color, strings.ToUpper(logger.Context), fmt.Sprint(data...), "\033[0m"))

		if os.Getenv("ENV") == "DEV" || os.Getenv("ENV") == "TEST" {
			os.Exit(1)
		}
	}
}

func (logger *Logger) LogOnError(err error, message ...interface{}) {
	if err != nil && err.Error() != dbError.NO_ROWS {
		logger.Error(message)
	}
}

func (logger *Logger) FailOnerror(err error, data ...interface{}) {
	if err != nil {
		logger.Error(err, data[0])
		os.Exit(1)
	}
}

func (logger *Logger) Debug(data ...interface{}) {
	if logger.isModeAllowed(ALL) {
		color := DebugColor

		if len(logger.Colors.Debug) > 0 {
			color = logger.Colors.Debug
		}

		fmt.Println(fmt.Sprintf("%s[%s] debug:%s%s", color, strings.ToUpper(logger.Context), fmt.Sprint(data...), "\033[0m"))
	}
}

func (logger *Logger) Info(data ...interface{}) {
	if logger.isModeAllowed(LOG) {
		color := LogColor

		if len(logger.Colors.Info) > 0 {
			color = logger.Colors.Info
		}

		fmt.Println(fmt.Sprintf("\n %s[%s] %s%s \n", color, strings.ToUpper(logger.Context), fmt.Sprint(data...), "\033[0m"))
	}
}
