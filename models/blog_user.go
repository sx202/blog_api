package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Bloguser struct {
	id       int        `json:"id"`
	username string		`json:"username"`
	password string		`json:"password"`
	email    string		`json:"email"`
	roles    int		`json:"roles"`
}

func BlogAllUser()  {
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
		var ID       int
		var USERNAME string
		var PASSWORD string
		var EMAIL    string
		var ROLES    int
		err = rows.Scan(&ID,&USERNAME,&PASSWORD,&EMAIL,&ROLES)
		if err != nil {
			log.Fatal(err)
		}

		var a map[string] *Bloguser
		 = ID
		a.username = USERNAME
		a.password = PASSWORD
		a.email = EMAIL
		a.roles = ROLES

		fmt.Println(a)



		aa,err := json.Marshal(a)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(aa))
	}

}
