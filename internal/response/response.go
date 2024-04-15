package response

import (
	"go-admin-template/internal/translator"
	"go-admin-template/pkg/util"

	"github.com/gin-gonic/gin"
)

// UnifiedResponse 统一返回
type UnifiedResponse struct {
	Code      int    `json:"code"`
	Data      any    `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

type Option func(*UnifiedResponse)

// HandleResponse 统一返回处理
func HandleResponse(c *gin.Context, data any, err error) {
	requestId, _ := c.Get(util.TrafficKey)
	if err != nil {
		c.JSON(200, UnifiedResponse{
			Code:      500,
			Data:      nil,
			Message:   translator.Translate(err),
			RequestId: requestId.(string),
		})
		return
	}

	c.JSON(200, UnifiedResponse{
		Code:      0,
		Data:      data,
		Message:   "成功",
		RequestId: requestId.(string),
	})
}

// HandleAbortResponse 统一 Abort 返回处理
func HandleAbortResponse(c *gin.Context, err string, opt ...Option) {
	requestId, _ := c.Get(util.TrafficKey)
	resp := UnifiedResponse{
		Code:      500,
		Data:      nil,
		Message:   err,
		RequestId: requestId.(string),
	}
	for _, fn := range opt {
		fn(&resp)
	}
	c.AbortWithStatusJSON(200, resp)
}

func WithCode(code int) Option {
	return func(r *UnifiedResponse) {
		r.Code = code
	}
}

func WithData(data any) Option {
	return func(r *UnifiedResponse) {
		r.Data = data
	}
}
