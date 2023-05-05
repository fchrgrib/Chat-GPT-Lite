package utils

import (
	"backend/db"
	"backend/models"
	"github.com/google/uuid"
)

func InsertQuestionAnswer(question, answer string) error {
	_db, err := db.GetDatabase()

	if err != nil {
		return err
	}

	questAns := models.QuestAns{
		Id:         uuid.New().String(),
		Question:   question,
		Answer:     answer,
		TimeCreate: GetJktTimeZone(),
	}

	if err := _db.Create(questAns); err.Error != nil {
		return err.Error
	}
	return nil
}

func UpdateAnswer(question, answer string) error {
	var questAns models.QuestAns
	_db, err := db.GetDatabase()
	if err != nil {
		return err
	}

	if err := _db.Where("question = ?", question).First(&questAns); err.Error != nil {
		return err.Error
	}

	questAns.Answer = answer
	questAns.TimeCreate = GetJktTimeZone()

	_db.Save(questAns)

	return nil
}
