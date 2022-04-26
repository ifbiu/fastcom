package organize

import (
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/logic/organize"
	"fastcom/utils"
	"fmt"
	"log"
	"strconv"
)

type SignOutOrganizeController struct {
	controllers.BaseController
}

func (this *SignOutOrganizeController) Get()  {
	openId := this.GetString("openid")
	uuidStr := this.GetString("uuid")
	if openId == "" || uuidStr == "" {
		msg := ""
		if openId == "" {
			msg += "openid "
		}
		if uuidStr == "" {
			msg += "uuid "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少字段："+msg}
		this.ServeJSON()
		return
	}
	uuid, err := strconv.Atoi(uuidStr)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "uuid格式错误",
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

	// 是否没有管理员权限
	authOrganize, err := organize.GetAuthOrganize(openId, uuid)
	if err != nil {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "服务异常！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	if (authOrganize == 1) {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "退出失败！超级管理员不可退出组织，请先转移权限！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	outOrganize, err := organize.SignOutOrganize(openId, uuid)
	if err != nil || !outOrganize{
		result := utils.ResultUtil{
			Code: 500,
			Msg: "服务异常！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	result := utils.ResultUtil{
		Code: 200,
		Msg: "退出成功",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}