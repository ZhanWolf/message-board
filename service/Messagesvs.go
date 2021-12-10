package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/dao"
)

func Setmessage(tousername string, fromusername string, messagecontent string, truename string) error {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
	}
	err = dao.Queryusername(tousername)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}

	err = dao.Insertmessage(fromusername, tousername, messagecontent, truename)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil
}

func Setcommnet(pid int, fromusername string, mesagecontent string, truename string) error {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
	}

	tousername, _ := dao.Querycommentuser(pid)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	dao.Insertcomment(pid, fromusername, mesagecontent, tousername, truename)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	return nil
}

func Listmsg(c *gin.Context, username string) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	dao.Queryuserallcmsg(c, username)
}

func Listcon(c *gin.Context, pid int) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	dao.Querymsgallcon(c, pid)
}

func Deletemsg(id int) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	dao.Deletemsg(id)
}

func Updatemsg(messagecomment string, id int) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = dao.Updatemsg(messagecomment, id)
	if err != nil {
		fmt.Println(err)
	}
}

func Listmymsg(c *gin.Context) {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	dao.Querymymsg(c)
}

func Likes(id int) error {
	err := dao.OpenDb()
	if err != nil {
		fmt.Println(err)
		return err
	}
	likes, err := dao.Querylke(id)
	if err != nil {
		fmt.Println(err, 1)
		return err
	}
	err = dao.Updatelke(likes, id)
	if err != nil {
		fmt.Println(err, 2)
		return err
	}
	return nil
}
