package login

import (
	"fastcom/utils"
	"github.com/astaxie/beego"
)


type NoAuthController struct {
	beego.Controller
}

func (this * NoAuthController) Get()  {
	this.Data["json"] = utils.ResultUtil{Code: 401,Msg: "登录已过期，请重新登录"}
	this.ServeJSON()
	return
}
