package methods

import (
	"iot-stand/server/services/ws/storage"
	"iot-stand/libs/mongo"
)

func GetTelemetry(owner *storage.StorageItem, reqID interface{}, params map[string]interface{}) {
	telemetry, err := mongo.GetMongoConnection().GetTelemetry()
	if err != nil {
		return
	}

	var resolveData []map[string]interface{} = make([]map[string]interface{}, 0)
	for _, v := range telemetry {
		resolveData = append(resolveData, map[string]interface{}{
			"battery":     v.Battery,
			"total_range": v.TotalRange,
			"temperature": v.Temperature,
			"mileage":     v.Mileage,
			"vin":         v.Vin,
			"created":     v.Created,
		})
	}

	sendSuccess("getTelemetry", owner, reqID, map[string]interface{}{
		"telemetry": resolveData,
	})
}
