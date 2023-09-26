package api

import (
	"encoding/json"
)

type ApiMessage struct {
	Error   int32
	Message string
	Data    string
}

func DeserializeMessage(message ApiMessage) string {
	b, err := json.Marshal(message)
	if err != nil {
		return "{\"Error\": -101, \"Message\": \"cannot deserialize message\", \"Data\": \"\"}"
	}
	return string(b)
}
