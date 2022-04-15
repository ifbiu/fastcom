package controllers

import (
	"fastcom/utils"
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404() {
	this.Data["json"] = utils.ResultUtil{Code: 404,Msg: "未定义接口"}
	this.ServeJSON()
	return
}
