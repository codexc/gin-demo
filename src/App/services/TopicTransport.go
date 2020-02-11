//这里放请求参数处理 和 响应的处理
//将参数或者响应绑定到endpoint
//此文件和框架深度耦合
package services

import (
	"github.com/gin-gonic/gin"
	"learn/topic/src/App"
)

func SetTopicListRequest() App.RequestHandler {
	return func(ginCtx *gin.Context) (request interface{}, err error) {
		req := &TopicRequest{}
		err = ginCtx.BindQuery(req)
		if err != nil {
			return nil, err
		}
		return req, nil
	}
}

func SetTopicListResponse() App.ResponseHandler {
	return func(ginCtx *gin.Context, response interface{}) (err error) {
		//rsp := &TopicResponse{} //这里写错了 要用断言
		rsp := response.(*TopicResponse)
		ginCtx.JSON(200, rsp)
		return nil
	}
}
