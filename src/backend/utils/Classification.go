package utils

import (
	"backend/db"
	"backend/models"
	"regexp"
)

func CheckDate(input string) (bool, string) {

	pattern := `\d{2}/\d{2}/\d{4}`
	re := regexp.MustCompile(pattern)
	matches := re.FindString(input)

	if len(matches) == 0 {
		return false, ""
	}

	return true, matches
}

func CheckCalculate(input string) (bool, string) {
	pattern := "Hitunglah (.+)"

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(input)

	if len(matches) == 0 {
		return false, ""
	}

	mathPattern := `^[\d*/+\-^()]*$`

	rex := regexp.MustCompile(mathPattern)
	matchesx := rex.MatchString(matches[1])

	if matchesx {
		return matchesx, matches[1]
	}

	return false, ""
}

func CheckInsertQuesAns(input string) (isInsert, isQuesSame bool, matches1, matches2 string) {
	pattern := "Tambahkan pertanyaan (.+) dengan jawaban (.+)"

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(input)

	if len(matches) == 0 {
		return false, false, "", ""
	}

	isInsert, matches1, matches2 = true, matches[1], matches[2]

	_db, err := db.GetDatabase()
	if err != nil {
		return
	}

	var question models.QuestAns

	if _ = _db.Where("question = ?", matches1).First(&question); len(question.Question) != 0 {
		isQuesSame = true
		return
	}
	isQuesSame = false

	return
}

func CheckEraseQues(input string) (isAsk, isFound bool, quest string) {
	pattern := "Hapus pertanyaan (.+)"

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 0 {
		isAsk = true
	} else {
		isAsk = false
		return
	}

	_db, err := db.GetDatabase()
	if err != nil {
		return
	}

	var question models.QuestAns

	if _ = _db.Where("question = ?", matches[1]).First(&question); len(question.Question) != 0 {
		_ = _db.Delete(question)
		isFound = true
		quest = question.Question
		return
	}
	quest = question.Question
	isFound = false
	return
}
