package nats

import (
	nats_controller "github.com/nats-io/go-nats"
	"fmt"
	"iot-stand/libs/logger"
	"iot-stand/libs/config"
	"iot-stand/libs/json_codec"
)

var connection *nats_controller.Conn = nil

// Get original connection for NATS server.
func GetConnection() *nats_controller.Conn {
	if connection == nil {
		if config.GetBool("docker") {
			nc, err := nats_controller.Connect(fmt.Sprintf("nats://nats:%d", config.GetInt("nats_port")),
				nats_controller.UserInfo(config.GetString("nats_user"), config.GetString("nats_pass")))
			if err != nil {
				logger.GetNats().Error("Error in connection", map[string]interface{}{
					"error": err,
				})
				return nil
			}
			connection = nc
		} else {
			nc, err := nats_controller.Connect(fmt.Sprintf("nats://localhost:%d", config.GetInt("nats_port")))
			if err != nil {
				logger.GetNats().Error("Error in connection", map[string]interface{}{
					"error": err,
				})
				return nil
			}
			connection = nc
		}
	}
	return connection
}

// Publish to NATS server topic.
func Publish(topic string, data map[string]interface{}) {
	nc := GetConnection()
	if nc == nil {
		return
	}

	result, err := json_codec.JsonEncode(data)
	if err != nil {
		logger.GetNats().Error("Error parse data to json in PUBLISH", map[string]interface{}{
			"data":  data,
			"topic": topic,
			"error": err,
		})
	}

	commonErr := nc.Publish(topic, []byte(result))
	if commonErr != nil {
		logger.GetNats().Error("Error send PUBLISH", map[string]interface{}{
			"error": err,
			"topic": topic,
			"data":  data,
		})
	}
	logger.GetNats().Info("Success send PUBLISH", map[string]interface{}{
		"topic": topic,
		"data":  data,
	})
}
