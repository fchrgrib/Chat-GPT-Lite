package models

import "time"

type QuestAns struct {
	Id         string    `json:"id" gorm:"primaryKey"`
	Question   string    `json:"question"`
	Answer     string    `json:"answer"`
	TimeCreate time.Time `json:"timeCreate"`
}
