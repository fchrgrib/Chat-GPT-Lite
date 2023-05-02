package models

type QuestAns struct {
	Id         string `json:"id" gorm:"primaryKey"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	TimeCreate string `json:"timeCreate"`
}
