package register

import (
	"fastcom/logic/register"
	"fastcom/utils"
	"github.com/astaxie/beego"
	"log"
)

type IsRegisterUserController struct {
	beego.Controller
}

func (this *IsRegisterUserController) Get()  {
	openid := this.GetString("openid")
	if openid == "" {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "缺少参数： openid",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	isExist,err := register.IsRegisterUser(openid)
	if err != nil {
		log.Panicln(err)
	}
	if !isExist {
		this.Data["json"] = utils.ResultUtil{
			Code: 500,
			Msg: "用户已存在",
		}
		this.ServeJSON()
		return
	}
	this.Data["json"] = utils.ResultUtil{
		Code: 200,
		Msg: "可以注册",
	}
	this.ServeJSON()
	return
}