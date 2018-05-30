// Package implement TELEMETRY SERVICE.
package telemetry

import (
	"iot-stand/device/services/telemetry/nats_api"
)

func CreateService() {
	go nats_api.GetProvider()
	for {
		select {}
	}
}
