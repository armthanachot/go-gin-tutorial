package model

import "time"

type Employee struct {
	Id        uint   `gorm:"primary_key" json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Position  string `json:"position"`
	CreatedAt  time.Time `json:"created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt  time.Time `json:"updated_at" sql:"DEFAULT:current_timestamp ON UPDATE current_timestamp"`
}