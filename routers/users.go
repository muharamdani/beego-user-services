package routers

import (
	"github.com/astaxie/beego"
	"user_services/controllers"
)

func init() {
	beego.Router(
		"/users", &controllers.UsersController{},
		"get:List")
	beego.Router(
		"/users/:id", &controllers.UsersController{},
		"get:GetByID")
	beego.Router(
		"/users", &controllers.UsersController{},
		"post:PostCreate")
	beego.Router(
		"/users/:id", &controllers.UsersController{},
		"patch:PostUpdate")
	beego.Router(
		"/users/:id", &controllers.UsersController{},
		"delete:PostDelete")
}
