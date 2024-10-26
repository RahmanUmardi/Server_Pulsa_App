package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var Log *log.Logger

type Logger struct {
	log *log.Logger
}

func NewLogger() Logger {

	log := logrus.New()

	file, err := os.OpenFile("server-pulsa-app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Fatal("failed to open log file: ", err)
	}

	log.Out = file

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})

	return Logger{log: log}

}
func (l *Logger) Info(message string, data any) {
	l.log.WithFields(logrus.Fields{
		"data": data,
	}).Info(message)
}

func (l *Logger) Error(message string, data any) {
	l.log.WithFields(logrus.Fields{
		"data": data,
	}).Error(message)
}
