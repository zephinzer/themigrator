package connection

import (
	"fmt"
	"time"

	"gitlab.com/zephinzer/themigrator/lib/log"
)

func Establish(connectionOptions Options, eventStream EventStream, retryInterval time.Duration, retries ...int) {
	var numberOfRetries int = 10
	if len(retries) > 0 {
		numberOfRetries = retries[0]
	}
	for {
		numberOfRetries -= 1
		dbConnection, dbConnectionError := Create(connectionOptions)
		if dbConnectionError != nil {
			if numberOfRetries == 0 {
				eventStream.Logs <- log.Entry{
					Code:    ErrorDBConnectionFailed,
					Level:   log.LevelError,
					Message: dbConnectionError.Error(),
				}
				eventStream.Error <- dbConnectionError
				break
			} else {
				eventStream.Logs <- log.Entry{
					Code:    ErrorDBConnectionFailed,
					Level:   log.LevelWarn,
					Message: fmt.Sprintf("%s, retrying again in %s (retries left: %v)...", dbConnectionError, retryInterval, numberOfRetries),
				}
			}
			<-time.After(retryInterval)
		} else {
			eventStream.Logs <- log.Entry{
				Code:    ErrorDBConnectionSucceeded,
				Level:   log.LevelInfo,
				Message: "connection succeeded",
			}
			eventStream.Connection <- dbConnection
			break
		}
	}
}
