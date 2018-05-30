// Package implement RPC protocol for WS.
package rpc

import (
	"iot-stand/services/ws/storage"
	"iot-stand/libs/logger"
	"iot-stand/services/ws/helpers"
)

var rpc_instance *RPC

// Type that encapsulate work with RPC API Methods.
type RPC struct {
	Storage *storage.Storage
	Map map[string]func(*storage.StorageItem, interface{}, map[string]interface{})
}

// Read input WS message and routing method/receipt/params to api.
func (self *RPC) ReadProcess(method string, receipt interface{},
	owner *storage.StorageItem, params map[string]interface{}) {
	methodApi, okMethod := self.Map[method]
	if !okMethod {
		helpers.SendError(nil, owner, 1003)
		return
	}
	methodApi(owner, receipt, params)
	logger.GetWS().Info("Finish calling API", map[string]interface{}{
		"method":  method,
		"receipt": receipt,
		"params":  params,
	})
}

// Return singleton for RPC.
func GetRPC() *RPC {
	if rpc_instance == nil {
		rpc_instance = new(RPC)
		rpc_instance.Storage = storage.GetStorage()
		rpc_instance.Map = mapApi
	}
	return rpc_instance
}
