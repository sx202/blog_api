package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sx202/blog_api/comm"
	"log"

)

var(
	QuestionList map[int]*comm.Question
)


func ViewQuestionAll()map[int]*comm.Question  {

	QuestionList = make(map[int]*comm.Question)

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

	nums := 0

	for rows.Next() {
		var a comm.Question
		err = rows.Scan(&a.Id,&a.Question,&a.OptionA,&a.OptionB,&a.OptionC,&a.OptionD,&a.OptionE,&a.OptionF,&a.OptionG,&a.CorrectAnswer1,&a.CorrectAnswer2,&a.CorrectAnswer3,&a.CorrectAnswer4,&a.CorrectAnswer5,&a.CorrectAnswer6,&a.CorrectAnswer7)
		if err != nil {
			log.Fatal(err)
		}

		QuestionList[nums]=&comm.Question{a.Id,a.Question,a.OptionA,a.OptionB,a.OptionC,a.OptionD,a.OptionE,a.OptionF,a.OptionG,a.CorrectAnswer1,a.CorrectAnswer2,a.CorrectAnswer3,a.CorrectAnswer4,a.CorrectAnswer5,a.CorrectAnswer6,a.CorrectAnswer7}
		nums=nums+1

	}


	return QuestionList
}