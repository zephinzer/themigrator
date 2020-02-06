package log

import (
	"github.com/usvc/logger"
)

// Handle handles the processing of a logs channel
func Handle(loggerInstance logger.Logger, logEntries chan Entry) {
	printLog := map[logger.Level]printer{
		logger.LevelTrace: noop,
		logger.LevelDebug: noop,
		logger.LevelInfo:  noop,
		logger.LevelWarn:  noop,
		logger.LevelError: noop,
	}

	if loggerInstance != nil {
		printLog = map[logger.Level]printer{
			logger.LevelTrace: loggerInstance.Trace,
			logger.LevelDebug: loggerInstance.Debug,
			logger.LevelInfo:  loggerInstance.Info,
			logger.LevelWarn:  loggerInstance.Warn,
			logger.LevelError: loggerInstance.Error,
		}
	}
	for logEntry := range logEntries {
		printLog[logEntry.GetLevel()]("[", logEntry.GetCode(), "] ", logEntry.GetMessage())
		if logEntry.GetData() != nil {
			printLog[logEntry.GetLevel()](logEntry.GetData())
		}
	}
}

// printer represents a logging function
type printer func(...interface{})

func noop(...interface{}) {}
