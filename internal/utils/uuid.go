package utils

import (
	"github.com/google/uuid"
)

func GenerateUUID5(data []byte) string {
	return uuid.NewSHA1(uuid.New(), data).String()
}
