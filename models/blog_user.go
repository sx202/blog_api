package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sx202/blog_api/comm"
	"log"
)

var(
	UserAllList map[int]*comm.Bloguser
)



func BlogAllUser()map[int]*comm.Bloguser  {

	UserAllList = make(map[int]*comm.Bloguser)

	db,err := sql.Open("sqlite3","./database/blog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows,err := db.Query("select * from blog_user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next(){

		var a comm.Bloguser

		err = rows.Scan(&a.Id,&a.Username,&a.Password,&a.Email,&a.Roles)
		if err != nil {
			log.Fatal(err)
		}


		UserAllList[a.Id] = &comm.Bloguser{a.Id,a.Username,a.Password,a.Email,a.Roles}


	}
	return UserAllList
}
