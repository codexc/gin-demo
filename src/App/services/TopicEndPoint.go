//这里放request 和 response的定义
//可以通过proto生成出来的
package services

import (
	"context"
	"learn/topic/src/App"
	"learn/topic/src/Models"
)

type TopicListRequest struct {
	Size int `form:"size"`
}

type TopicListResponse struct {
	Result *Models.TopicList `json:"result"`
}

func TopicListEndPoint(topic *TopicService) App.EndPoint {
	return func(context context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*TopicListRequest)
		ret, err := topic.GetList(req)
		return &TopicListResponse{Result: ret}, err
	}
}

type TopicDetailRequest struct {
	Id int `form:"id"`
}
type TopicDetailResponse struct {
	Detail *Models.Topic
}

func TopicDetailEndPoint(topic *TopicService) App.EndPoint {
	return func(context context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*TopicDetailRequest)
		ret, err := topic.GetDetail(req)
		return &TopicDetailResponse{Detail: ret}, err
	}
}
