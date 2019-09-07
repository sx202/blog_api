package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type Test struct {
	beego.Controller
}

func (c *Test) TEST ()  {
	txt := c.Ctx.Input.RequestBody
	fmt.Println(string(txt))
}