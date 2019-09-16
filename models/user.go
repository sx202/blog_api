package models

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/sx202/blog_api/comm"
	"github.com/sx202/blog_api/database"
)

func Login(user comm.User)(err error,b bool) {

	b = false

	db,err := database.LinkDb()
	if err == nil {
		rows,err :=db.Query("SELECT username,password FROM user")
		if err == nil {
			for rows.Next() {
				var username string
				var password string
				err = rows.Scan(&username,&password)
				if err == nil {
					if username == user.UserName && password == user.PassWord {
						b = true
						err = nil
					}
				}
			}
		}
	}
	return err,b
}

func AddUser()  {

}

func GetUser()  {

}

func GetAllUsers()  {

}

func UpdateUser()  {

}

func DeleteUser()  {

}