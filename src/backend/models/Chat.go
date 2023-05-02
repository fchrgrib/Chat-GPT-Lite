package models

type Chat struct {
	IdHistoryChat string `json:"idHistoryChat"`
	IdChat        string `json:"id_chat" gorm:"primaryKey"`
	From          string `json:"from"`
	Chat          string `json:"chat"`
	Time          string `json:"time"`
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
