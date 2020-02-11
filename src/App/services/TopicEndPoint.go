//这里放request 和 response的定义
//可以通过proto生成出来的
package services

import (
	"context"
	"learn/topic/src/App"
	"learn/topic/src/Models"
)

type TopicRequest struct {
	Size int `form:"size"`
}

type TopicResponse struct {
	Result *Models.TopicList `json:"result"`
}

func TopicEndPoint(topic *TopicService) App.EndPoint {
	return func(context context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*TopicRequest)
		return &TopicResponse{Result: topic.GetList(req)}, nil
	}
}
