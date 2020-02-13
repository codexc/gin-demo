//这里放的是框架无关的纯业务处理代码
//为了统一约定涉及到的对象一律使用指针
package services

import (
	"fmt"
	"learn/topic/src/AppInit"
	"learn/topic/src/Models"
)

type TopicService struct {
}

func (t *TopicService) GetList(req *TopicListRequest) (*Models.TopicList, error) {
	topicList := &Models.TopicList{}
	err := AppInit.GetDB().Limit(req.Size).Find(topicList).Error
	if err != nil {
		return nil, err
	}
	return topicList, nil
}

func (t *TopicService) GetDetail(req *TopicDetailRequest) (*Models.Topic, error) {
	id := req.Id
	topic := &Models.Topic{}
	if AppInit.GetDB().Find(topic, id).RowsAffected != 1 {
		return nil, fmt.Errorf("no topic found")
	}
	return topic, nil
}
