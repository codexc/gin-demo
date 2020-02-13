//这里放的是框架无关的纯业务处理代码
//为了统一约定涉及到的对象一律使用指针
package services

import (
	"learn/topic/src/AppInit"
	"learn/topic/src/Models"
)

type TopicService struct {
}

func (t *TopicService) GetList(req *TopicListRequest) *Models.TopicList {
	topicList := &Models.TopicList{}
	AppInit.GetDB().Limit(req.Size).Find(topicList)
	return topicList
}

func (t *TopicService) GetDetail(req *TopicDetailRequest) *Models.Topic {
	id := req.Id
	topic := &Models.Topic{}
	AppInit.GetDB().Find(topic, id)
	return topic
}
