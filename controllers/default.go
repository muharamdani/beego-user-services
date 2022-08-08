package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Redirect("/users", http.StatusFound)
}
