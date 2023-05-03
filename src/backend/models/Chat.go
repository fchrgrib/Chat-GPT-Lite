package models

type Chat struct {
	IdHistoryChat string      `json:"idHistoryChat" gorm:"size:255"`
	IdChat        string      `json:"id" gorm:"primaryKey"`
	From          string      `json:"from"`
	Chat          string      `json:"chat"`
	Type          string      `json:"type"`
	Time          string      `json:"time"`
	ChatHistory   ChatHistory `gorm:"foreignKey:id_history_chat"`
}

type ChatHistory struct {
	Id       string `json:"id" gorm:"primaryKey"`
	LastChat string `json:"lastChat"`
	UpdateAt string `json:"updateAt"`
	Chat     []Chat `gorm:"foreignKey:id_history_chat;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
}
