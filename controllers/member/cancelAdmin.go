package member

import (
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/logic/member"
	"fastcom/logic/organize"
	"fastcom/utils"
	"fmt"
	"log"
	"strconv"
)

type CancelAdminController struct {
	controllers.BaseController
}

func (this *CancelAdminController) Get()  {
	openId := this.GetString("openid")
	setOpenId := this.GetString("setOpenid")
	uuidStr := this.GetString("uuid")
	if openId == "" || uuidStr == "" || setOpenId== "" {
		msg := ""
		if openId == "" {
			msg += "openid "
		}
		if uuidStr == "" {
			msg += "uuid "
		}
		if setOpenId == "" {
			msg += "setOpenid "
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
	authOrganize1, err := organize.GetAuthOrganize(openId, uuid)
	if err != nil {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "组织与成员无联系！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	if authOrganize1 != 1 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉，您没有设置权限！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	authOrganize2, err := organize.GetAuthOrganize(setOpenId, uuid)
	if err != nil {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "被设置成员不在该组织！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	if authOrganize2 == 1 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "你不可以给创建者降权！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	if authOrganize2 != 2 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉，该成员没有管理权限，无需降权！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	err = member.CancelAdmin(uuid, setOpenId)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "设置管理员失败！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	result := utils.ResultUtil{
		Code: 200,
		Msg: "降权成功",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}