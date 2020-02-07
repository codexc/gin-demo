package src

type GetByPageQuery struct {
    UserName string `json:"username" form:"username" binding:"required"`
    Page int `json:"p" form:"page"`
    Size int `json:"s" form:"size"`
}

type TopicEntity struct {
    Title string `json:"title" binding:"min=4,max=10"`
    SubTitle string `json:"subtitle" binding:"required,nefield=Title"`
    UserIp string `json:"ip" binding:"ip"`
    Score int `json:"score" binding:"omitempty,gt=5"`
}

type MultiTopicEntity struct {
    MultiTopics []TopicEntity `json:"topics" binding:"gt=0,lte=2,dive"`
}
