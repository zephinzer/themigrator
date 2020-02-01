package log

type Level string

const (
	LevelTrace Level = "TRAC"
	LevelDebug Level = "DEBU"
	LevelInfo  Level = "INFO"
	LevelWarn  Level = "WARN"
	LevelError Level = "ERRO"
)
