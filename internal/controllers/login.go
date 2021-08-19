package controllers

import (
	"message-processor/internal/models"
	"message-processor/internal/utils"
)

func CreateLogin(user models.User) {
	var data []byte
	var uuid string

	data = utils.StructToByte(user)
	uuid = utils.GenerateUUID5(data)

	user.Id = uuid
}
