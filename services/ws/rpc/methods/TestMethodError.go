package methods

import (
	"iot-stand/services/ws/storage"
)

func TestMethodError(owner *storage.StorageItem, reqID interface{}, params map[string]interface{}) {
	sendError("TestMethodError", owner, reqID, 1004, map[string]interface{}{
		"any_detail_field": "any error detail field value",
	})
}
