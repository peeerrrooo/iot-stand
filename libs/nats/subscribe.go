package nats

import (
	nats_controller "github.com/nats-io/go-nats"
	"iot-stand/libs/logger"
	"iot-stand/libs/json_codec"
	"fmt"
)

// Type that encapsulate logic for NATS subscription.
type Subscribe struct {
	Service string
	MethodsMap map[string]func(map[string]interface{})
}

// Run logic for subscribe for NATS server.
func (self *Subscribe) run() {
	nc := GetConnection()
	if nc == nil {
		return
	}
	logger.GetNats().Info("Success subscribe", map[string]interface{}{
		"service": self.Service,
	})

	ch := make(chan *nats_controller.Msg)
	_, err := nc.ChanSubscribe(self.Service, ch)
	if err != nil {
		logger.GetNats().Error("Error in subscribe", map[string]interface{}{
			"error":   err,
			"service": self.Service,
		})
		return
	}

	logger.GetNats().Info("Success connect", map[string]interface{}{
		"topic": self.Service,
	})
	for msg := range ch {
		logger.GetNats().Info("Success get message", map[string]interface{}{
			"msg":     string(msg.Data),
			"service": self.Service,
		})

		receive, err := json_codec.JsonParse(string(msg.Data))
		if err != nil {
			logger.GetNats().Error("Incorrect JSON format for input message", map[string]interface{}{
				"service": self.Service,
			})
			continue
		}
		methodField, methodErr := json_codec.GetString("method", receive)
		if methodErr != nil {
			logger.GetNats().Error("Miss field 'method'", map[string]interface{}{
				"service": self.Service,
			})
			continue
		}
		method, okMethod := self.MethodsMap[methodField]
		if !okMethod {
			logger.GetNats().Error("Unknown method", map[string]interface{}{
				"service": self.Service,
				"method":  methodField,
			})
			continue
		}

		data, dataErr := json_codec.GetJson("data", receive)
		if dataErr != nil {
			go method(nil)
			logger.GetNats().Info("Success call method", map[string]interface{}{
				"service": self.Service,
				"method":  methodField,
			})
		} else {
			go method(data)
			logger.GetNats().Info("Success call method", map[string]interface{}{
				"service": self.Service,
				"method":  methodField,
			})
		}
	}
}

// Init subscribe.
func InitSubscribe(service string, methodsMap map[string]func(map[string]interface{})) *Subscribe {
	subscribe := new(Subscribe)
	subscribe.Service = fmt.Sprintf("%s-service", service)
	subscribe.MethodsMap = methodsMap
	go subscribe.run()
	return subscribe
}
