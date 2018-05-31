package mqtt

import (
	"iot-stand/libs/logger"
	"time"
)

/**
 * Send MQTT message.
 */
func Subscribe(reconnectInterval int, mapApi MapApi) error {
	logger.GetMQTT().Info("Success subscribe")

	for {
		c, err := getConnection(getHandler(mapApi))

		if err == nil {
			if token := c.Subscribe(getSubscribeTopic(), 0, nil); token.Wait() && token.Error() != nil {
				err := token.Error()
				logger.GetMQTT().Error("Error in SUBSCRIBE", map[string]interface{}{
					"error": err,
				})
				c.Disconnect(250)
				return err
			}

			logger.GetMQTT().Info("Successful SUBSCRIBE")
			return nil
		}

		time.Sleep(time.Duration(reconnectInterval) * time.Second)
	}
}
