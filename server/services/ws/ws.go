// Package implement WS server.
package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"iot-stand/libs/config"
	"iot-stand/libs/logger"
	"iot-stand/server/services/ws/storage"
	"iot-stand/server/services/ws/rpc"
	"iot-stand/server/services/ws/helpers"
	"iot-stand/libs/json_codec"
	"iot-stand/server/services/ws/topic"
	"github.com/satori/go.uuid"
	"iot-stand/server/services/ws/nats_api"
	"iot-stand/server/services/ws/websocket_nats_api"
)

// Create WS Server.
func createWsServer() {
	storage.GetStorage()
	rpc.GetRPC()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/ws", getWsApiRouter())

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.GetInt("ws_service_port")),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.GetWS().Info("Start WS Server")
	err := s.ListenAndServe()
	if err != nil {
		logger.GetWS().Error("Error in start WS server", map[string]interface{}{
			"error": err,
		})
	}
}

// Create Service.
func CreateService() {
	go nats_api.GetProvider()
	go websocket_nats_api.GetProvider()
	createWsServer()
}

// Get WS Api Router.
func getWsApiRouter() func(c *gin.Context) {
	wsUpgradeOptions := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return func(c *gin.Context) {
		func(resWriter http.ResponseWriter, request *http.Request) {

			conn, err := wsUpgradeOptions.Upgrade(resWriter, request, nil)
			defer conn.Close()
			if err != nil {
				logger.GetWS().Error("Error in get WS UPGRADE", map[string]interface{}{
					"error": err,
				})
				return
			}

			gUUID, err := uuid.NewV4()
			if err != nil {
				logger.GetWS().Error("Error generate UUID", map[string]interface{}{
					"error": err,
				})
				return
			}

			uuid := fmt.Sprintf("%s", gUUID)
			logger.GetWS().Info("Success connect new client", map[string]interface{}{
				"uuid": uuid,
			})
			readMessageProcess(uuid, conn, storage.GetStorage())
		}(c.Writer, c.Request)
	}
}

// Implement logic for INPUT/OUTPUT WS.
func readMessageProcess(uuid string, conn *websocket.Conn, sg *storage.Storage) {
	owner := sg.AddToStorage(uuid, conn)
	sg.AddTopic(uuid, topic.COMMON_TOPIC)
	defer sg.RemoveFromStorage(uuid)
	defer func(uuid string) {
		logger.GetWS().Info("Success disconnect client", map[string]interface{}{
			"uuid": uuid,
		})
	}(uuid)

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		if t == websocket.TextMessage {
			message := string(msg)
			logger.GetWS().Info("Get INPUT message", map[string]interface{}{
				"message": message,
				"type":    "TEXT",
			})

			fields, err := json_codec.JsonParse(message)
			if err != nil {
				helpers.SendError(nil, owner, 1000)
				continue
			}

			// Parse message for PRC protocol.
			method, commonErr := json_codec.GetString("method", fields)
			if commonErr != nil {
				helpers.SendError(nil, owner, 1001)
				continue
			}
			reqID, commonErr := json_codec.GetAny("req_id", fields)
			if commonErr != nil {
				helpers.SendError(nil, owner, 1002)
				continue
			}
			params, commonErr := json_codec.GetJson("params", fields)
			if commonErr != nil {
				params = map[string]interface{}{}
			}
			rpc.GetRPC().ReadProcess(method, reqID, owner, params)
		} else {
			logger.GetWS().Info("Get INPUT message", map[string]interface{}{
				"message": msg,
				"type":    "BINARY",
			})
		}
	}
}
