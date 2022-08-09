package utils

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

const (
	SUCCESS = true
	FAIL    = false
	UNKNOWN = false
)

type ResResult struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) ParseBody(v interface{}) error {
	return json.Unmarshal(this.Ctx.Input.RequestBody, v)
}

func (this *BaseController) ResParseError(err error) {
	this.Error(FAIL, "The requested data is not in the correct format！", err)
}

func (this *BaseController) ResJson(v interface{}) {
	this.Data["json"] = v
	this.ServeJSON()
}

func (this *BaseController) Success(data interface{}, msg string) {
	result := ResResult{
		Status:  SUCCESS,
		Data:    data,
		Message: msg,
	}
	this.ResJson(result)
}

func (this *BaseController) UnprocessableEntity(data map[string]string) {
	result := ResResult{
		Status:  FAIL,
		Message: "Unprocessable Entity",
		Data:    data,
	}
	this.Ctx.Output.Status = http.StatusUnprocessableEntity
	this.ResJson(result)
}

func (this *BaseController) Error(status bool, msg string, err error) {
	result := ResResult{
		Status:  status,
		Message: msg,
	}
	beego.Error(fmt.Sprintf("%s:%v", msg, err))
	this.ResJson(result)
}
