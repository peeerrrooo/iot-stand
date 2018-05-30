package methods

import (
	"iot-stand/libs/json_codec"
	"iot-stand/libs/nats"
	"iot-stand/libs/mongo"
	"iot-stand/libs/mongo/schema"
	"time"
	"iot-stand/libs/logger"
)

func UpdateTelemetry(param interface{}) {
	data, okData := param.(map[string]interface{})
	if !okData {
		return
	}
	battery, err := json_codec.GetInt64("battery", data)
	if err != nil {
		logger.GetMQTTService().Error("Cannot get 'battery' from telemetry", map[string]interface{}{
			"method": "updateTelemetry",
			"param":  data,
			"err":    err,
		})
		return
	}
	totalRange, err := json_codec.GetInt64("total_range", data)
	if err != nil {
		logger.GetMQTTService().Error("Cannot get 'total_range' from telemetry", map[string]interface{}{
			"method": "updateTelemetry",
			"param":  data,
			"err":    err,
		})
		return
	}
	temperature, err := json_codec.GetInt64("temperature", data)
	if err != nil {
		logger.GetMQTTService().Error("Cannot get 'temperature' from telemetry", map[string]interface{}{
			"method": "updateTelemetry",
			"param":  data,
			"err":    err,
		})
		return
	}
	mileage, err := json_codec.GetInt64("mileage", data)
	if err != nil {
		logger.GetMQTTService().Error("Cannot get 'mileage' from telemetry", map[string]interface{}{
			"method": "updateTelemetry",
			"param":  data,
			"err":    err,
		})
		return
	}
	vin, err := json_codec.GetString("vin", data)
	if err != nil {
		logger.GetMQTTService().Error("Cannot get 'vin' from telemetry", map[string]interface{}{
			"method": "updateTelemetry",
			"param":  data,
			"err":    err,
		})
		return
	}
	mongo.GetMongoConnection().AddTelemetry(schema.Telemetry{
		Battery:     battery,
		TotalRange:  totalRange,
		Mileage:     mileage,
		Temperature: temperature,
		Vin:         vin,
		Created:     time.Now(),
	})
	nats.CallMethod("ws", "updateTelemetry", nil)
}
