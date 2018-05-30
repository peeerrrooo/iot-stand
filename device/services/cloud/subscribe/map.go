package subscribe

import (
	"iot-stand/libs/mqtt"
	"iot-stand/device/services/cloud/subscribe/methods"
)

var mapApi mqtt.MapApi = mqtt.MapApi{
	"hiJack": methods.HiJack,
}

/**
 * Return api map of function.
 */
func GetMap() mqtt.MapApi {
	return mapApi
}
