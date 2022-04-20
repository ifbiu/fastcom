package organize

import (
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/logic/organize"
	"fastcom/utils"
	"log"
)

type OrganizeMenuController struct {
	controllers.BaseController
}

func (this *OrganizeMenuController) Get()  {
	openId := this.GetString("openid")
	status := this.GetString("status")
	if openId == "" || status == "" {
		msg := ""
		if openId == "" {
			msg += "openid "
		}
		if status == "" {
			msg += "status "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少字段："+msg}
		this.ServeJSON()
		return
	}

	if status != "admin" && status != "member" {
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "status参数传入无效"}
		this.ServeJSON()
		return
	}

	auth, err := common.CheckAuth(openId)
	if err != nil {
		log.Panicln(err)
		return
	}
	if(!auth){
		this.Redirect("/noAuth",302)
		return
	}

	menus, err := organize.GetMenu(openId,status)
	if err != nil {
		log.Panicln(err)
		return 
	}

	this.Data["json"] = utils.ResultDataUtil{
		Code: 200,
		Data: menus,
	}
	this.ServeJSON()
	return
}
