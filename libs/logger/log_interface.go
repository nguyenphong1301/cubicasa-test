package logger

import (
	"log"
)

//LoggerConfig - logger configs
type LogConfig struct {
	FileName    string
	Level       int
	MaxSize     int64
	MaxDays     int
	KeepConsole bool
}

//Logger interface
type LogInterface interface {
	Init(conf LogConfig)
	Level(level int)
	Debug(v ...interface{})
	Info(v ...interface{})
	Notice(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Alert(v ...interface{})
	Emergency(v ...interface{})
	Trace(v ...interface{})
	Print(v ...interface{})
	GetLogger(prefixes ...string) *log.Logger
}
