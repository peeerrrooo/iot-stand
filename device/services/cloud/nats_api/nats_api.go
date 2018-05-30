// Package implement NATS API.
package nats_api

import (
	"iot-stand/libs/nats"
)

var provider *nats.Subscribe = nil

func GetProvider() *nats.Subscribe {
	if provider == nil {
		provider = nats.InitSubscribe("cloud", methodsMap)
	}
	return provider
}
