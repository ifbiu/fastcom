package message

import (
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/logic/message"
	"fastcom/utils"
	"fmt"
	"log"
)

type AngleMarkController struct {
	controllers.BaseController
}

func (this *AngleMarkController) Get()  {
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
	angleMark, err := message.GetAngleMark(openId)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "查询角标失败"}
		this.ServeJSON()
		return
	}
	this.Data["json"] = utils.ResultDataUtil{Code: 200,Data: angleMark}
	this.ServeJSON()
	return
}