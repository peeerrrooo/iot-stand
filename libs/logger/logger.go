package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
)

// Create new logger.
func NewLogger(prefix string) *Logger {
	log := logrus.New()
	form := prefixed.TextFormatter{
		ForceColors:      true,
		ForceFormatting:  true,
		FullTimestamp:    true,
		QuoteEmptyFields: true,
	}
	log.Formatter = &form

	newLogger := new(Logger)
	newLogger.Logger = log
	newLogger.Prefix = prefix

	return newLogger
}

// Type that encapsulate work with logger INFO/ERROR message
// with custom prefix and other structure fields.
type Logger struct {
	Logger *logrus.Logger
	Prefix string
}

// Print INFO LEVEL message.
func (self *Logger) Info(message string, args ...interface{}) {
	var fields map[string]interface{}
	if len(args) > 0 {
		fields = args[0].(map[string]interface{})
		self.Logger.WithFields(logrus.Fields(makeFieldsWithPrefix(self.Prefix, fields))).Info(message)
	} else {
		self.Logger.WithFields(logrus.Fields{
			"prefix": self.Prefix,
		}).Info(message)
	}
}

// Print ERROR LEVEL message.
func (self *Logger) Error(message string, args ...interface{}) {
	var fields map[string]interface{}
	if len(args) > 0 {
		fields = args[0].(map[string]interface{})
		self.Logger.WithFields(logrus.Fields(makeFieldsWithPrefix(self.Prefix, fields))).Error(message)
	} else {
		self.Logger.WithFields(logrus.Fields{
			"prefix": self.Prefix,
		}).Info(message)
	}
}

// Get logger with one custom field - prefix.
func (self *Logger) GetPrefixedLogger() *logrus.Logger {
	return self.Logger.WithFields(logrus.Fields(map[string]interface{}{
		"prefix": self.Prefix,
	})).Logger
}

// Get result fields with custom prefix for prepare logger.
func makeFieldsWithPrefix(prefix string, fields map[string]interface{}) map[string]interface{} {
	newFields := make(map[string]interface{})
	for k, v := range fields {
		newFields[k] = v
	}
	newFields["prefix"] = prefix
	return newFields
}
