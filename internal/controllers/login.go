package controllers

import (
	"encoding/json"
	"log"
	"message-processor/internal/models"
	"message-processor/internal/service/aws"
	"message-processor/internal/utils"
)

func CreateLogin(user models.User) {
	var data []byte
	var uuid string

	data, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error on convert objetc to []byte %v", err)
	}
	uuid = utils.GenerateUUID5(data)

	user.Id = uuid

	aws.SendMessage(user)
}
