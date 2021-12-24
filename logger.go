// Package logger for logging program process
// This's example of package documentation
// Add new information
package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	timeFormat string
	debug      bool
}

// Instance function get instance of logger
// Example for change doc
func Instance(timeFormat string, debug bool) *Logger {
	return &Logger{
		timeFormat: timeFormat,
		debug:      debug,
	}
}

// Log function write log in instance.
func (l Logger) Log(level string, s string) {
	switch level {
	case "info", "warning":
		if !l.debug {
			return
		}
	}
	fmt.Printf("[%s] %s %s\n", level, time.Now().Format(l.timeFormat), s)
}
