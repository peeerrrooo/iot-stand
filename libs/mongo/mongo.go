// Package implement connections and logic with collection for MongoDB.
package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"iot-stand/libs/config"
	"iot-stand/libs/logger"
)

var mongoConnection *MongoConnection

// Type that encapsulation all logic with MongoDB connections and CRUD operations.
type MongoConnection struct {
	OriginalSession     *mgo.Session
	DB                  string
	TelemetryCollection string
}

// Create one original connection session with MongoDB.
func (self *MongoConnection) createOriginalSession() error {
	if self.OriginalSession == nil {
		var host string
		if config.GetBool("docker") {
			host = fmt.Sprintf("%s:%s@mongo:%d",
				config.GetString("mongo_user"),
				config.GetString("mongo_pass"),
				config.GetInt("mongo_port"))
		} else {
			host = fmt.Sprintf("localhost:27017")
		}

		session, err := mgo.Dial(host)
		if err != nil {
			logger.GetMongo().Error("Error get session", map[string]interface{}{
				"error": err,
			})
			return err
		}

		self.OriginalSession = session
		self.DB = "iot-stand"
		self.TelemetryCollection = "telemetry"

		logger.GetMongo().Info("Success connect")
		return nil
	}
	return nil
}

// Get MongoDB encapsulation object.
func GetMongoConnection() *MongoConnection {
	if mongoConnection == nil {
		mongoConnection = new(MongoConnection)
		mongoConnection.createOriginalSession()
	}
	return mongoConnection
}
