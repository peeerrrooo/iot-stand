package methods

import (
	"iot-stand/server/services/ws/storage"
	"iot-stand/server/services/ws/notify"
)

func TestMethodSuccess(owner *storage.StorageItem, reqID interface{}, params map[string]interface{}) {
	sendSuccess("TestMethodSuccess", owner, reqID, map[string]interface{}{
		"example_field": "Any...",
		"our_params":    params,
	})
	notify.TestNotify(owner)
}
