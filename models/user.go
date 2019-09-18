package models

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/sx202/blog_api/comm"
	"github.com/sx202/blog_api/database"
)

var UserId  []int

func GetUserId()(s []int,err error)  {

	if len(UserId) < 1 {
		db,err := database.LinkDb()
		if err != nil {
			return UserId,err
		}

		rows,err := db.Query("SELECT id FROM user")
		if err != nil {
			return UserId,err
		}

		for rows.Next() {
			var id int
			err = rows.Scan(&id)
			if err != nil {
				return UserId,err
			}
			UserId = append(UserId,id)
		}
	}

	return UserId,err
}

func Login(user comm.User)(err error) {

	db,err := database.LinkDb()
	if err != nil {
		return err
	}

	rows,err :=db.Query("SELECT username,password FROM user")
	if err != nil {
		return err
	}

	for rows.Next() {
		var username string
		var password string
		err = rows.Scan(&username,&password)
		if err != nil {
			return err
		}
		if username == user.UserName && password == user.PassWord {
			err = nil
		}else {
			err.Error() = "username or password is not Correct!"
		}
	}
	return err
}

func GetUser(num int)(m comm.User,err error)  {

	var user comm.User

	db,err := database.LinkDb()
	if err != nil {
		return user,err
	}

	stmt,err := db.Prepare("SELECT * FROM user WHERE id = ? LIMIT 1")
	if err != nil {
		return user,err
	}

	var b comm.User
	err = stmt.QueryRow(num).Scan(&b.Id,&b.UserName,&b.PassWord,&b.Email,&b.Roles)
	if err != nil {
		return user,err
	}

	//user = comm.User{b.Id,b.UserName,b.PassWord,b.Email,b.Roles}
	user = b

	return user,err
}

func GetAllUsers()(m map[int]*comm.User,err error)  {

	userlist := make(map[int]*comm.User)

	db,err := database.LinkDb()
	if err != nil {
		return userlist,err
	}

	rows,err := db.Query("SELECT * FROM user")
	if err != nil {
		return userlist,err
	}

	nums := 0
	for rows.Next() {
		var a comm.User
		err = rows.Scan(&a.Id,&a.UserName,&a.PassWord,&a.Email,&a.Roles)
		if err != nil {
			return userlist,err
		}
		userlist[nums] = &a
		nums = nums + 1
	}

	return userlist,err
}



func AddUser(user comm.User)(err error)  {

	db,err := database.LinkDb()
	if err != nil {
		return err
	}

	tx,err := db.Begin()
	if err != nil{
		return err
	}

	stmt,err := db.Prepare("INSERT INTO user(username,password,email,roles) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}

	_,err = stmt.Exec(user.UserName,user.PassWord,user.Email,user.Roles)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

func UpdateUser(newuser comm.User)(err error)  {

	 u := map[string]string{}

	 olduser,err := GetUser(newuser.Id)
	 if err != nil {
	 	return err
	 }

	 if newuser.UserName != olduser.UserName {
	 	u["Username"] = newuser.UserName
	 }

	 if newuser.PassWord != olduser.PassWord {
	 	u["PassWord"] = newuser.PassWord
	 }

	 if newuser.Email != olduser.Email {
	 	u["Email"] = newuser.Email
	 }

	 if newuser.Roles != olduser.Roles {
	 	u["Roles"] = newuser.Roles
	 }
	 //字符串拼接
	 update := "UPDATE user SET "
	 column := ""
	 for key,value := range u {
	 	column = column + key + "=" + value + " "
	 }
	 update = update + column + "WHERE id=?"

	 db,err := database.LinkDb()
	 if err != nil {
	 	return err
	 }

	 tx,err := db.Begin()
	 if err != nil {
	 	return err
	 }

	 stmt,err := db.Prepare(update)
	 if err != nil {
	 	return err
	 }

	 _,err = stmt.Exec(newuser.Id)
	 if err != nil {
	 	return err
	 }

	 err = tx.Commit()
	 if err != nil {
	 	return err
	 }

	return err
}

func DeleteUser(id int)(err error)  {

	db,err := database.LinkDb()
	if err != nil {
		return err
	}

	tx,err := db.Begin()
	if err != nil {
		return err
	}

	stmt,err := db.Prepare("DELETE FROM user WHERE id=? ")
	if err != nil {
		return err
	}

	_,err = stmt.Exec(id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}