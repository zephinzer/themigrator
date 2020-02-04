package log

import (
	"fmt"

	"github.com/usvc/logger"
)

// NewEntry constructs a new log entry and returns it
func NewEntry(level logger.Level, code, message string, data ...interface{}) Entry {
	var entryData interface{}
	if len(data) == 0 {
		entryData = nil
	} else if len(data) == 1 {
		entryData = data[0]
	} else {
		entryData = data
	}
	return logEntry{
		Level:   level,
		Code:    code,
		Message: message,
		Data:    entryData,
	}
}

// Entry represents a log entry
type Entry interface {
	Error() string
	GetCode() string
	GetData() interface{}
	GetLevel() logger.Level
	GetMessage() string
}

// logEntry represents a log entry
type logEntry struct {
	Code    string
	Data    interface{}
	Level   logger.Level
	Message string
}

// GetCode retrieves the code attached to this log.Entry instance
func (le logEntry) GetCode() string {
	if len(le.Code) == 0 {
		return "UNKNOWN_CODE"
	}
	return le.Code
}

// GetMessage retrieves the message attached to this log.Entry instance
func (le logEntry) GetMessage() string {
	return le.Message
}

// GetLevel retrieves the logger.Level attached to this log.Entry instance
func (le logEntry) GetLevel() logger.Level {
	if len(le.Level) == 0 {
		return logger.LevelInfo
	}
	return le.Level
}

// GetData retrieves any data objects attached to this log.Entry instance
func (le logEntry) GetData() interface{} {
	return le.Data
}

// Error returns a string to comply with the `error` interface
func (le logEntry) Error() string {
	toString := fmt.Sprintf(
		"%s: %s",
		le.GetCode(),
		le.GetMessage(),
	)
	if le.Data != nil {
		toString = fmt.Sprintf("%s, data: %v", toString, le.GetData())
	}
	return toString
}
