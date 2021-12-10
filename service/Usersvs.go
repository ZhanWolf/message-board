package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
	"net/http"
)

func UserLoginser(username string, password string) *http.Cookie {
	dao.OpenDb()
	turepassword := dao.Queryuserpassword(username)
	if turepassword != password {
		return nil
	}

	cookie := &http.Cookie{
		Name:     "now_user_login",
		Value:    username,
		MaxAge:   300,
		Path:     "/",
		HttpOnly: true,
	}
	return cookie
}

func Checkuseraliveser(username string) error {
	dao.OpenDb()
	err := dao.Queryusername(username)
	return err
}

func UserSingup(username string, password string, passwordagain string, protectionQ string, protectionA string) (*http.Cookie, bool) {
	dao.OpenDb()
	err := dao.Insertuser(username, password, protectionQ, protectionA)
	if err != nil {
		fmt.Println("注册错误", err)
	}
	if passwordagain != password {
		return nil, false
	}
	cookie := &http.Cookie{
		Name:     "now_user_login",
		Value:    username,
		MaxAge:   300,
		Path:     "/",
		HttpOnly: true,
	}
	return cookie, true
}

func PasswordReset(c *gin.Context, username string, password string, protectionA string, passwordagain string) {
	dao.OpenDb()
	protectionQ, trueprotectionA := dao.Queryprotection(username)
	if protectionQ == "" {
		c.JSON(http.StatusOK, "该账户未设置密保")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			username:  "你好！",
			"宁的密保问题是": protectionQ,
		})
	}

	if trueprotectionA == protectionA && password == passwordagain {
		dao.Updatepassword(password, username)
		c.JSON(http.StatusOK, "密码修改成功")
	} else if passwordagain != password && trueprotectionA == protectionA {
		c.JSON(http.StatusOK, "两次输入密码不相同")
	} else if trueprotectionA != protectionA {
		c.JSON(http.StatusOK, "密保答案错误")
	}

}
