// Package implement CLOUD SERVICE.
package cloud

import (
	"iot-stand/device/services/cloud/nats_api"
	"iot-stand/device/services/cloud/subscribe"
)

func CreateService() {
	go nats_api.GetProvider()
	go subscribe.CreateSubscribeService()
	for {
		select {}
	}
}
