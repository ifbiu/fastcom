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

type GetMemberInfoController struct {
	controllers.BaseController
}

func (this *GetMemberInfoController) Get() {
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
			Msg: "组织与成员无联系！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	memberInfo,err := member.GetMemberInfo(authOrganize,uuid)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = utils.ResultUtil{
			Code: 500,
			Msg: "查询服务异常！",
		}
		this.ServeJSON()
		return
	}

	this.Data["json"] = utils.ResultDataUtil{
		Code: 200,
		Data: memberInfo,
	}
	this.ServeJSON()
	return
}