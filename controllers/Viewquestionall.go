package controllers

import (
	"github.com/astaxie/beego"
	"github.com/sx202/blog_api/models"
)

type Viewquestionall struct {
	beego.Controller
}

func (c *Viewquestionall) Questionall()  {
	questionall := models.ViewQuestionAll()
	c.Data["json"] = questionall
	c.ServeJSON()
}