package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/rxwen/go-logrus-test/cal"
)

func main() {
	log.SetFormatter(&log.TextFormatter{}) // ..JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.WithField("field1", "filed1 value").Info("log 1")

	//log.SetLevel(log.WarnLevel)
	log.WithField("result", cal.Add(12, 30)).Debug("add")
	log.WithField("result", cal.Multiply(12, 3)).Warn("multiply")
}
