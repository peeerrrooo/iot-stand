package mqtt

import (
	"fmt"
	"iot-stand/libs/config"
	"time"
	"iot-stand/libs/logger"
	mqtt_protocol "github.com/eclipse/paho.mqtt.golang"
)

/**
 * Get topic for PUBLISH.
 */
func getPublishTopic() string {
	return config.GetString("mqtt_pub_topic")
}

/**
 * Get topic for SUBSCRIBE.
 */
func getSubscribeTopic() string {
	return config.GetString("mqtt_sub_topic")
}

/**
 * Handler for MQTT subscribe.- Убрать подписка/ученики для админа
- А добавление ученика работает? Ошибка запроса на сервере...
 */

func getHandler(args ...interface{}) mqtt_protocol.MessageHandler {
	var handler mqtt_protocol.MessageHandler = func(client mqtt_protocol.Client, msg mqtt_protocol.Message) {
		message := fmt.Sprintf("%s", msg.Payload())

		logger.GetMQTT().Info("Get message", map[string]interface{}{
			"topic":   fmt.Sprintf("%s", msg.Topic()),
			"message": message,
		})

		if len(args) > 0 {
			mapApi, okConvert := args[0].(MapApi)
			if okConvert {
				callApi(mapApi, message)
			} else {
				logger.GetMQTT().Error("Cannot convert message to MapApi")
			}
		}
	}
	return handler
}

/**
 * Get MQTT connection.
 */
func getConnection(handler mqtt_protocol.MessageHandler) (mqtt_protocol.Client, error) {
	host := config.GetString("mqtt_host")
	port := config.GetInt("mqtt_port")

	opts := mqtt_protocol.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", host, port))
	opts.SetKeepAlive(5 * time.Second)
	opts.SetPingTimeout(2 * time.Second)
	opts.SetConnectTimeout(2 * time.Second)
	opts.SetMaxReconnectInterval(8 * time.Second)
	opts.SetDefaultPublishHandler(handler)

	logger.GetMQTT().Info("Prepare data for MQTT", map[string]interface{}{
		"host": host,
		"port": port,
	})

	c := mqtt_protocol.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		err := token.Error()
		logger.GetMQTT().Error("Error in get connection", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}
	return c, nil
}
