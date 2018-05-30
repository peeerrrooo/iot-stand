package mqtt

import (
	"iot-stand/libs/json_codec"
	"fmt"
	"iot-stand/libs/logger"
)

/**
 * Type of map for mapping function call of API MQTT.
 * Structure, example:
 * map[string]map[string]interface{}{
 * 		"SetAny": map[string]interface{}{
 *			"method": SetAny,
 * 			"param": true,
 *		},
 *		"SetSome": map[string]interface{}{
 *			"method": SetSome,
 *		},
 * }
 */
type MapApi map[string]interface{}

/**
 * Parse input subscribe message and mapping to MapApi.
 */
func callApi(mapApi MapApi, message string) {
	param, err := json_codec.JsonParse(message)
	if err != nil {
		logger.GetMQTT().Error(fmt.Sprintf("Incorrect input JSON"))
		return
	}
	_, okMethod := param["method"]
	if okMethod {
		method, okMethodField := param["method"].(string)
		if okMethodField {
			_, okMethodApi := mapApi[method]
			if okMethodApi {
				resultParam, okJsonParam := param["param"]
				if okJsonParam {
					go mapApi[method].(func(interface{}))(resultParam)
					logger.GetMQTT().Info("Success call method", map[string]interface{}{
						"method": method,
					})
					return
				} else {
					go mapApi[method].(func(interface{}))(nil)
					logger.GetMQTT().Info("Success call method", map[string]interface{}{
						"method": method,
					})
					return
				}
				return
			}
			logger.GetMQTT().Error(fmt.Sprintf("Unknown API method"))
			return
		}
		logger.GetMQTT().Error(fmt.Sprintf("Field 'method' from input JSON must be string"))
		return
	}
	logger.GetMQTT().Error(fmt.Sprintf("Miss 'method' filed from input JSON"))
	return
}
