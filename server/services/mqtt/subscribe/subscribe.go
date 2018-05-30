package subscribe

import (
	"iot-stand/libs/logger"
	"iot-stand/libs/mqtt"
)

/**
 * Create Subscribe service.
 */
func CreateSubscribeService() {
	logger.GetMQTTService().Info("Start MQTT subscribe")
	mqtt.Subscribe(5, GetMap())
}
