package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var(
	QuestionList map[int]*Question
)
type Question struct {
	Id              int     `json:"id"`
	Question        string  `json:"Question"`
	OptionA         string	`json:"Option_A"`
	OptionB         string	`json:"Option_B"`
	OptionC         string	`json:"Option_C"`
	OptionD         string	`json:"Option_D"`
	OptionE         string	`json:"Option_E"`
	OptionF         string	`json:"Option_F"`
	OptionG         string	`json:"Option_G"`
	CorrectAnswer1 	string	`json:"Correct_Answer_1"`
	CorrectAnswer2 	string	`json:"Correct_Answer_2"`
	CorrectAnswer3 	string	`json:"Correct_Answer_3"`
	CorrectAnswer4 	string	`json:"Correct_Answer_4"`
	CorrectAnswer5 	string	`json:"Correct_Answer_5"`
	CorrectAnswer6 	string	`json:"Correct_Answer_6"`
	CorrectAnswer7 	string	`json:"Correct_Answer_7"`
}

func ViewQuestionAll()map[int]*Question  {

	QuestionList = make(map[int]*Question)

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
		var a Question
		err = rows.Scan(&a.Id,&a.Question,&a.OptionA,&a.OptionB,&a.OptionC,&a.OptionD,&a.OptionE,&a.OptionF,&a.OptionG,&a.CorrectAnswer1,&a.CorrectAnswer2,&a.CorrectAnswer3,&a.CorrectAnswer4,&a.CorrectAnswer5,&a.CorrectAnswer6,&a.CorrectAnswer7)
		if err != nil {
			log.Fatal(err)
		}

		QuestionList[nums]=&Question{a.Id,a.Question,a.OptionA,a.OptionB,a.OptionC,a.OptionD,a.OptionE,a.OptionF,a.OptionG,a.CorrectAnswer1,a.CorrectAnswer2,a.CorrectAnswer3,a.CorrectAnswer4,a.CorrectAnswer5,a.CorrectAnswer6,a.CorrectAnswer7}
		nums=nums+1

	}


	return QuestionList
}