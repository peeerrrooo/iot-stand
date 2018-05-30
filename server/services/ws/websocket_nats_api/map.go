package websocket_nats_api

import (
	"iot-stand/server/services/ws/websocket_nats_api/methods"
)

var methodsMap map[string]func(map[string]interface{}) = map[string]func(map[string]interface{}){
	"getWsMessage": methods.GetWsMessage,
}
