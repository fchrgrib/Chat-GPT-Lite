package utils

import (
	db2 "backend/db"
	"backend/models"
)

func GetChats(id string) ([]models.Chat, error) {
	var chats []models.Chat

	db, err := db2.GetDatabase()

	if err != nil {
		return nil, err
	}

	if err := db.Table("chats").Where("id_history_chat = ?", id).Order("time").Find(&chats); err.Error != nil {
		return nil, err.Error
	}

	return chats, nil
}

func GetHistoryChats() ([]models.ChatHistory, error) {
	var chats []models.ChatHistory

	db, err := db2.GetDatabase()

	if err != nil {
		return nil, err
	}

	if err := db.Table("chat_histories").Find(&chats); err.Error != nil {
		return nil, err.Error
	}

	return chats, nil
}
