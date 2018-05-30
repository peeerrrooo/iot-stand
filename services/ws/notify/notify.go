// Package implement NOTIFY protocol for WS.
package notify

import (
	"iot-stand/services/ws/storage"
	"iot-stand/services/ws/nats_api"
)

func sendNotify(event string, owner *storage.StorageItem, topic string, fields ... map[string]interface{}) {
	if len(fields) > 0 {
		nats_api.PublishWsTopic(topic, map[string]interface{}{
			"event":  event,
			"fields": fields,
		}, owner)
	} else {
		nats_api.PublishWsTopic(topic, map[string]interface{}{
			"event": event,
		}, owner)
	}
}
