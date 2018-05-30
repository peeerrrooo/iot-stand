package methods

import (
	"iot-stand/services/ws/storage"
	"iot-stand/libs/json_codec"
	"iot-stand/services/ws/nats_api"
)

func SendFeedback(owner *storage.StorageItem, reqID interface{}, params map[string]interface{}) {
	name, err := json_codec.GetString("name", params)
	if err != nil {
		sendError("sendFeedback", owner, reqID, 1005, map[string]interface{}{
			"field": "name",
		})
	}
	email, err := json_codec.GetString("email", params)
	if err != nil {
		sendError("sendFeedback", owner, reqID, 1005, map[string]interface{}{
			"field": "email",
		})
	}
	description, err := json_codec.GetString("description", params)
	if err != nil {
		sendError("sendFeedback", owner, reqID, 1005, map[string]interface{}{
			"field": "description",
		})
	}

	sendSuccess("sendFeedback", owner, reqID, map[string]interface{}{
		"success":     true,
		"name":        name,
		"email":       email,
		"description": description,
	})

	nats_api.GetProvider().Send("telegram", "sendFeedback", map[string]interface{}{
		"name":        name,
		"email":       email,
		"description": description,
	})
}
