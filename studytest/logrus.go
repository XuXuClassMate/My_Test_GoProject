package main

import (
	"github.com/sirupsen/logrus"
)

func logMain() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		//DisableColors: true,
		FullTimestamp: true,
	})

	log.Info("This is an info log message")
	log.Warn("This is a warning log message")
	log.Error("This is an error log message")

}
