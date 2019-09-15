// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego/plugins/cors"
	"github.com/sx202/blog_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.InsertFilter("*",beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		//AllowOrigins:     nil,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"token", "key", "Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		//MaxAge:           0,
	}))

	beego.Router("/", &controllers.INDEX{},"*:Index")

	beego.Router("/user", &controllers.UserController{}, "GET:GetAll")
	beego.Router("/bloguser", &controllers.UserController{}, "GET:Bloguser")

	beego.Router("/getquestionid", &controllers.System{}, "GET:GetQuestionId")

	beego.Router("/getquestion", &controllers.System{}, "GET:GetQuestion")
	beego.Router("/getquestionall", &controllers.System{}, "GET:GetAllQuestion")
	beego.Router("/insertquestion", &controllers.System{}, "POST:InsertQuestion")
	beego.Router("/updatequestion", &controllers.System{}, "POST:UpdateQuestion")
	beego.Router("/deletequestion", &controllers.System{}, "POST:DeleteQuestion")



}
