package utils

import (
	db2 "backend/db"
	"backend/models"
)

func GetChats(id string) (models.ListChats, error) {
	var chats []models.Chat

	db, err := db2.GetDatabase()

	if err != nil {
		return models.ListChats{}, err
	}

	if err := db.Where("id = ?", id).Find(&chats); err.Error != nil {
		return models.ListChats{}, err.Error
	}

	allChats := models.ListChats{
		ListChat: chats,
	}

	return allChats, nil
}
