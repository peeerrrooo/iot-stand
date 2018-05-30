package notify

import (
	"iot-stand/services/ws/topic"
	"iot-stand/services/ws/storage"
)

func TestNotify(owner *storage.StorageItem) {
	sendNotify("testNotify", owner, topic.COMMON_TOPIC, map[string]interface{}{
		"any_field": "any value",
	})
}
