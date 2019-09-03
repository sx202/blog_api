package controllers

import (
	"github.com/astaxie/beego"
	"github.com/sx202/blog_api/models"
)

type SingleQuestion struct {
	beego.Controller
}

func (c *SingleQuestion) singleQuestion() {
	singlequestion := models.SingleQuestion()
	c.Data["json"] = singlequestion
	c.ServeJSON()
}