package models

import "time"

type Chat struct {
	id   string    `json:"id"`
	from string    `json:"from"`
	chat string    `json:"chat"`
	time time.Time `json:"time"`
}

type ListChat struct {
	listChat *Chat
}
