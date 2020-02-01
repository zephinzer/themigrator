package common

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/zephinzer/themigrator/lib/log"
)

var logger = logrus.New()

func SetLogLevel(logLevel log.Level) {
	levels := map[log.Level]logrus.Level{
		log.LevelTrace: logrus.TraceLevel,
		log.LevelDebug: logrus.DebugLevel,
		log.LevelInfo:  logrus.InfoLevel,
		log.LevelWarn:  logrus.WarnLevel,
		log.LevelError: logrus.ErrorLevel,
	}
	logger.SetLevel(levels[logLevel])
}

type Printer func(...interface{})

func HandleLogs(logs chan log.Entry) {
	printLog := map[log.Level]Printer{
		log.LevelTrace: logger.Trace,
		log.LevelDebug: logger.Debug,
		log.LevelInfo:  logger.Info,
		log.LevelWarn:  logger.Warn,
		log.LevelError: logger.Error,
	}
	for logEntry := range logs {
		printLog[logEntry.GetLevel()](logEntry.GetCode(), "| ", logEntry.GetMessage())
		if logEntry.GetData() != nil {
			printLog[logEntry.GetLevel()](logEntry.GetData())
		}
	}

}
