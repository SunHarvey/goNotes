package models

import (
	"time"
)

type Answer struct {
	Id               int    `gorm:"primary_key" json:"id"`
	Body             string `sql:"type:text;" json:"body"`
	CreateAt         time.Time
	UpdateAt         time.Time
	Views            int
	Liks             int
	DisLike          int
	IsAcceptedAnswer bool
	UserID           int `gorm:"size:10"`
	QuestionID       int `gorm:"size:10"`
	Question         Question
	Qeestion_id      int `sql:"type:integer REFERENCES questions(id)"`
	User_id          int `sql:"type:integer REFERENCES users(id)"`
}
