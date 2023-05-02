package utils

import (
	"backend/algorithm"
	"backend/db"
	"backend/models"
	"errors"
	"github.com/google/uuid"
	"time"
)

func InsertQuestAns(input string) error {
	_db, err := db.GetDatabase()

	if err != nil {
		return err
	}

	idxQues, _ := algorithm.BM("Tambahkan pertanyaan ", input)
	idxAns, _ := algorithm.BM(" dengan jawaban ", input)

	var questChars, ansChars []byte

	for i := idxQues + len("Tambahkan pertanyaan "); i < idxAns+1; i++ {
		questChars = append(questChars, input[i])
	}
	for i := idxAns + len(" dengan jawaban "); i < len(input); i++ {
		ansChars = append(ansChars, input[i])
	}

	if err := _db.Where("question = ?", questChars); err != nil {
		return errors.New("pertanyaan sudah ada")
	}

	if err := _db.Create(models.QuestAns{
		Id:         uuid.New().String(),
		Question:   string(questChars),
		Answer:     string(ansChars),
		TimeCreate: time.Now().Local(),
	}); err.Error != nil {
		return err.Error
	}

	return nil
}
