package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommonResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    code,
		"message": message,
	})
}

// Response 是一个通用的响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 返回一个成功的响应
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}
