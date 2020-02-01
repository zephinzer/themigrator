package log

import "fmt"

type Entry struct {
	Code    string
	Message string
	Level   Level
	Data    interface{}
}

func (e Entry) GetCode() string {
	if len(e.Code) == 0 {
		return "UNKNOWN_CODE"
	}
	return e.Code
}

func (e Entry) GetMessage() string {
	return e.Message
}

func (e Entry) GetLevel() Level {
	if len(e.Level) == 0 {
		return LevelInfo
	}
	return e.Level
}

func (e Entry) GetData() interface{} {
	return e.Data
}

func (e Entry) Error() string {
	toString := fmt.Sprintf(
		"[%s] %s: %s",
		e.GetLevel(),
		e.GetCode(),
		e.GetMessage(),
	)
	if e.Data != nil {
		toString = fmt.Sprintf(
			"%s, data: %v",
			toString,
			e.GetData(),
		)
	}
	return toString
}
