package utils

import (
	"bytes"
	"encoding/json"
	"message-processor/internal/models"
)

func StructToByte(user models.User) []byte {
	bodyBytes := new(bytes.Buffer)
	json.NewEncoder(bodyBytes).Encode(user)

	return bodyBytes.Bytes()
}
