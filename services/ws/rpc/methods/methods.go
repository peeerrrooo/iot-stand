// Main package implement ALL methods of RPC Api.
package methods

import (
	"iot-stand/services/ws/storage"
	"iot-stand/services/ws/helpers"
)

func sendSuccess(method string, owner *storage.StorageItem, reqID interface{}, result map[string]interface{}) {
	rpcMetadata := map[string]interface{}{
		"method": method,
		"req_id": reqID,
	}
	helpers.SendSuccess(rpcMetadata, owner, map[string]interface{}{
		"result": result,
	})
}

func sendError(method string, owner *storage.StorageItem, reqID interface{}, code int, details ...map[string]interface{}) {
	rpcMetadata := map[string]interface{}{
		"method": method,
		"req_id": reqID,
	}
	helpers.SendError(rpcMetadata, owner, code, details...)
}
