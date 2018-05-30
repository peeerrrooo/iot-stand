package nats_api

import (
	"iot-stand/server/services/ws/nats_api/methods"
)

var methodsMap map[string]func(map[string]interface{}) = map[string]func(map[string]interface{}){
	"updateTelemetry": methods.UpdateTelemetry,
}
