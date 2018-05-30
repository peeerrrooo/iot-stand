package methods

import (
	"iot-stand/server/services/ws/storage"
	"iot-stand/libs/nats"
)

func HiJack(owner *storage.StorageItem, reqID interface{}, params map[string]interface{}) {
	sendSuccess("hiJack", owner, reqID, map[string]interface{}{
		"is_success": true,
	})

	nats.CallMethod("mqtt", "hiJack", map[string]interface{}{})
}
