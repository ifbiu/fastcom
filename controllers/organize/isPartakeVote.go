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

type IsPartakeVoteController struct {
	controllers.BaseController
}

func (this *IsPartakeVoteController) Get()  {
	openId := this.GetString("openid")
	typeIdStr := this.GetString("typeId")
	if openId == "" || typeIdStr == "" {
		msg := ""
		if openId == "" {
			msg += "openid "
		}
		if typeIdStr == "" {
			msg += "typeId "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少字段："+msg}
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

	auth, err := common.CheckAuth(openId)
	if err != nil {
		log.Panicln(err)
		return
	}
	if(!auth){
		this.Redirect("/noAuth",302)
		return
	}

	authOrganize, err := organize.GetAuthVote(openId,typeId)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "找不到此用户关联组织",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	if authOrganize!=1 && authOrganize!=2 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉您无权限！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	info, err := organize.IsPartakeVote(typeId,openId)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "服务异常！",
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
