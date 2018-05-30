// Package implement logger logic.
package logger

var logger_config *Logger
var logger_nats *Logger
var logger_ws *Logger
var logger_telegram *Logger
var logger_mongo *Logger

// Return CONFIG Logger.
func GetConfig() *Logger {
	if logger_config == nil {
		logger_config = NewLogger("CONFIG")
	}
	return logger_config
}

// Return NATS Logger.
func GetNats() *Logger {
	if logger_nats == nil {
		logger_nats = NewLogger("NATS")
	}
	return logger_nats
}

// Return WS Logger.
func GetWS() *Logger {
	if logger_ws == nil {
		logger_ws = NewLogger("WS")
	}
	return logger_ws
}

// Return TELEGRAM Logger.
func GetTelegram() *Logger {
	if logger_telegram == nil {
		logger_telegram = NewLogger("TELEGRAM")
	}
	return logger_telegram
}

// Return MONGO Logger.
func GetMongo() *Logger {
	if logger_mongo == nil {
		logger_mongo = NewLogger("MONGO")
	}
	return logger_mongo
}
