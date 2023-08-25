package models

import "time"

type Message struct {
	Id        string
	Text      string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
