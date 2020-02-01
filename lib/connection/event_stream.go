package connection

import (
	"database/sql"

	"gitlab.com/zephinzer/themigrator/lib/log"
)

func NewEventStream() EventStream {
	return EventStream{
		Connection: make(chan *sql.DB, 1),
		Error:      make(chan error, 1),
		Logs:       make(chan log.Entry, 128),
	}
}

type EventStream struct {
	Connection chan *sql.DB
	Error      chan error
	Logs       chan log.Entry
}
