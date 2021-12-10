package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/service"
	"net/http"
	"time"
	"unicode/utf8"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	err := service.Checkuseraliveser(username)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, "没有此账户")
		return
	}
	cookie := service.UserLoginser(username, password)
	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, gin.H{
		"登录成功":   "",
		"Hello!": username,
	})

}

func Singup(c *gin.Context) {
	username := c.PostForm("username")           //用户名
	password := c.PostForm("password")           //密码
	passwordagain := c.PostForm("passwordagain") //重复输入密码
	protectionQ := c.PostForm("protectionQ")     //密保问题
	protectionA := c.PostForm("protectionA")     //密保答案

	err := service.Checkuseraliveser(username)
	if err == nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, "该账户已存在")
		return
	}
	if utf8.RuneCountInString(password) <= 3 {
		c.JSON(http.StatusOK, "密码长度应该大于3")
		return
	}

	cookie, flag := service.UserSingup(username, password, passwordagain, protectionQ, protectionA)
	if flag == false {
		c.JSON(http.StatusOK, "两次密码输入不正确")
		return
	}
	if utf8.RuneCountInString(password) <= 3 {
		c.JSON(http.StatusOK, "密码长度应该大于3")
	}
	c.JSON(http.StatusOK, "注册成功")
	http.SetCookie(c.Writer, cookie)
}

func Reset(c *gin.Context) {
	username := c.PostForm("username") //用户名
	password := c.PostForm("password") //密码
	passwordagain := c.PostForm("passwordagain")
	protectionA := c.PostForm("protectionA") //密保答案

	err := service.Checkuseraliveser(username)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, "没有此账户")
		return
	}
	service.PasswordReset(c, username, password, protectionA, passwordagain)
}

func Clock(c *gin.Context) {
	username, _ := c.Cookie("now_user_login")
	c.JSON(http.StatusOK, gin.H{
		"hello": username,
		"现在时间":  time.Now(),
	})

}
