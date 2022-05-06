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

type IsMessageReadController struct {
	controllers.BaseController
}

func (this *IsMessageReadController) Get() {
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
		this.Data["json"] = utils.ResultUtil{Code: 500, Msg: "缺少字段：" + msg}
		this.ServeJSON()
		return
	}
	theType, err := strconv.Atoi(typeStr)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg:  "type格式错误！",
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
			Msg:  "type格式错误！",
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
	if (!auth) {
		this.Redirect("/noAuth", 302)
		return
	}

	read, err := message.IsMessageRead(theType, typeId, openId)
	if err != nil || !read {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg:  "服务异常！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	result := utils.ResultUtil{
		Code: 200,
		Msg:  "Is Read",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}