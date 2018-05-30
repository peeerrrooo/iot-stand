package methods

import (
	"iot-stand/libs/nats"
)

func HiJack(param interface{}) {
	nats.CallMethod("hmi", "hiJack", nil)
}
