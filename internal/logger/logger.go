package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var Log *log.Logger

func InitLogger() {
	if Log == nil {
		Log = logrus.New()

		file, err := os.OpenFile("server-pulsa-app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			Log.Fatal("failed to open log file: ", err)
		}

		Log.Out = file

		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}
}

func GetLogger() *logrus.Logger {
	return Log
}
