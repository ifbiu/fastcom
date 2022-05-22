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


type PublishMessageController struct {
	controllers.BaseController
}

type RequestPublishMessage struct {
	Openid string `json:"openid"`
	Uuid string `json:"uuid"`
	Title string `json:"title"`
	Content string `json:"content"`
	Members []string `json:"members"`
}

func (this *PublishMessageController) Post()  {
	publishParam := &RequestPublishMessage{}
	json.Unmarshal(this.Ctx.Input.RequestBody, publishParam)
	if publishParam.Openid=="" ||
		publishParam.Uuid=="" ||
		publishParam.Title=="" ||
		publishParam.Content == ""||
		len(publishParam.Members)==0 {
		var (
			isOpenid = ""
			isUuid = ""
			isTitle = ""
			isContent = ""
			isMembers = ""
		)
		if publishParam.Openid=="" {
			isOpenid = "openid "
		}
		if publishParam.Uuid=="" {
			isUuid = "uuid "
		}
		if publishParam.Title=="" {
			isTitle = "title "
		}
		if publishParam.Content=="" {
			isContent = "content "
		}
		if len(publishParam.Members)==0 {
			isMembers = "members "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenid+isUuid+isTitle+isContent+isMembers}
		this.ServeJSON()
		return
	}
	log.Println(publishParam)
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

	auth, err := common.CheckAuth(publishParam.Openid)
	if err != nil {
		log.Panicln(err)
		return
	}
	if(!auth){
		this.Redirect("/noAuth",302)
		return
	}

	// 是否没有管理员权限
	authOrganize, err := organize.GetAuthOrganize(publishParam.Openid, uuid)
	if err != nil {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "组织与成员无联系！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	if authOrganize == 3 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉，您没有发布的权限！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	//openids, err := message.SelectOpenIds(publishParam.Uuid)
	//if err != nil {
	//	fmt.Println(err)
	//	result := utils.ResultUtil{
	//		Code: 500,
	//		Msg: "抱歉，搜索不到成员！",
	//	}
	//	this.Data["json"] = &result
	//	this.ServeJSON()
	//	return
	//}

	_, err = message.PublishMessage(publishParam.Openid,publishParam.Members,publishParam.Uuid,publishParam.Title,publishParam.Content)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "公告发布失败！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	err = common.AmqpMessage(publishParam.Members)
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
		Msg: "发布公告成功！",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}