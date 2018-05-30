// Package implement NATS API.
package nats_api

import (
	"iot-stand/services/ws/storage"
	"iot-stand/libs/nats"
)

var provider *nats.Subscribe = nil

func GetProvider() *nats.Subscribe {
	if provider == nil {
		provider = nats.InitSubscribe("ws", methodsMap)
	}
	return provider
}

// Publish data with topic for WS to NATS server.
func PublishWsTopic(topic string, data map[string]interface{}, excludeOwners ...*storage.StorageItem) {
	excludeOwnersID := []string{}
	for _, v := range excludeOwners {
		excludeOwnersID = append(excludeOwnersID, v.UUID)
	}
	if len(excludeOwnersID) > 0 {
		GetProvider().Send("ws", "getWSMessage", map[string]interface{}{
			"topic":          topic,
			"data":           data,
			"exclude_owners": excludeOwnersID,
		})
	} else {
		GetProvider().Send("ws", "getWSMessage", map[string]interface{}{
			"topic": topic,
			"data":  data,
		})
	}
}
