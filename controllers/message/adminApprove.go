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

type AdminApproveController struct {
	controllers.BaseController
}

func (this *AdminApproveController) Get()  {
	openId := this.GetString("openid")
	approveStr := this.GetString("approve")
	typeIdStr := this.GetString("typeId")
	if openId == "" || approveStr == "" || typeIdStr == "" {
		msg := ""
		if openId == "" {
			msg += "openid "
		}
		if approveStr == "" {
			msg += "approve "
		}
		if typeIdStr == "" {
			msg += "typeId "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少字段："+msg}
		this.ServeJSON()
		return
	}
	approve, err := strconv.Atoi(approveStr)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "approve格式错误！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	if approve != 1 && approve != 2 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉您没有权限！",
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
			Msg: "typeId格式错误！",
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

	approveAuth, err := message.CheckApproveAuth(openId, typeId)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "当前用户不属于该组织！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	if approveAuth==3 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "当前用户权限不够！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}else if approveAuth==1 || approveAuth==2 {
		err := message.AdminApprove(openId, typeId, approve)
		if err != nil {
			fmt.Println(err)
			result := utils.ResultUtil{
				Code: 500,
				Msg: "审核异常！",
			}
			this.Data["json"] = &result
			this.ServeJSON()
			return
		}
		result := utils.ResultUtil{
			Code: 200,
			Msg: "审核完成",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
}