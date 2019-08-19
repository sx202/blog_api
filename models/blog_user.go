package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var(
	UserAllList map[int]*bloguser
)

type bloguser struct {
	Id       int        `json:"Id"`
	Username string		`json:"Username"`
	Password string		`json:"Password"`
	Email    string		`json:"Email"`
	Roles    int		`json:"Roles"`
}

func BlogAllUser()map[int]*bloguser  {

	UserAllList = make(map[int]*bloguser)

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

		var a bloguser

		err = rows.Scan(&a.Id,&a.Username,&a.Password,&a.Email,&a.Roles)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(a)

		//aa,err := json.Marshal(a)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Println(string(aa))


		//bb := string(aa)

		UserAllList[a.Id] = &bloguser{a.Id,a.Username,a.Password,a.Email,a.Roles}


	}
	return UserAllList
}
