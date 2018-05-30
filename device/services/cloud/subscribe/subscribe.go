package subscribe

import (
	"iot-stand/libs/logger"
	"iot-stand/libs/mqtt"
)

/**
 * Create Subscribe service.
 */
func CreateSubscribeService() {
	logger.GetCloud().Info("Start MQTT subscribe")
	mqtt.Subscribe(5, GetMap())
}
