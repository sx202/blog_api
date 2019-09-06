package controllers

import "github.com/astaxie/beego"

type INDEX struct {
	beego.Controller
}

func (c *INDEX) Index()  {
	c.Ctx.WriteString("welcome blog_api !!")
}