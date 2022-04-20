package login

import (
	"fastcom/controllers"
	"fastcom/utils"
)


type NoAuthController struct {
	controllers.BaseController
}

func (this * NoAuthController) Get()  {
	this.Data["json"] = utils.ResultUtil{Code: 401,Msg: "登录已过期，请重新登录"}
	this.ServeJSON()
	return
}
