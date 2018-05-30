package methods

import (
	"iot-stand/libs/logger"
	"iot-stand/libs/mqtt"
)

func HiJack(data map[string]interface{}) {
	err := mqtt.Publish("hiJack", nil)
	if err != nil {
		logger.GetMQTTService().Error("Error send 'HiJack")
		return
	}
	logger.GetMQTTService().Info("Success send 'HiJack'")
}
