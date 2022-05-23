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

type HistoryInfoController struct {
	controllers.BaseController
}

func (this *HistoryInfoController) Get()  {
	openId := this.GetString("openid")
	uuidStr := this.GetString("uuid")
	typeStr := this.GetString("type")
	typeIdStr := this.GetString("typeId")
	if openId == "" || uuidStr == "" || typeStr == "" || typeIdStr == "" {
		msg := ""
		if openId == "" {
			msg += "openid "
		}
		if uuidStr == "" {
			msg += "uuid "
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
			Msg: "type格式错误",
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
			Msg: "type格式错误",
		}
		this.Data["json"] = &result
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

	if authOrganize != 1 && authOrganize != 2 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉您没有权限！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	info, err := organize.GetHistoryInfo(theType, typeId)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "获取详情失败！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	result := utils.ResultDataUtil{
		Code: 200,
		Data: info,
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}