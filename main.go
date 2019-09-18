package main

import (
	_ "github.com/sx202/blog_api/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

//这个程序采用前后端分离的架构，前端所需要的所有数据都是通过API接口提供
//如果数据库采用的是sqlite3，只能做单用户系统，就是同一时间只能一个用户进行操作。


	beego.Run()
}
