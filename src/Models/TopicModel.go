package Models

import "time"

type Topic struct {
	Id       int       `json:"id" gorm:"PRIMARY_KEY"`
	Title    string    `json:"title" binding:"min=4,max=10"`
	Content  string    `json:"content" binding:"required,nefield=Title"`
	Username string    `json:"username" binding:"gt=2" gorm:"Column:userName"`
	Category string    `json:"category" gorm:"Column:topic_category"`
	Ctime    time.Time `json:"ctime" `
	Mtime    time.Time `json:"mtime"`
}

type TopicList []*Topic
