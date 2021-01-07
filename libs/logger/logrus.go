package logger

import (
	"github.com/sirupsen/logrus"
	"log"
)

//BeeLogger -
type LogRus struct {
}

//Init -
func (l *LogRus) Init(config LogConfig) {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	l.Level(config.Level)
}

//Level - update logger level
func (l *LogRus) Level(level int) {
	switch level {
	case 1:
		logrus.SetLevel(logrus.PanicLevel)
	case 2:
		logrus.SetLevel(logrus.FatalLevel)
	case 3:
		logrus.SetLevel(logrus.ErrorLevel)
	case 4:
		logrus.SetLevel(logrus.WarnLevel)
	case 5:
		logrus.SetLevel(logrus.InfoLevel)
	case 6:
		logrus.SetLevel(logrus.DebugLevel)
	default:
		logrus.SetLevel(logrus.TraceLevel)
	}
}

//Debug -
func (l *LogRus) Debug(v ...interface{}) {
	logrus.Debug(v...)
}

//Info -
func (l *LogRus) Info(v ...interface{}) {
	logrus.Info(v...)
}

//Warn -
func (l *LogRus) Warn(v ...interface{}) {
	logrus.Warn(v...)
}

//Error -
func (l *LogRus) Error(v ...interface{}) {
	logrus.Error(v...)
}

//Trace -
func (l *LogRus) Trace(v ...interface{}) {
	logrus.Trace(v...)
}

//Critical -
func (l *LogRus) Critical(v ...interface{}) {
	// method not implement
}

//Alert -
func (l *LogRus) Alert(v ...interface{}) {
	// method not implement
}

//Emergency -
func (l *LogRus) Emergency(v ...interface{}) {
	// method not implement
}

//Notice -
func (l *LogRus) Notice(v ...interface{}) {
	// method not implement
}

func (l *LogRus) GetLogger(prefixes ...string) *log.Logger {
	return nil
}

// Print
func (LogRus) Print(v ...interface{}) {
	logrus.Println(v...)
}
