// Package implement WEBSOCKET NATS API.
package websocket_nats_api

import (
	"iot-stand/libs/nats"
)

var provider *nats.Subscribe = nil

func GetProvider() *nats.Subscribe {
	if provider == nil {
		provider = nats.InitSubscribe("websocket", methodsMap)
	}
	return provider
}
