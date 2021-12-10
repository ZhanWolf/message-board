package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/api"
)

func Messageroute(r *gin.Engine) {
	ms := r.Group("/message", cookie)
	{
		ms.POST("/sendmsg", api.Messagepost)
		ms.POST("/sendcom", api.Commentpost)
		ms.GET("/listmsg", api.Listmsg)
		ms.POST("/listcom", api.Liscon)
		ms.POST("/delete", api.Deletemsg)
		ms.POST("/update", api.Updatemsg)
		ms.GET("/mymsg", api.Listmymsg)
		ms.POST("/nonamemsg", api.Nonamemsg)
		ms.POST("/nonamecom", api.Nonamecom)
		ms.POST("/likes", api.Likes)
	}
}

func cookie(c *gin.Context) {
	ck, err := c.Cookie("now_user_login")
	if err != nil {
		fmt.Println(err)
		c.JSON(403, "未登录")
		c.Abort()
	} else {
		c.Set("cookie", ck)
		c.Next()
	}
}
