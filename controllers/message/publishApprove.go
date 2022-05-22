package message

import (
	"encoding/json"
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/logic/message"
	"fastcom/logic/organize"
	"fastcom/utils"
	"fmt"
	"log"
	"strconv"
)


type PublishApproveController struct {
	controllers.BaseController
}

type RequestPublishApprove struct {
	Openid string `json:"openid"`
	Uuid string `json:"uuid"`
}

func (this *PublishApproveController) Post()  {
	publishParam := &RequestPublishApprove{}
	json.Unmarshal(this.Ctx.Input.RequestBody, publishParam)
	if publishParam.Openid=="" ||
		publishParam.Uuid==""  {
		var (
			isOpenid = ""
			isUuid = ""
		)
		if publishParam.Openid=="" {
			isOpenid = "openid "
		}
		if publishParam.Uuid=="" {
			isUuid = "uuid "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenid+isUuid}
		this.ServeJSON()
		return
	}

	auth, err := common.CheckAuth(publishParam.Openid)
	if err != nil {
		log.Panicln(err)
		return
	}
	if(!auth){
		this.Redirect("/noAuth",302)
		return
	}

	uuid, err := strconv.Atoi(publishParam.Uuid)
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
	// 是否没有管理员权限
	authOrganize, _ := organize.GetAuthOrganize(publishParam.Openid, uuid)

	if authOrganize != nil{
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉，您已经在该组织中！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	selectApprove, err := message.SelectApprove(publishParam.Openid, publishParam.Uuid)
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
	if selectApprove {
		result := utils.ResultUtil{
			Code: 200,
			Msg: "审核中...",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	openids, err := message.SelectApproveOpenIds(publishParam.Uuid)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉，搜索不到成员！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	_, err = message.PublishApprove(publishParam.Openid,openids,publishParam.Uuid)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "发布审核消息失败！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	err = common.AmqpMessage(openids)
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
		Msg: "申请成功",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}