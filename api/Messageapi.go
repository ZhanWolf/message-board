package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"message-board/service"
	"net/http"
	"strconv"
	"unicode/utf8"
)

func Messagepost(c *gin.Context) {
	touser := c.PostForm("touser")   //发送给谁
	message := c.PostForm("message") //评论的内容
	fromuser, _ := c.Cookie("now_user_login")

	err := service.Checkuseraliveser(touser)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, "没有该账户")
		return
	}
	if utf8.RuneCountInString(message) <= 5 {
		c.JSON(http.StatusOK, "评论长度应该大于5")
		return
	}

	err = service.Setmessage(touser, fromuser, message, fromuser)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, "发送失败")
	}
	c.JSON(http.StatusOK, "发送成功")
}

func Commentpost(c *gin.Context) {
	id := c.PostForm("id")           //母message的id
	message := c.PostForm("message") //评论的内容
	fromuser, _ := c.Cookie("now_user_login")
	id2, _ := strconv.Atoi(id)
	fmt.Println(id2)
	if dao.Querypid(id2) == false {
		c.JSON(http.StatusOK, "没有这条评论")
		return
	}
	if utf8.RuneCountInString(message) <= 5 {
		c.JSON(http.StatusOK, "评论长度应该大于5")
		return
	}

	err := service.Setcommnet(id2, fromuser, message, fromuser)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, "发送失败")
		return
	}
	c.JSON(http.StatusOK, "发送成功")
}

func Updatemsg(c *gin.Context) {
	id := c.PostForm("id")           //修改message的id
	message := c.PostForm("message") //修改后的内容
	username, _ := c.Cookie("now_user_login")
	id2, _ := strconv.Atoi(id)
	if dao.Checkueser(id2) != username && dao.Checktruename(id2) != username {
		c.JSON(http.StatusOK, "这不是您的评论")
		return
	}

	service.Updatemsg(message, id2)
	c.JSON(http.StatusOK, gin.H{
		"修改评论成功": "",
		"修改为":    message,
	})
}

func Deletemsg(c *gin.Context) {
	id := c.PostForm("id") //删除评论的id
	username, _ := c.Cookie("now_user_login")
	id2, _ := strconv.Atoi(id)
	if dao.Checkueser(id2) != username && dao.Checktruename(id2) != username {
		c.JSON(http.StatusOK, "这不是您的评论")
		return
	}
	service.Deletemsg(id2)
	c.JSON(http.StatusOK, "删除信息成功")
}

func Listmsg(c *gin.Context) {
	username, _ := c.Cookie("now_user_login")
	service.Listmsg(c, username)
}

func Liscon(c *gin.Context) {
	id := c.PostForm("id") //显示那个message下面的评论
	id2, _ := strconv.Atoi(id)
	service.Listcon(c, id2)
}

func Listmymsg(c *gin.Context) {
	service.Listmymsg(c)
}

func Nonamemsg(c *gin.Context) {
	touser := c.PostForm("touser")   //发送给谁
	message := c.PostForm("message") //message的内容
	fromuser := "noname"
	truename, _ := c.Cookie("now_user_login")

	err := service.Checkuseraliveser(touser)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, "没有该账户")
		return
	}
	if utf8.RuneCountInString(message) <= 5 {
		c.JSON(http.StatusOK, "评论长度应该大于5")
		return
	}

	err = service.Setmessage(touser, fromuser, message, truename)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, "发送失败")
		return
	}
	c.JSON(http.StatusOK, "发送成功")
}

func Nonamecom(c *gin.Context) {
	id := c.PostForm("id")           //为哪条id品论
	message := c.PostForm("message") //message的内容
	fromuser := "noname"
	truename, _ := c.Cookie("now_user_login")
	id2, _ := strconv.Atoi(id)
	fmt.Println(id2)
	if dao.Querypid(id2) == false {
		c.JSON(http.StatusOK, "没有这条评论")
		return
	}
	if utf8.RuneCountInString(message) <= 5 {
		c.JSON(http.StatusOK, "评论长度应该大于5")
		return
	}

	err := service.Setcommnet(id2, fromuser, message, truename)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, "发送失败")
		return
	}
	c.JSON(http.StatusOK, "发送成功")
}

func Likes(c *gin.Context) {
	id := c.PostForm("id") //为哪条评论点赞
	id2, _ := strconv.Atoi(id)
	if dao.Querypid(id2) == false {
		c.JSON(http.StatusOK, "没有这条评论")
		return
	}
	err := service.Likes(id2)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, "点赞失败")
		return
	}
	c.JSON(http.StatusOK, "点赞成功")
}
