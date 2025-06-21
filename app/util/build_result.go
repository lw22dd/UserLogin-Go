package util

import (
	"UserLogin/app/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 统一通过函数返回给前端
func Response(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, Model.Result{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
