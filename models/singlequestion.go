package models

import (
	"database/sql"
	"log"
)

var(
	singlequestion map[int]*Question
)

func SingleQuestion()map[int]*Question  {
	singlequestion = make(map[int]*Question)

	db,err := sql.Open("sqlite3","./database/blog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows,err := db.Query("select * from question_063")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()





	return singlequestion
}

