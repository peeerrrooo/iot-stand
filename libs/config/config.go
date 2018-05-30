// Package implement logic for config.{yaml|json}.
package config

import (
	"github.com/spf13/viper"
	"iot-stand/libs/logger"
)

var isInit bool = false

// Init dirs and paths to config(yaml/json) file.
func initConfig() error {
	viper.SetEnvPrefix("")

	// Common variables.
	viper.BindEnv("dev")
	viper.BindEnv("docker")
	viper.SetDefault("dev", true)
	viper.SetDefault("docker", false)

	// NATS.
	viper.BindEnv("nats_user")
	viper.BindEnv("nats_pass")
	viper.BindEnv("nats_port")
	viper.SetDefault("nats_user", "")
	viper.SetDefault("nats_pass", "")
	viper.SetDefault("nats_port", 4222)

	// MONGO.
	viper.BindEnv("mongo_user")
	viper.BindEnv("mongo_pass")
	viper.BindEnv("mongo_port")
	viper.SetDefault("mongo_user", "")
	viper.SetDefault("mongo_pass", "")
	viper.SetDefault("mongo_port", 27017)

	// MQTT.
	viper.BindEnv("mqtt_host")
	viper.BindEnv("mqtt_port")
	viper.BindEnv("mqtt_sub_topic")
	viper.BindEnv("mqtt_pub_topic")
	viper.SetDefault("mqtt_host", "localhost")
	viper.SetDefault("mqtt_port", 1883)
	viper.SetDefault("mqtt_sub_topic", "server")
	viper.SetDefault("mqtt_pub_topic", "device")


	// Game service.
	viper.BindEnv("ws_service_port")
	viper.SetDefault("ws_service_port", 9120)

	isInit = true
	logger.GetConfig().Info("Successful init")
	return nil
}

// Get INT value from config.
func GetInt(key string) int {
	if !isInit {
		initConfig()
	}
	value := viper.GetInt(key)
	logger.GetConfig().Info("Get INT value", map[string]interface{}{
		"type":  "INT",
		"key":   key,
		"value": value,
	})
	return value
}

// Get INT64 value from config.
func GetInt64(key string) int64 {
	if !isInit {
		initConfig()
	}
	value := viper.GetInt64(key)
	logger.GetConfig().Info("Get INT64 value", map[string]interface{}{
		"type":  "INT64",
		"key":   key,
		"value": value,
	})
	return value
}

// Get FLOAT value from config.
func GetFloat64(key string) float64 {
	if !isInit {
		initConfig()
	}
	value := viper.GetFloat64(key)
	logger.GetConfig().Info("Get FLOAT64 value", map[string]interface{}{
		"type":  "FLOAT64",
		"key":   key,
		"value": value,
	})
	return value
}

// Get BOOLEAN value from config.
func GetBool(key string) bool {
	if !isInit {
		initConfig()
	}
	value := viper.GetBool(key)
	logger.GetConfig().Info("Get BOOl value", map[string]interface{}{
		"type":  "BOOl",
		"key":   key,
		"value": value,
	})
	return value
}

// Get STRING value from config.
func GetString(key string) string {
	if !isInit {
		initConfig()
	}
	value := viper.GetString(key)
	logger.GetConfig().Info("Get STRING value", map[string]interface{}{
		"type":  "STRING",
		"key":   key,
		"value": value,
	})
	return value
}
