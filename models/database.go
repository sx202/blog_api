package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sx202/blog_api/comm"
	"log"
	"strconv"
)

var ID []int

func LinkDb()*sql.DB  {

	db,err := sql.Open("sqlite3","./database/blog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}

func QueryId()  {

	db := LinkDb()

	rows,err := db.Query("select id from quesiton_063")
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next(){
		var id int
		err = rows.Scan(&id)
		if err != nil{
			log.Fatal(err)
		}
		ID = append(ID,id)
	}
}

func QuerySingleQuestion(txt string)*comm.Question {

	db := LinkDb()

	if len(ID) <1 {
		QueryId()
	}

	a,err := strconv.Atoi(txt)
	if err != nil {
		log.Fatal(err)
	}

	stmt,err := db.Prepare("select * from question_063 where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var singlequestion *comm.Question
	err=stmt.QueryRow(ID[a]).Scan(&singlequestion.Id,&singlequestion.Question,&singlequestion.OptionA,&singlequestion.OptionB,&singlequestion.OptionC,&singlequestion.OptionD,&singlequestion.OptionE,&singlequestion.OptionF,&singlequestion.OptionG,&singlequestion.CorrectAnswer1,&singlequestion.CorrectAnswer2,&singlequestion.CorrectAnswer3,&singlequestion.CorrectAnswer4,&singlequestion.CorrectAnswer5,&singlequestion.CorrectAnswer6,&singlequestion.CorrectAnswer7)


	return singlequestion
}