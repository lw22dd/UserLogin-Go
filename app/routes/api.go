package routes

import (
	"UserLogin/app/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/login", controllers.LoginHandler)
		user.POST("/register", controllers.RegisterHandler)
		user.POST("/logout", controllers.LogoutHandler)
		user.GET("/test", controllers.Test)

	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}
