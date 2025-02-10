package profile

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

type DefaultLogger struct {
	level  LogLevel
	writer io.Writer
}

func NewLogger(level string) *DefaultLogger {
	// Check for AZTX_LOG_FILE environment variable
	logFile := os.Getenv("AZTX_LOG_FILE")
	var writer io.Writer = os.Stderr

	if logFile != "" {
		if f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
			writer = f
		}
	}

	return &DefaultLogger{
		level:  parseLevel(level),
		writer: writer,
	}
}

func parseLevel(level string) LogLevel {
	switch strings.ToLower(level) {
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "warn":
		return LevelWarn
	case "error":
		return LevelError
	default:
		return LevelInfo
	}
}

func (l *DefaultLogger) log(level LogLevel, msg string, args ...interface{}) {
	if level < l.level {
		return
	}

	levelStr := [...]string{"DEBUG", "INFO", "WARN", "ERROR"}
	formatted := fmt.Sprintf(msg, args...)
	logLine := fmt.Sprintf("[AZTX][%s] %s\n", levelStr[level], formatted)
	fmt.Fprint(l.writer, logLine)
}

func (l *DefaultLogger) Debug(msg string, args ...interface{}) {
	l.log(LevelDebug, msg, args...)
}

func (l *DefaultLogger) Info(msg string, args ...interface{}) {
	l.log(LevelInfo, msg, args...)
}

func (l *DefaultLogger) Warn(msg string, args ...interface{}) {
	l.log(LevelWarn, msg, args...)
}

func (l *DefaultLogger) Error(msg string, args ...interface{}) {
	l.log(LevelError, msg, args...)
}
