package main

import (
	"github.com/sx202/blog_api/models"
	_ "github.com/sx202/blog_api/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	models.BlogAllUser()

	beego.Run()
}
