package methods

import (
	"iot-stand/libs/json_codec"
	"iot-stand/server/services/ws/storage"
)

func GetWsMessage(data map[string]interface{}) {
	topic, commonErr := json_codec.GetString("topic", data)
	if commonErr != nil {
		return
	}
	receiveData, commonErr := json_codec.GetJson("data", data)
	if commonErr != nil {
		return
	}
	excludeOwners, isExclude := json_codec.GetArrayString("exclude_owners", data)

	for _, item := range storage.GetStorage().Storage {
		for _, t := range item.Topics {
			if t == topic {
				if isExclude == nil {
					isExcludeOwner := false
					for _, uuid := range excludeOwners {
						if item.UUID == uuid {
							isExcludeOwner = true
							break
						}
					}
					if !isExcludeOwner {
						item.Conn.WriteJSON(receiveData)
					}
					break
				} else {
					item.Conn.WriteJSON(receiveData)
					break
				}
			}
		}
	}
}
