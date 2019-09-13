package controllers

import (
	"github.com/astaxie/beego"
	"github.com/sx202/blog_api/models"
)

type System struct {
	beego.Controller
}

func (c *System) GetQuestion()  {

	//注意这个地方传数据没有用json格式，而是简单的传了一个字符（题号）
	txt := c.Ctx.Input.RequestBody

	singlequestion,err := models.GetQuestion(string(txt))
	if err == nil {
		c.Data["json"] = singlequestion
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

}

func (c *System) UpdateQuestion ()  {

}

func (c *System) DeleteQuestion ()  {

}
