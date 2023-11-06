package server

import (
	"github.com/gin-gonic/gin"
	"otp/go-gin-poc/controllers"
)

func initializeRouter(router *gin.Engine) {
	router.GET("/health", controllers.Health_check())
	user := router.Group("user")
	{
		v1 := user.Group("v1")
		{
			api := v1.Group("api")
			{
				api.POST("/user-login", controllers.SingleLoginRequired())
				api.POST("/send-login/otp", controllers.Health_check())
				api.POST("/submit-otp", controllers.Health_check())
			}
		}
	}
}
