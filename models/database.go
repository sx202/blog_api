package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sx202/blog_api/comm"
	"strconv"
)

//使用切片存储题库中id字段的值，这样做的原因是：在数据库中id字段虽然是不重复递增，
//但是并不连续，为了能够顺序遍历，所以就把id字段的值插入到一个切片中，然后通过遍历
//切片的方式来达到遍历题库中每一条记录的目的。
var ID []int

//初始化数据库的连接
func LinkDb()(db *sql.DB,err error) {

	db,err = sql.Open("sqlite3","./database/blog.db")
	if err != nil {
		return nil,err
	}
	return db,err
}

//获取题库中每个题的id值,没有返回值
func getId()(err error) {

	db,err := LinkDb()

	if err == nil {
		rows,err := db.Query("SELECT id FROM quesiton_063")
		if err == nil{
			for rows.Next(){
				var id int
				err = rows.Scan(&id)
				if err == nil{
					ID = append(ID,id)
					err = nil
				}
			}
		}
	}
	return err
}

//返回所有题的id的值
func GetId() (s []int,err error)  {
	
	if len(ID) <1 {
		err := getId()
		if err != nil {
			return nil,err
		}
	}
	return ID,nil
}

//获取指定题的内容
func GetSingleQuestion(txt string)(m map[int]*comm.Question,err error) {

	question := make(map[int]*comm.Question)
	question=nil

	db,err := LinkDb()
	if err == nil {
		//把字符类型转换成数值类型
		a,err := strconv.Atoi(txt)
		if err == nil {
			stmt,err := db.Prepare("SELECT * FROM question_063 WHERE id = ? LIMIT 1")
			if err == nil {
				var b comm.Question
				err=stmt.QueryRow(a).Scan(&b.Question, &b.OptionA,&b.OptionB,&b.OptionC,&b.OptionD,&b.OptionE,&b.OptionF,&b.OptionG,&b.CorrectAnswer1,&b.CorrectAnswer2,&b.CorrectAnswer3,&b.CorrectAnswer4,&b.CorrectAnswer5,&b.CorrectAnswer6,&b.CorrectAnswer7)
				if err == nil{
					question[a] = &comm.Question{b.Id,b.Question,b.OptionA,b.OptionB,b.OptionC,b.OptionD,b.OptionE,b.OptionF,b.OptionG,b.CorrectAnswer1,b.CorrectAnswer2,b.CorrectAnswer3,b.CorrectAnswer4,b.CorrectAnswer5,b.CorrectAnswer6,b.CorrectAnswer7}
				}
			}
		}
	}
	return question,err
}

//获取所有题的内容
func GetAllQuestion()(m map[int]*comm.Question,err error) {

	QuestionList := make(map[int]*comm.Question)
	QuestionList = nil

	db,err := LinkDb()
	if err == nil {
		//这条查询语句，因为是出题系统，数据量不会太大，所以没有加限制条件
		rows,err := db.Query("SELECT * FROM question_063")
		if err == nil {
			nums := 0
			for rows.Next() {
				var a comm.Question
				err = rows.Scan(&a.Id,&a.Question,&a.OptionA,&a.OptionB,&a.OptionC,&a.OptionD,&a.OptionE,&a.OptionF,&a.OptionG,&a.CorrectAnswer1,&a.CorrectAnswer2,&a.CorrectAnswer3,&a.CorrectAnswer4,&a.CorrectAnswer5,&a.CorrectAnswer6,&a.CorrectAnswer7)
				if err == nil {
					QuestionList[nums]=&comm.Question{a.Id,a.Question,a.OptionA,a.OptionB,a.OptionC,a.OptionD,a.OptionE,a.OptionF,a.OptionG,a.CorrectAnswer1,a.CorrectAnswer2,a.CorrectAnswer3,a.CorrectAnswer4,a.CorrectAnswer5,a.CorrectAnswer6,a.CorrectAnswer7}
					nums=nums+1
				}
			}
		}
	}
	return QuestionList,err
}

//往数据库中插入一道题
func InsertSingleQuestion(question comm.Question)(err error)  {

	db,err := LinkDb()
	if err == nil {
		tx,err := db.Begin()
		if err == nil {
			stmt, err := db.Prepare("INSERT INTO question_063(Question,Option_A,Option_B,Option_C,Option_D, Option_E,Option_F,Option_G,Correct_Answer_1, Correct_Answer_2,Correct_Answer_3, Correct_Answer_4,Correct_Answer_5, Correct_Answer_6,Correct_Answer_7) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
			if err == nil {
				_, err := stmt.Exec(question.Question,question.OptionA,question.OptionB,question.OptionC,question.OptionD,question.OptionE,question.OptionF,question.OptionG,question.CorrectAnswer1,question.CorrectAnswer2,question.CorrectAnswer3,question.CorrectAnswer4,question.CorrectAnswer5,question.CorrectAnswer6,question.CorrectAnswer7)
				if err == nil {
					//不是很清楚这个数字代表队的是什么，暂时认为可以返回这个值，代表是第几道题插入成功
					err = tx.Commit()
				}
			}
		}
	}
	return err
}

func UpdateSingleQuestion(newquestion *comm.Question,oldquestion *comm.Question)(err error)  {

	db,err := LinkDb()
	if err == nil {
		if newquestion.Question != oldquestion.Question {

		}else {
			if newquestion.OptionA != oldquestion.OptionA {

			}
		}
	}

	return err
}
