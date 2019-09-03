package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type SingleQuestion struct {
	beego.Controller
}

func (c *SingleQuestion) singleQuestion() {

	txt := c.Ctx.Input.RequestBody
	fmt.Println(txt)
	//singlequestion := models.SingleQuestion()
	//c.Data["json"] = singlequestion
	//c.ServeJSON()
}