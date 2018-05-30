package methods

import (
	"iot-stand/libs/mqtt"
	"iot-stand/libs/logger"
)

func UpdateTelemetry(data map[string]interface{}) {
	mqtt.Publish("updateTelemetry", data)
	logger.GetCloud().Info("Success send telemetry")
}
