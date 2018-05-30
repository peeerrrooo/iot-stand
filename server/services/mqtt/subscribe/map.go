package subscribe

import (
	"iot-stand/libs/mqtt"
	"iot-stand/server/services/mqtt/subscribe/methods"
)

var mapApi mqtt.MapApi = mqtt.MapApi{
	"updateTelemetry": methods.UpdateTelemetry,
}

/**
 * Return api map of function.
 */
func GetMap() mqtt.MapApi {
	return mapApi
}
