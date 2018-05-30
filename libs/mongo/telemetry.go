package mongo

import (
	"iot-stand/libs/logger"
	"iot-stand/libs/mongo/schema"
	"gopkg.in/mgo.v2/bson"
)

// Add TELEMETRY.
func (self *MongoConnection) AddTelemetry(el schema.Telemetry) error {
	session := self.OriginalSession.Copy()
	defer session.Close()

	err := session.DB(self.DB).C(self.TelemetryCollection).Insert(el)
	if err != nil {
		logger.GetMongo().Error("Error ADD TELEMETRY", map[string]interface{}{
			"telemetry": el,
			"error":     err,
		})
		return err
	}

	logger.GetMongo().Info("Success ADD TELEMETRY", map[string]interface{}{
		"telemetry": el,
	})
	return nil
}

// Get TELEMETRY.
func (self *MongoConnection) GetTelemetry() ([]schema.Telemetry, error) {
	session := self.OriginalSession.Copy()
	defer session.Close()

	var result []schema.Telemetry
	err := session.DB(self.DB).C(self.TelemetryCollection).Find(bson.M{}).All(&result)
	if err != nil {
		logger.GetMongo().Error("Error GET TELEMETRY", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return result, nil
}

// Purge TELEMETRY.
func (self *MongoConnection) PurgeTelemetry() (error) {
	session := self.OriginalSession.Copy()
	defer session.Close()

	err := session.DB(self.DB).C(self.TelemetryCollection).DropCollection()
	if err != nil {
		logger.GetMongo().Error("Error PURGE TELEMETRY", map[string]interface{}{
			"error": err,
		})
		return err
	}

	return nil
}
