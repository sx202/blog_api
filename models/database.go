package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sx202/blog_api/comm"
	"log"
	"strconv"
)

//使用切片存储题库中id字段的值，这样做的原因是：在数据库中id字段虽然是不重复递增，
//但是并不连续，为了能够顺序遍历，所以就把id字段的值插入到一个切片中，然后通过遍历
//切片的方式来达到遍历题库中每一条记录的目的。
var ID []int

//初始化数据库的连接
func LinkDb()*sql.DB  {

	db,err := sql.Open("sqlite3","./database/blog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}

//获取题库中每个题的id值,没有返回值
func getId()  {

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

	db.Close()
}

//返回所有题的id的值
func GetId() []int  {
	
	if len(ID) <1 {
		getId()
	}
	return ID
}

//获取指定题的内容
func GetSingleQuestion(txt string)map[int]*comm.Question {

	db := LinkDb()
	
	//把字符类型转换成数值类型
	a,err := strconv.Atoi(txt)
	if err != nil {
		log.Fatal(err)
	}

	question := make(map[int]*comm.Question)

	stmt,err := db.Prepare("select * from question_063 where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var b comm.Question
	err=stmt.QueryRow(a).Scan(&b.Question,
		&b.OptionA,&b.OptionB,&b.OptionC,&b.OptionD,&b.OptionE,&b.OptionF,&b.OptionG,&b.CorrectAnswer1,&b.CorrectAnswer2,&b.CorrectAnswer3,&b.CorrectAnswer4,&b.CorrectAnswer5,&b.CorrectAnswer6,&b.CorrectAnswer7)

	question[a] = &comm.Question{b.Id,b.Question,b.OptionA,b.OptionB,b.OptionC,b.OptionD,b.OptionE,b.OptionF,b.OptionG,b.CorrectAnswer1,b.CorrectAnswer2,b.CorrectAnswer3,b.CorrectAnswer4,b.CorrectAnswer5,b.CorrectAnswer6,b.CorrectAnswer7}

	db.Close()
	return question
}

//获取所有题的内容
func GetAllQuestion()map[int]*comm.Question  {

	QuestionList := make(map[int]*comm.Question)

	db := LinkDb()

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

	db.Close()
	return QuestionList
}

//往数据库中插入一道题
func InsertSingleQuestion(question comm.Question)int64  {

	db := LinkDb()


	tx,err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("insert into question_063(Question,Option_A,Option_B,Option_C,Option_D, Option_E,Option_F,Option_G,Correct_Answer_1, Correct_Answer_2,Correct_Answer_3, Correct_Answer_4,Correct_Answer_5, Correct_Answer_6,Correct_Answer_7) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(question.Question,question.OptionA,question.OptionB,question.OptionC,question.OptionD,question.OptionE,question.OptionF,question.OptionG,question.CorrectAnswer1,question.CorrectAnswer2,question.CorrectAnswer3,question.CorrectAnswer4,question.CorrectAnswer5,question.CorrectAnswer6,question.CorrectAnswer7)
	if err != nil{
		log.Fatal(err)
	}

	//不是很清楚这个数字代表队的是什么，暂时认为可以返回这个值，代表是第几道题插入成功
	id,err := res.LastInsertId()
	if err != nil{
		log.Fatal(err)
	}

	tx.Commit()

	db.Close()

	return id

}

