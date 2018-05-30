// Package implement MQTT server.
package mqtt

import (
	"iot-stand/server/services/mqtt/nats_api"
	"iot-stand/server/services/mqtt/subscribe"
)

func CreateService() {
	go nats_api.GetProvider()
	go subscribe.CreateSubscribeService()
	for {
		select {}
	}
}
