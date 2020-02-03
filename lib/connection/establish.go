package connection

import (
	"fmt"
	"time"

	"github.com/usvc/logger"
	"gitlab.com/zephinzer/themigrator/lib/log"
)

// Establish establishes a new database connection
func Establish(connectionOptions Options, eventStream EventStream, retryInterval time.Duration, retries ...int) {
	var numberOfRetries int = 10
	if len(retries) > 0 {
		numberOfRetries = retries[0]
	}
	for {
		numberOfRetries--
		dbConnection, dbConnectionError := Create(connectionOptions)
		if dbConnectionError != nil {
			if numberOfRetries == 0 {
				eventStream.Logs <- log.NewEntry(logger.LevelError, ErrorDBConnectionFailed, "ran out of retries", dbConnectionError)
				eventStream.Error <- dbConnectionError
				break
			} else {
				eventStream.Logs <- log.NewEntry(logger.LevelWarn, ErrorDBConnectionFailed, fmt.Sprintf("connection failed, retrying again in %s (retries left: %v)...", retryInterval, numberOfRetries), dbConnectionError)
			}
			<-time.After(retryInterval)
		} else {
			eventStream.Logs <- log.NewEntry(logger.LevelInfo, ErrorDBConnectionSucceeded, "connection succeeded")
			eventStream.Connection <- dbConnection
			break
		}
	}
}
