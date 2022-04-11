package login

import (
	"fastcom/utils"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}


func (ctx *LoginController) Get()  {
	result := utils.ResultUtil{
		Code: 200,
		Msg: "登录成功",
	}
	ctx.Data["json"] = &result
	ctx.ServeJSON()
}