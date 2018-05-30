package rpc

import (
	"iot-stand/services/ws/storage"
	"iot-stand/services/ws/rpc/methods"
)

var mapApi map[string]func(*storage.StorageItem, interface{}, map[string]interface{}) = map[string]func(*storage.StorageItem, interface{}, map[string]interface{}){
	"testMethodSuccess": methods.TestMethodSuccess,
	"testMethodError":   methods.TestMethodError,
	"sendFeedback":      methods.SendFeedback,
}
