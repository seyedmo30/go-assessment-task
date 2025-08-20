package pkg

import (
	"github.com/sirupsen/logrus"
)

func InitLog() {
	logrus.SetReportCaller(true)
	logLevel := Config.LogLevel
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("log level is not valid")
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
}
