package routes

import "github.com/gin-gonic/gin"

func RegisterAPIRoutes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/login", controllers.LoginHandler)
		user.POST("/register", controllers.RegisterHandler)
		user.POST("/logout", controllers.LogoutHandler)
	}
}
