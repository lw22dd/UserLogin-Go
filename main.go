package main

import (
	"UserLogin/app/Service"
	"UserLogin/app/config"
	"UserLogin/app/routes"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//链接数据库
	config.DataBaseInit()

	//数据库测试
	fmt.Println("=============== 开始测试查询 ===============")
	users, err := Service.GetAllUsers()
	if err != nil {
		fmt.Println("查询 user 表失败:", err)
		return
	}

	fmt.Println("\nuser 表的数据：")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s , Password: %s\n", user.ID, user.Name, user.Email, user.Password)
	}
	fmt.Println("=============== 查询结束 ===============")

	//初始化路由
	r := gin.Default()
	// 注意跨域问题，后端没有返回 Access-Control-Allow-Origin 响应头，前端请求被浏览器拦截。
	r.Use(cors.Default())
	routes.RegisterAPIRoutes(r)

	if err := r.Run(":8080"); err != nil {
		panic("启动失败: " + err.Error())
	}

}
