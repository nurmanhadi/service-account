package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogrus() *logrus.Logger {
	log := logrus.New()

	_, err := os.Stat("./log")
	if os.IsNotExist(err) {
		if err := os.Mkdir("./log", os.ModePerm); err != nil {
			panic(err)
		}
	}
	file, err := os.OpenFile("./log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)

	return log
}
