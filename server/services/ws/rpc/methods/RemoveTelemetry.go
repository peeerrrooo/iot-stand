package methods

import (
	"iot-stand/server/services/ws/storage"
	"iot-stand/libs/mongo"
)

func RemoveTelemetry(owner *storage.StorageItem, reqID interface{}, params map[string]interface{}) {
	err := mongo.GetMongoConnection().PurgeTelemetry()
	if err != nil {
		return
	}

	sendSuccess("removeTelemetry", owner, reqID, map[string]interface{}{
		"success": true,
	})
}
