package ginresp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewSuccess(c *gin.Context, data interface{}) {
	newResp(c, true, Success, mapper[Success], data)
}

func NewFailure(c *gin.Context, code AppCode, data interface{}) {
	err, ok := data.(error)
	if ok {
		data = err.Error()
	}
	newResp(c, false, code, mapper[code], data)
}

// newResp 创建一个新的 Response 结构体
func newResp(c *gin.Context, success bool, code AppCode, message string, data interface{}) {
	c.JSON(http.StatusOK, AppResponse{
		Success: success,
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// AppResponse 消息返回类，返回指定的数据结构。
type AppResponse struct {
	Success bool        `json:"success"`
	Code    AppCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewAppResponse(success bool, code AppCode, message string, data interface{}) *AppResponse {
	return &AppResponse{Success: success, Code: code, Message: message, Data: data}
}
