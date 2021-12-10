package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/Struct"
	"net/http"
	"time"
)

func Insertmessage(fromusername string, tousername string, messagecontent string, truename string) error {
	time := time.Now()
	_, err := Db.Exec("insert into message(fromuser,touser,messagecontent,messagetime,truename,likes) values(?,?,?,?,?,0);", fromusername, tousername, messagecontent, time, truename)
	return err
}

func Insertcomment(pid int, fromusername string, messagecontent string, tousername string, truename string) {
	time := time.Now()
	_, err := Db.Exec("insert into message(fromuser,pid,messagecontent,messagetime,touser,truename,likes) values(?,?,?,?,?,?,0);", fromusername, pid, messagecontent, time, tousername, truename)
	if err != nil {
		fmt.Println(err)
	}
}

func Querycommentuser(id int) (username string, err error) {
	err = Db.QueryRow("select  touser from  message where id=?;", id).Scan(&username)
	return username, err
} //通过id查询username

func Queryuserallcmsg(c *gin.Context, username string) {
	var Ms Struct.Message
	sqlStr := "select messagecontent,id,fromuser,messagetime from message where touser=?;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, username)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		c.JSON(http.StatusOK, "无评论")
		goto sign
	}

	for rows.Next() {
		err := rows.Scan(&Ms.Messagecontent, &Ms.Id, &Ms.Fromusername, &Ms.Time)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		time := utos(Ms.Time)
		c.JSON(http.StatusOK, gin.H{
			"该评论id为": Ms.Id,
			"来自":     Ms.Fromusername,
			"评论了":    Ms.Messagecontent,
			"时间":     time,
		})
	}
	rows.Close()
sign:
}

func Querymsgallcon(c *gin.Context, pid int) {
	var Ms Struct.Message
	sqlStr := "select messagecontent,id,fromuser,messagetime from message where pid=?;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, pid)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		c.JSON(http.StatusOK, "无评论")
		goto sign
	}

	for rows.Next() {
		err := rows.Scan(&Ms.Messagecontent, &Ms.Id, &Ms.Fromusername, &Ms.Time)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		time := utos(Ms.Time)
		c.JSON(http.StatusOK, gin.H{
			"回复的评论id为": pid,
			"该评论id为":   Ms.Id,
			"来自":       Ms.Fromusername,
			"评论了":      Ms.Messagecontent,
			"时间":       time,
		})
	}
	rows.Close()
sign:
}

func Querypid(pid int) bool {
	var fromusername string
	Db.QueryRow("select fromuser from message where id=?;", pid).Scan(&fromusername)
	if fromusername == "" {
		return false
	}
	return true
}

func Deletemsg(id int) error {
	_, err := Db.Exec("delete from message where id = ?;", id)
	return err
}

func Updatemsg(messsagecomment string, id int) error {
	time := time.Now()
	_, err := Db.Exec("update message set messagecontent=?,messagetime=? where id=?;", messsagecomment, time, id)
	return err
}

func Checkueser(id int) string {
	OpenDb()
	var username string
	Db.QueryRow("select fromuser from message where id=?;", id).Scan(&username)
	return username
}

func utos(u []uint8) string {
	by := []byte{}
	for _, b := range u {
		by = append(by, b)
	}
	return string(by)
}
func Querymymsg(c *gin.Context) {
	var Ms Struct.Message
	myname, _ := c.Cookie("now_user_login")
	sqlStr := "select messagecontent,id,fromuser,messagetime from message where fromuser=? or truename=?;" //遍历写给登录用户的评论
	rows, err := Db.Query(sqlStr, myname, myname)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		c.JSON(http.StatusOK, "无评论")
		goto sign
	}

	for rows.Next() {
		err := rows.Scan(&Ms.Messagecontent, &Ms.Id, &Ms.Fromusername, &Ms.Time)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		time := utos(Ms.Time)
		c.JSON(http.StatusOK, gin.H{
			"该评论id为": Ms.Id,
			"来自":     Ms.Fromusername,
			"评论了":    Ms.Messagecontent,
			"时间":     time,
		})
	}
	rows.Close()
sign:
}

func Updatelke(likes, id int) error {
	_, err := Db.Exec("update message set likes=? where id=?;", likes, id)
	return err
}

func Querylke(id int) (likes int, err error) {
	err = Db.QueryRow("select likes from  message where id=?;", id).Scan(&likes)

	if likes == 0 {
		likes = 0
	}
	likes += 1
	return likes, err
} //通过id查询likes

func Checktruename(id int) (username string) {
	OpenDb()
	Db.QueryRow("select truename from message where id=?;", id).Scan(&username)
	return username
}
