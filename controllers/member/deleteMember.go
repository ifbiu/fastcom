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

type DeleteMemberController struct {
	controllers.BaseController
}

func (this *DeleteMemberController) Get()  {
	openId := this.GetString("openid")
	delOpenId := this.GetString("delOpenid")
	uuidStr := this.GetString("uuid")
	if openId == "" || uuidStr == "" || delOpenId== "" {
		msg := ""
		if openId == "" {
			msg += "openid "
		}
		if uuidStr == "" {
			msg += "uuid "
		}
		if delOpenId == "" {
			msg += "delOpenid "
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
	if authOrganize1 == 3 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉，您没有操作权限！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	// 判断被删除者权限
	authOrganize2, err := organize.GetAuthOrganize(delOpenId, uuid)
	if err != nil {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "被删除成员不在该组织！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	if authOrganize2 == 1 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉，组织不能删除超级管理员！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	if authOrganize2 == 2 && authOrganize1 != 1 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉，您没有删除管理员的权限！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	deleteMember, err := member.DeleteMember(uuid,delOpenId)
	if err != nil || !deleteMember{
		fmt.Println(err)
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
		Msg: "删除成员成功",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}