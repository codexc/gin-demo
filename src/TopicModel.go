package src

import "time"

type GetByPageQuery struct {
	UserName string `json:"username" form:"username" binding:"required"`
	Page     int    `json:"p" form:"page"`
	Size     int    `json:"s" form:"size"`
}

type Topic struct {
	Id       int       `json:"id" gorm:"PRIMARY_KEY"`
	Title    string    `json:"title" binding:"min=4,max=10"`
	Content  string    `json:"content" binding:"required,nefield=Title"`
	Username string    `json:"username" binding:"gt=5" gorm:"Column:userName"`
	Category string    `json:"category" gorm:"Column:topic_category"`
	Ctime    time.Time `json:"ctime" `
	Mtime    time.Time `json:"mtime"`
}

type MultiTopicEntity struct {
	MultiTopics []Topic `json:"topics" binding:"gt=0,lte=2,dive"`
}
