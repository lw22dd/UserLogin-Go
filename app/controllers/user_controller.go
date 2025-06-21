package controllers

import (
	"UserLogin/app/Model"
	"UserLogin/app/Service"
	"UserLogin/app/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		util.Response(c, http.StatusBadRequest, nil, "参数错误")
		return
	}

	user, err := Service.GetUserByEmail(credentials.Email)
	if err != nil {
		util.Response(c, http.StatusUnauthorized, nil, "邮箱或密码错误")
		return
	}

	if user.Password != credentials.Password {
		util.Response(c, http.StatusUnauthorized, nil, "密码错误")
		return
	}

	util.Response(c, http.StatusOK, user, "登录成功")
}

func RegisterHandler(c *gin.Context) {
	var newUser struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&newUser); err != nil {
		util.Response(c, http.StatusBadRequest, nil, "参数错误")
		return
	}

	err := Service.CreateUser(Model.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password, // 实际建议加密存储
	})
	if err != nil {
		util.Response(c, http.StatusInternalServerError, nil, "注册失败")
		return
	}

	util.Response(c, http.StatusCreated, nil, "注册成功")
}

func LogoutHandler(c *gin.Context) {
	// TODO: 清除 token 或会话
	util.Response(c, http.StatusOK, nil, "已退出")
}
func Test(c *gin.Context) {
	str := c.Query("str") // 获取查询参数 "str"
	if str == "" {
		util.Response(c, http.StatusBadRequest, nil, "缺少字符串参数")
		return
	}
	fmt.Println("接收到的字符串参数为:", str)
	util.Response(c, http.StatusOK, gin.H{"input": str}, "测试成功")
}
