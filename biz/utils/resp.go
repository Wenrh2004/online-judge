package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"online-judge/biz/utils/errno"
)

// Resp ==> the struct of unified response
type Resp[T any] struct {
	Code    int32  `json:"status_code"`
	Message string `json:"status_msg"`
	Data    T      `json:"data"`
}

// Response ==> unified response
func Response(ctx context.Context, c *app.RequestContext, err error, data ...interface{}) {
	resp := errno.BuildBaseResp(err)

	if resp.StatusCode != 0 {
		c.JSON(http.StatusBadRequest, Resp[interface{}]{
			Code:    resp.StatusCode,
			Message: resp.StatusMsg,
		})
		return
	}

	c.JSON(http.StatusOK, Resp[interface{}]{
		Code:    resp.StatusCode,
		Message: resp.StatusMsg,
		Data:    data,
	})
}
