package rpc

import (
	"iot-stand/server/services/ws/storage"
	"iot-stand/server/services/ws/rpc/methods"
)

var mapApi map[string]func(*storage.StorageItem, interface{}, map[string]interface{}) = map[string]func(*storage.StorageItem, interface{}, map[string]interface{}){
	"testMethodSuccess": methods.TestMethodSuccess,
	"testMethodError":   methods.TestMethodError,
	"getTelemetry":      methods.GetTelemetry,
	"removeTelemetry":   methods.RemoveTelemetry,
	"hiJack":            methods.HiJack,
}
