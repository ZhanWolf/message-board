package cmd

import (
	"github.com/gin-gonic/gin"
	"message-board/api"
)

func Userroute(r *gin.Engine) {
	us := r.Group("/user")
	{
		us.POST("/login", api.Login)
		us.POST("/Singup", api.Singup)
		us.POST("/Reset", api.Reset)
		us.GET("/clock", cookie, api.Clock)
	}
}
