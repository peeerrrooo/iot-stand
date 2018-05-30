package methods

import (
	"iot-stand/libs/nats"
	"time"
	"math/rand"
	"iot-stand/libs/logger"
)

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func UpdateTelemetry(data map[string]interface{}) {
	nats.CallMethod("cloud", "updateTelemetry", map[string]interface{}{
		"battery":     random(1, 100),
		"total_range": random(1, 1000),
		"mileage":     random(1, 1000),
		"temperature": random(1, 120),
		"vin":         "54Idgkjqpd",
	})
	logger.GetTelemetry().Info("Success send telemetry")
}
