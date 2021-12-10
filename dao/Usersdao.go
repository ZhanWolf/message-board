package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"message-board/Struct"
)

func Queryuserpassword(username string) string {
	U := new(Struct.User)
	err := Db.QueryRow("select username,password,id from user where username = ?;", username).Scan(&U.Username, &U.Password, &U.Id)
	if err != nil {
		fmt.Println("错误:", err)
	}

	return U.Password
}

func Queryusername(username string) error {
	U := new(Struct.User)
	err := Db.QueryRow("select username,id from user where username = ?;", username).Scan(&U.Username, &U.Id)
	if err != nil {
		fmt.Println("查询错误", err)
		return err
	}
	return nil
}

func Insertuser(username string, password string, protectionQ string, protectionA string) error {
	_, err := Db.Exec("insert into user(username,password,protectionQ,protectionA,) values (?,?,?,?,);", username, password, protectionQ, protectionA)
	if err != nil {
		fmt.Println("插入错误", err)
	}
	return nil
}

func Updatepassword(newpassword string, username string) error {
	_, err := Db.Exec("update user set password=? where username=?;", newpassword, username)
	return err
}

func Queryprotection(username string) (string, string) {
	U := new(Struct.User)
	err := Db.QueryRow("select protectionQ,protectionA from user where username=?;", username).Scan(&U.ProtectionQ, &U.ProtectionA)
	if err != nil {
		fmt.Println(err)
	}
	return U.ProtectionQ, U.ProtectionA
}
