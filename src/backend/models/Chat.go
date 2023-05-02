package models

type Chat struct {
	Id   string `json:"id" gorm:"primaryKey"`
	From string `json:"from"`
	Chat string `json:"chat"`
	Time string `json:"time"`
}

type ChatHistory struct {
	Id       string `json:"id" gorm:"primaryKey"`
	LastChat string `json:"lastChat"`
	UpdateAt string `json:"updateAt"`
}

type ListChatHistory struct {
	ListHistory []ChatHistory `json:"listHistory"`
}

type ListChats struct {
	Id       string `json:"id" gorm:"primaryKey"'`
	ListChat []Chat
}
