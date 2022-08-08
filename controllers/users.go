package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/core/validation"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"user_services/models"
)

type UsersController struct {
	BaseController
	users models.Users
}

func (this *UsersController) List() {
	// Pagination | Temporary use case, still explore for best practices
	perPage, _ := this.GetInt("per_page")
	page, _ := this.GetInt("page")
	if page >= 1 {
		page -= 1
	}
	if perPage == 0 {
		perPage = 10
	}
	if perPage > 100 {
		perPage = 100
		this.Ctx.WriteString("Maximum data per page is 100 data!")
	}
	page *= perPage

	// Get data by limit and offset
	result, err := this.users.List(perPage, page)
	if err != nil {
		this.Error(UNKNOWN, "An exception occurred when querying the users list", err)
		return
	}
	this.Success(result, "Get data successful")
}

func (this *UsersController) GetByID() {
	result, err := this.users.GetByID(bson.ObjectIdHex(this.GetString(":id")))
	if err != nil {
		this.Error(UNKNOWN, "Query users data exception!", err)
		return
	}
	this.Success(result, "Get data successful")
}

func (this *UsersController) PostCreate() {
	if err := this.ParseBody(&this.users); err != nil {
		this.ResParseError(err)
		return
	}
	// Validation
	valid := validation.Validation{}
	b, err := valid.Valid(&this.users)
	if err != nil {
		this.ResParseError(err)
		return
	}
	if !b {
		var msg string
		for _, err := range valid.Errors {
			msg += fmt.Sprintf("%s: %s\r", err.Key, err.Message)
		}
		this.Ctx.ResponseWriter.WriteHeader(400)
		this.Data["json"] = &map[string]interface{}{"status": "error", "message": msg}
		fmt.Println(msg)
		return
	}
	result, err := this.users.Create()
	if err != nil {
		beego.Error(err)
		return
	}
	this.Success(result, "Create success")
}

func (this *UsersController) PostUpdate() {
	if err := this.ParseForm(&this.users); err != nil {
		this.ResParseError(err)
		return
	}
	err := this.users.UpdateByID(bson.ObjectIdHex(this.GetString(":id")))
	if err != nil {
		beego.Error(err)
		return
	}
	this.Redirect("/users", http.StatusFound)
}

func (this *UsersController) PostDelete() {
	err := this.users.DeleteByID(bson.ObjectIdHex(this.GetString(":id")))
	if err != nil {
		beego.Error(err)
		return
	}
	this.Redirect("/users", http.StatusFound)
}
