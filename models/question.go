package models

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/sx202/blog_api/comm"
	"github.com/sx202/blog_api/database"
)

//使用切片存储题库中id字段的值，这样做的原因是：在数据库中id字段虽然是不重复递增，
//但是并不连续，为了能够顺序遍历，所以就把id字段的值插入到一个切片中，然后通过遍历
//切片的方式来达到遍历题库中每一条记录的目的。
var QuestionID []int

//返回所有题的id的值
func GetQuestionId() (s []int,err error)  {
	
	if len(QuestionID) < 1 {

		db,err := database.LinkDb()
		if err != nil {
			return QuestionID,err
		}

		rows,err := db.Query("SELECT id FROM question_063")
		if err != nil{
			return QuestionID,err
		}

		for rows.Next(){
			var id int
			err = rows.Scan(&id)
			if err != nil{
				return QuestionID,err
			}
			QuestionID = append(QuestionID,id)
		}
	}

	return QuestionID,err
}

//获取指定题的内容
func GetQuestion(num int)(m comm.Question,err error) {

	var question comm.Question


	db,err := database.LinkDb()
	if err != nil {
		return question,err
	}

	stmt,err := db.Prepare("SELECT * FROM question_063 WHERE id = ? LIMIT 1")
	if err != nil {
		return question,err
	}

	var b comm.Question
	err=stmt.QueryRow(num).Scan(&b.Id,&b.Question, &b.OptionA,&b.OptionB,&b.OptionC,&b.OptionD,&b.OptionE,&b.OptionF,&b.OptionG,&b.CorrectAnswer1,&b.CorrectAnswer2,&b.CorrectAnswer3,&b.CorrectAnswer4,&b.CorrectAnswer5,&b.CorrectAnswer6,&b.CorrectAnswer7)
	if err != nil {
		return question,err
	}

	question = comm.Question{b.Id,b.Question,b.OptionA,b.OptionB,b.OptionC,b.OptionD,b.OptionE,b.OptionF,b.OptionG,b.CorrectAnswer1,b.CorrectAnswer2,b.CorrectAnswer3,b.CorrectAnswer4,b.CorrectAnswer5,b.CorrectAnswer6,b.CorrectAnswer7}
	question = b

	return question,err
}

//获取所有题的内容
func GetAllQuestion()(m map[int]*comm.Question,err error) {

	QuestionList := make(map[int]*comm.Question)


	db,err := database.LinkDb()
	if err != nil {
		return QuestionList,err
	}

	//这条查询语句，因为是出题系统，数据量不会太大，所以没有加限制条件
	rows,err := db.Query("SELECT * FROM question_063")
	if err != nil {
		return QuestionList,err
	}

	nums := 0
	for rows.Next() {
		var a comm.Question
		err = rows.Scan(&a.Id,&a.Question,&a.OptionA,&a.OptionB,&a.OptionC,&a.OptionD,&a.OptionE,&a.OptionF,&a.OptionG,&a.CorrectAnswer1,&a.CorrectAnswer2,&a.CorrectAnswer3,&a.CorrectAnswer4,&a.CorrectAnswer5,&a.CorrectAnswer6,&a.CorrectAnswer7)
		if err != nil {
			return QuestionList,err
		}
		QuestionList[nums]=&comm.Question{a.Id,a.Question,a.OptionA,a.OptionB,a.OptionC,a.OptionD,a.OptionE,a.OptionF,a.OptionG,a.CorrectAnswer1,a.CorrectAnswer2,a.CorrectAnswer3,a.CorrectAnswer4,a.CorrectAnswer5,a.CorrectAnswer6,a.CorrectAnswer7}
		nums=nums+1
	}

	return QuestionList,err
}

//往数据库中插入一道题
func InsertQuestion(question comm.Question)(err error)  {

	db,err := database.LinkDb()
	if err != nil {
		return err
	}

	tx,err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO question_063(Question,Option_A,Option_B,Option_C,Option_D, Option_E,Option_F,Option_G,Correct_Answer_1, Correct_Answer_2,Correct_Answer_3, Correct_Answer_4,Correct_Answer_5, Correct_Answer_6,Correct_Answer_7) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(question.Question,question.OptionA,question.OptionB,question.OptionC,question.OptionD,question.OptionE,question.OptionF,question.OptionG,question.CorrectAnswer1,question.CorrectAnswer2,question.CorrectAnswer3,question.CorrectAnswer4,question.CorrectAnswer5,question.CorrectAnswer6,question.CorrectAnswer7)
	if err != nil {
		return err

	}

	//不是很清楚这个数字代表的是什么，暂时认为可以返回这个值，代表是第几道题插入成功
	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}

//更新题库中指定的一条题库内容
func UpdateQuestion(newquestion comm.Question)(err error)  {

	m := map[string]string{}

	oldquestion,err := GetQuestion(newquestion.Id)
	if err != nil {
		return err
	}

	if newquestion.Question != oldquestion.Question {
		m["Question"] = newquestion.Question
	}
	///这是优化代码，没有试验不太感确定
	//s := [7]string{"A","B","C","D","E","F","G"}
	//for i:=0;i<len(s);i++ {
	//	option := "Option"+s[i]
	//	newOption := "newquestion.Option"+ s[i]
	//	oldOption := "oldquestion.Option"+ s[i]
	//
	//	if  newOption != oldOption{
	//		m[option] = newOption
	//	}
	//}
	///这段代码和上面的代码功能一样
	if newquestion.OptionA != oldquestion.OptionA {
		m["OptionA"] = newquestion.OptionA
	}
	if newquestion.OptionB != oldquestion.OptionB {
		m["OptionB"] = newquestion.OptionB
	}
	if newquestion.OptionC != oldquestion.OptionC {
		m["OptionC"] = newquestion.OptionC
	}
	if newquestion.OptionD != oldquestion.OptionD {
		m["OptionD"] = newquestion.OptionD
	}
	if newquestion.OptionE != oldquestion.OptionE {
		m["OptionE"] = newquestion.OptionE
	}
	if newquestion.OptionF != oldquestion.OptionF {
		m["OptionF"] = newquestion.OptionF
	}
	if newquestion.OptionG != oldquestion.OptionG {
		m["OptionG"] = newquestion.OptionG
	}

	if newquestion.CorrectAnswer1 != oldquestion.CorrectAnswer1 {
		m["CorrectAnswer1"] = newquestion.CorrectAnswer1
	}
	if newquestion.CorrectAnswer2 != oldquestion.CorrectAnswer2 {
		m["CorrectAnswer2"] = newquestion.CorrectAnswer2
	}
	if newquestion.CorrectAnswer3 != oldquestion.CorrectAnswer3 {
		m["CorrectAnswer3"] = newquestion.CorrectAnswer3
	}
	if newquestion.CorrectAnswer4 != oldquestion.CorrectAnswer4 {
		m["CorrectAnswer4"] = newquestion.CorrectAnswer4
	}
	if newquestion.CorrectAnswer5 != oldquestion.CorrectAnswer5 {
		m["CorrectAnswer5"] = newquestion.CorrectAnswer5
	}
	if newquestion.CorrectAnswer6 != oldquestion.CorrectAnswer6 {
		m["CorrectAnswer6"] = newquestion.CorrectAnswer6
	}
	if newquestion.CorrectAnswer7 != oldquestion.CorrectAnswer7 {
		m["CorrectAnswer7"] = newquestion.CorrectAnswer7
	}

	//字符串拼接
	update := "update question_063 set "
	column := ""
	for key,value := range m {
		column = column + key + "=" + value + " "
	}
	update = update + column + "where id=?"

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

	_,err = stmt.Exec(newquestion.Id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}

//删除指定考题
func DeleteQuestion(id int)(err error)  {

	db,err := database.LinkDb()
	if err != nil {
		return err
	}

	tx,err := db.Begin()
	if err != nil {
		return err
	}

	stmt,err := db.Prepare("DELETE FROM question_063 WHERE id=?")
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
