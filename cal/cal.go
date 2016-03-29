package cal

import (
	log "github.com/Sirupsen/logrus"
)

func Add(a, b int) int {
	log.WithFields(log.Fields{
		"a": a,
		"b": b,
	}).Warn("add")
	return a + b
}

func Multiply(a, b int) int {
	log.WithFields(log.Fields{
		"a": a,
		"b": b,
	}).Debug("multiply")
	return a * b
}
