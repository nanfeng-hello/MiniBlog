package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}, msg ...string) {
	default_msg := "success"
	if len(msg) > 0 {
		default_msg = msg[0]
	}

	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  default_msg,
		Data: data,
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
	})
}
