package login

import (
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/logic/organize"
	"fastcom/utils"
	"fmt"
	"log"
)

type GetUserInfoController struct {
	controllers.BaseController
}


func (this *GetUserInfoController) Get() {
	openId := this.GetString("openid")
	if openId == "" {
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少参数：openid"}
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
	info, err := organize.GetUserInfo(openId)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "接口异常！"}
		this.ServeJSON()
		return
	}
	this.Data["json"] = utils.ResultDataUtil{Code: 500,Data: info}
	this.ServeJSON()
	return

}




