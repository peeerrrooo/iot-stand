// Package implement NOTIFY protocol for WS.
package notify

import (
	"iot-stand/server/services/ws/storage"
	"iot-stand/libs/nats"
	"iot-stand/libs/logger"
)

func sendNotify(event string, owner *storage.StorageItem, topic string, fields ... map[string]interface{}) {
	if len(fields) > 0 {
		if owner != nil {
			publishWsTopic(topic, map[string]interface{}{
				"event":  event,
				"fields": fields[0],
			}, owner)
		} else {
			publishWsTopic(topic, map[string]interface{}{
				"event":  event,
				"fields": fields[0],
			})
		}
	} else {
		if owner != nil {
			publishWsTopic(topic, map[string]interface{}{
				"event": event,
			}, owner)
		} else {
			publishWsTopic(topic, map[string]interface{}{
				"event": event,
			})
		}
	}
	logger.GetWS().Info("Send notify message", map[string]interface{}{
		"event": event,
	})
}

// Publish data with topic for WS to NATS server.
func publishWsTopic(topic string, data map[string]interface{}, excludeOwners ...*storage.StorageItem) {
	excludeOwnersID := []string{}
	for _, v := range excludeOwners {
		excludeOwnersID = append(excludeOwnersID, v.UUID)
	}
	if len(excludeOwnersID) > 0 {
		nats.CallMethod("websocket", "getWsMessage", map[string]interface{}{
			"topic":          topic,
			"data":           data,
			"exclude_owners": excludeOwnersID,
		})
	} else {
		nats.CallMethod("websocket", "getWsMessage", map[string]interface{}{
			"topic": topic,
			"data":  data,
		})
	}
}
