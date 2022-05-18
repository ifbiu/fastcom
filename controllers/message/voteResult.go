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

type VoteResultController struct {
	controllers.BaseController
}

func (this *VoteResultController) Get()  {
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
		this.Data["json"] = utils.ResultUtil{Code: 500, Msg: "缺少字段：" + msg}
		this.ServeJSON()
		return
	}
	typeId, err := strconv.Atoi(typeIdStr)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg:  "typeId格式错误！",
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
	if (!auth) {
		this.Redirect("/noAuth", 302)
		return
	}
	_, err = message.IsAuthVote(openId, typeId)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg:  "您不属于该组织！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	voteResult, err := message.VoteResult(typeId)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg:  "获取数据失败！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	result := utils.ResultDataUtil{
		Code: 500,
		Data:  voteResult,
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return

}
