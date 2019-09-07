package controllers

import (
	"github.com/astaxie/beego"
	"github.com/sx202/blog_api/models"
)

type SingleQuestion struct {
	beego.Controller
}

func (c *SingleQuestion) Singlequestion() {

	//c.Ctx.WriteString("hello")
	txt := c.Ctx.Input.RequestBody
	c.Ctx.WriteString(string(txt))
	singlequestion := models.GetSingleQuestion(string(txt))
	c.Data["json"] = singlequestion
	c.ServeJSON()
}