package message

import (
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/logic/message"
	"fastcom/utils"
	"fmt"
	"log"
	"strconv"
)

type IsAuthDelController struct {
	controllers.BaseController
}

func (this *IsAuthDelController) Get()  {
	openId := this.GetString("openid")
	typeStr := this.GetString("type")
	typeIdStr := this.GetString("typeId")
	if openId == "" || typeStr == "" || typeIdStr == "" {
		msg := ""
		if openId == "" {
			msg += "openid "
		}
		if typeStr == "" {
			msg += "type "
		}
		if typeIdStr == "" {
			msg += "typeId "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少字段："+msg}
		this.ServeJSON()
		return
	}
	theType, err := strconv.Atoi(typeStr)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "type格式错误！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	typeId, err := strconv.Atoi(typeIdStr)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "type格式错误！",
		}
		this.Data["json"] = &result
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

	res, err := message.IsAuthDel(theType, typeId,openId)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "验证权限失败！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	result := utils.ResultDataUtil{
		Code: 200,
		Data: res,
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return

}