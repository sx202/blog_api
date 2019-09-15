package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/sx202/blog_api/comm"
	"github.com/sx202/blog_api/models"
	"strconv"
)

type System struct {
	beego.Controller
}

func (c *System) GetQuestionId ()  {

	id,err := models.GetQuestionId()
	if err == nil {
		//这个地方需要测试一下，id是切片，不知道这样是否能把切片数据转换成json格式传输出去
		c.Data["json"] = id
	}else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *System) GetQuestion()  {

	txt := c.Ctx.Input.RequestBody

	num,err := strconv.Atoi(string(txt))
	if err == nil {
		singlequestion,err := models.GetQuestion(num)
		if err == nil {
			c.Data["json"] = singlequestion
		}else {
			c.Data["json"] = err.Error()
		}
	}else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *System) GetAllQuestion()  {

	questionall,err := models.GetAllQuestion()
	if err == nil{
		c.Data["json"] = questionall
	}else {
		c.Data["json"] = err.Error()

	}
	c.ServeJSON()
}

func (c *System) InsertQuestion ()  {

	var question comm.Question

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &question)
	if err == nil {
		err = models.InsertQuestion(question)
		if err == nil {
			c.Data["json"] = []string{"newquestion","ok"}
		}else {
			c.Data["json"] = err.Error()
		}
	}else {
		c.Data["json"] = err.Error()
	}

	c.ServeJSON()
}

func (c *System) UpdateQuestion ()  {

	var question comm.Question

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &question)
	if err == nil {
		err = models.UpdateQuestion(question)
		if err == nil {
			c.Data["json"] = []string{"updatequestion","ok"}
		}else {
			c.Data["json"] = err.Error()
		}
	}else {
		c.Data["json"] = err.Error()
	}

	c.ServeJSON()

}

func (c *System) DeleteQuestion ()  {

	txt := c.Ctx.Input.RequestBody

	num,err := strconv.Atoi(string(txt))
	if err == nil {
		err = models.DeleteQuestion(num)
		if err == nil {
			c.Data["json"] = []string{"deletequestion","ok"}
		}else {
			c.Data["json"] = err.Error()
		}
	}else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
