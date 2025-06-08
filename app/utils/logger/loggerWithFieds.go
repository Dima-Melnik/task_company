package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogWithFieldsFatalf(component string, action string, err error) {
	logrus.WithFields(logrus.Fields{
		"component": component,
		"action":    action,
		"error":     err,
	}).Fatalf("Error: %v", err)
}

func LogWithFieldsErrorf(component string, action string, err error) {
	logrus.WithFields(logrus.Fields{
		"component": component,
		"action":    action,
		"error":     err,
	}).Errorf("Error: %v", err)
}

func LogService(c *gin.Context, action string, message string, err error) {
	logrus.WithFields(logrus.Fields{
		"component": "service",
		"action":    action,
		"method":    c.Request.Method,
		"host":      c.Request.Host,
		"error":     err,
	}).Errorf("Error in service %s", message)
}

func LogHandler(c *gin.Context, message string, err error) {
	logrus.WithFields(logrus.Fields{
		"component": "handler",
		"method":    c.Request.Method,
		"error":     err,
	}).Errorf("Error in handler %s", message)
}
