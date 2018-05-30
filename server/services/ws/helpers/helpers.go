// Package implement any common RPC API helpers.
package helpers

import (
	"iot-stand/libs/logger"
	"iot-stand/server/services/ws/storage"
	"iot-stand/libs/errors_handler"
)

// Send success response to WS with receipt field.
func SendSuccess(protocolMetadata map[string]interface{}, owner *storage.StorageItem, data map[string]interface{}) {
	resolveFields := map[string]interface{}{}
	if protocolMetadata != nil {
		for k, v := range protocolMetadata {
			resolveFields[k] = v
		}
	}
	if data != nil {
		for k, v := range data {
			resolveFields[k] = v
		}
	}
	owner.WriteJSON(resolveFields)
}

// Send error message to WS.
func SendError(protocolMetadata map[string]interface{}, owner *storage.StorageItem, code int, details ...map[string]interface{}) {
	if len(details) > 0 {
		errorInstance := errors_handler.GetError(code, details[0])
		logger.GetWS().Error("Call API with error", map[string]interface{}{
			"code":    errorInstance.Code(),
			"message": errorInstance.Error(),
			"details": errorInstance.Details(),
		})
		SendSuccess(protocolMetadata, owner, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    errorInstance.Code(),
				"message": errorInstance.Error(),
				"details": errorInstance.Details(),
			},
		})
	} else {
		errorInstance := errors_handler.GetError(code)
		logger.GetWS().Error("Call API with error", map[string]interface{}{
			"code":    errorInstance.Code(),
			"message": errorInstance.Error(),
		})
		SendSuccess(protocolMetadata, owner, map[string]interface{}{
			"error": map[string]interface{}{
				"code":    errorInstance.Code(),
				"message": errorInstance.Error(),
			},
		})
	}
}
