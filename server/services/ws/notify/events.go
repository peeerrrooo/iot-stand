package notify

import (
	"iot-stand/server/services/ws/topic"
	"iot-stand/server/services/ws/storage"
)

func TestNotify(owner *storage.StorageItem) {
	sendNotify("testNotify", owner, topic.COMMON_TOPIC, map[string]interface{}{
		"any_field": "any value",
	})
}

func UpdateTelemetry(telemetry []map[string]interface{}) {
	sendNotify("updateTelemetry", nil, topic.COMMON_TOPIC, map[string]interface{}{
		"telemetry": telemetry,
	})
}
