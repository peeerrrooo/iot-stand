// Package implement logger logic.
package logger

var logger_config *Logger
var logger_nats *Logger
var logger_ws *Logger
var logger_mqtt_service *Logger
var logger_cloud *Logger
var logger_mqtt *Logger
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

// Return WS_SERVICE Logger.
func GetWS() *Logger {
	if logger_ws == nil {
		logger_ws = NewLogger("WS")
	}
	return logger_ws
}

// Return MQTT_SERVICE Logger.
func GetMQTTService() *Logger {
	if logger_mqtt_service == nil {
		logger_mqtt_service = NewLogger("MQTT-SERVICE")
	}
	return logger_mqtt_service
}

// Return CLOUD Logger.
func GetCloud() *Logger {
	if logger_cloud == nil {
		logger_cloud = NewLogger("CLOUD")
	}
	return logger_cloud
}

// Return MQTT Logger.
func GetMQTT() *Logger {
	if logger_mqtt == nil {
		logger_mqtt = NewLogger("MQTT")
	}
	return logger_mqtt
}

// Return MONGO Logger.
func GetMongo() *Logger {
	if logger_mongo == nil {
		logger_mongo = NewLogger("MONGO")
	}
	return logger_mongo
}
