package log

import "github.com/usvc/logger"

var Level = map[int]logger.Level{
	5: logger.LevelTrace,
	4: logger.LevelDebug,
	3: logger.LevelInfo,
	2: logger.LevelWarn,
	1: logger.LevelError,
}

var Format = map[string]logger.Format{
	"json": logger.FormatJSON,
	"text": logger.FormatText,
}

type Options struct {
	Level  int
	Format string
}
