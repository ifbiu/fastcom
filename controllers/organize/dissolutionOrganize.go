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

type DissolutionOrganizeController struct {
	controllers.BaseController
}

func (this *DissolutionOrganizeController) Get()  {
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
	if (authOrganize != 1) {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "解散失败，权限不够！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	outOrganize, err := organize.DissolutionOrganize(uuid)
	if err != nil || !outOrganize{
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "解散失败！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	var openIds []string
	openIds, err = organize.GetOrganizeOpenIds(uuid)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "获取消息推送人失败！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	err = common.AmqpMessage(openIds)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "消息推送失败！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	result := utils.ResultUtil{
		Code: 200,
		Msg: "解散成功",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}
