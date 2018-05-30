package mqtt

import (
	"time"
	"iot-stand/libs/logger"
	"iot-stand/libs/json_codec"
)

/**
 * Send MQTT message.
 */
func Publish(method string, param map[string]interface{}) error {
	c, err := getConnection(getHandler())
	if err != nil {
		return err
	}

	message := ""
	if param != nil {
		message, _ = json_codec.JsonEncode(map[string]interface{}{
			"method": method,
			"param":  param,
		})
	} else {
		message, _ = json_codec.JsonEncode(map[string]interface{}{
			"method": method	,
		})
	}

	token := c.Publish(getPublishTopic(), 0, false, message)
	token.Wait()
	time.Sleep(3 * time.Second)
	c.Disconnect(250)

	logger.GetMQTT().Info("Successful publish method", map[string]interface{}{
		"method": method,
		"param":  param,
	})
	return nil
}
