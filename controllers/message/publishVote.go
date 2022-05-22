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


type PublishVoteController struct {
	controllers.BaseController
}

type RequestPublishVote struct {
	Openid string `json:"openid"`
	Uuid string `json:"uuid"`
	Title string `json:"title"`
	MaxNum int `json:"maxNum"`
	IsAbstained int `json:"isAbstained"`
	EndTime int64 `json:"endTime"`
	Items []string `json:"items"`
	Members []string `json:"members"`
}

func (this *PublishVoteController) Post()  {
	publishParam := &RequestPublishVote{}
	json.Unmarshal(this.Ctx.Input.RequestBody, publishParam)
	if publishParam.Openid=="" ||
		publishParam.Uuid=="" ||
		publishParam.Title=="" ||
		publishParam.MaxNum == 0||
		publishParam.IsAbstained == 0||
		publishParam.EndTime == 0||
		len(publishParam.Items)==0 ||
		len(publishParam.Members)==0 {
		var (
			isOpenid = ""
			isUuid = ""
			isTitle = ""
			isMaxNum = ""
			isIsAbstained = ""
			isEndTime = ""
			isItems = ""
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
		if publishParam.IsAbstained==0 {
			isIsAbstained = "isAbstained "
		}
		if publishParam.EndTime==0 {
			isEndTime = "endTime "
		}
		if publishParam.MaxNum==0 {
			isMaxNum = "maxNum "
		}
		if len(publishParam.Items)==0 {
			isItems = "items "
		}
		if len(publishParam.Members)==0 {
			isMembers = "members "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenid+isUuid+isTitle+isMaxNum+isIsAbstained+isEndTime+isItems+isMembers}
		this.ServeJSON()
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


	_, err = message.PublishVote(publishParam.Openid,publishParam.Members,publishParam.Uuid,publishParam.Title,publishParam.MaxNum,publishParam.IsAbstained,publishParam.EndTime,publishParam.Items)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "投票发布失败！",
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
		Msg: "发布投票成功！",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}