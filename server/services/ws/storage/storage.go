// Package that encapsulate storage for WS connections.
package storage

import (
	"sync"
	"github.com/gorilla/websocket"
	"iot-stand/libs/logger"
)

var sg *Storage = nil

// Type that encapsulate logic for CLIENT connections storage.
type Storage struct {
	sync.Mutex
	Storage []*StorageItem
}

// Type that encapsulate logic for CLIENT.
type StorageItem struct {
	sync.Mutex
	UUID   string
	Conn   *websocket.Conn
	Topics []string
}

// Add CLIENT to STORAGE.
func (self *Storage) addToStorage(UUID string, conn *websocket.Conn) *StorageItem {
	item := new(StorageItem)
	item.UUID = UUID
	item.Conn = conn
	item.Topics = []string{}
	self.Storage = append(self.Storage, item)

	findIndex := -1
	for i, v := range self.Storage {
		if v.UUID == UUID {
			findIndex = i
			break
		}
	}

	if findIndex == -1 {
		self.Storage = append(self.Storage, item)
	}
	return item
}

// Add CLIENT to STORAGE.
func (self *Storage) AddToStorage(UUID string, conn *websocket.Conn) *StorageItem {
	self.Lock()
	defer self.Unlock()
	return self.addToStorage(UUID, conn)
}

// Remove CLIENT to STORAGE.
func (self *Storage) removeFromStorage(UUID string) {
	if len(self.Storage) > 0 {
		findIndex := -1
		for i, v := range self.Storage {
			if v.UUID == UUID {
				findIndex = i
				break
			}
		}
		self.Storage = append(self.Storage[:findIndex], self.Storage[findIndex+1:]...)
	}
}

// Remove CLIENT to STORAGE.
func (self *Storage) RemoveFromStorage(UUID string) {
	self.Lock()
	defer self.Unlock()
	self.removeFromStorage(UUID)
}

// Add TOPIC for CLIENT to STORAGE.
func (self *Storage) addTopic(UUID string, topic string) {
	findIndex := -1
	for i, v := range self.Storage {
		if v.UUID == UUID {
			findIndex = i
			break
		}
	}
	findIndexTopic := -1
	for i, v := range self.Storage[findIndex].Topics {
		if v == topic {
			findIndexTopic = i
			break
		}
	}
	if findIndexTopic == -1 {
		self.Storage[findIndex].Topics = append(self.Storage[findIndex].Topics, topic)
		logger.GetWS().Info("Success subscribe client to topic", map[string]interface{}{
			"uuid":  UUID,
			"topic": topic,
		})
	}
}

// Add TOPIC for CLIENT to STORAGE.
func (self *Storage) AddTopic(UUID string, topic string) {
	self.Lock()
	defer self.Unlock()
	self.addTopic(UUID, topic)
}

// Write JSON to WS connection.
func (self *StorageItem) writeJSON(value map[string]interface{}) error {
	return self.Conn.WriteJSON(value)
}

// Write JSON to WS connection.
func (self *StorageItem) WriteJSON(value map[string]interface{}) error {
	self.Lock()
	defer self.Unlock()
	return self.writeJSON(value)
}

func GetStorage() *Storage {
	if sg == nil {
		sg = new(Storage)
		sg.Storage = []*StorageItem{}
	}
	return sg
}
