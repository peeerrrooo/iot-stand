package nats_api

import (
	"iot-stand/server/services/mqtt/nats_api/methods"
)

var methodsMap map[string]func(map[string]interface{}) = map[string]func(map[string]interface{}){
	"hiJack": methods.HiJack,
}
