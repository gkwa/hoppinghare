package log

import (
	"fmt"
	"os"
)

// LogLevel represents the verbosity level of logging
type LogLevel int

const (
	// LevelError is the default log level
	LevelError LogLevel = iota
	// LevelWarn is used for warnings
	LevelWarn
	// LevelInfo is used for information
	LevelInfo
	// LevelDebug is used for debugging
	LevelDebug
)

var currentLevel = LevelError

// SetLevel sets the current log level
func SetLevel(level LogLevel) {
	currentLevel = level
}

// IncreaseLevel increases the verbosity by one level
func IncreaseLevel() {
	if currentLevel < LevelDebug {
		currentLevel++
	}
}

// Error logs an error message
func Error(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

// Warn logs a warning message if the log level is at least LevelWarn
func Warn(format string, args ...interface{}) {
	if currentLevel >= LevelWarn {
		fmt.Fprintf(os.Stderr, "WARN: "+format+"\n", args...)
	}
}

// Info logs an informational message if the log level is at least LevelInfo
func Info(format string, args ...interface{}) {
	if currentLevel >= LevelInfo {
		fmt.Fprintf(os.Stderr, "INFO: "+format+"\n", args...)
	}
}

// Debug logs a debug message if the log level is at least LevelDebug
func Debug(format string, args ...interface{}) {
	if currentLevel >= LevelDebug {
		fmt.Fprintf(os.Stderr, "DEBUG: "+format+"\n", args...)
	}
}

