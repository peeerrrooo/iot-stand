// Package implement MongoDB schemas.
package schema

import "time"

// Schema for FEEDBACK Collection.
type Telemetry struct {
	Battery     int64     `"json:"Battery" bson:"Battery"`
	TotalRange  int64     `"json:"TotalRange" bson:"TotalRange"`
	Temperature int64     `"json:"Temperature" bson:"Temperature"`
	Mileage     int64     `"json:"Mileage" bson:"Mileage"`
	Vin         string    `"json:"Vin" bson:"Vin"`
	Created     time.Time `"json:"Created" bson:"Created"`
}
