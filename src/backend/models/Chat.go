package models

import "time"

type Chat struct {
	Id   string    `json:"id"`
	From string    `json:"from"`
	Chat string    `json:"chat"`
	Time time.Time `json:"time"`
}

type ListChats struct {
	ListChat []Chat
}
