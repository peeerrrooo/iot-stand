package methods

import (
	"iot-stand/libs/nats"
	"iot-stand/libs/logger"
)

func HiJack(param interface{}) {
	nats.CallMethod("hmi", "hiJack", nil)
	logger.GetCloud().Info("Success send 'hiJack'")
}
