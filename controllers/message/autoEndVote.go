package message

import (
	"fastcom/controllers"
	"fastcom/logic/message"
	"fastcom/utils"
	"fmt"
	"strconv"
)

type AutoEndVoteController struct {
	controllers.BaseController
}

func (this *AutoEndVoteController) Get()  {
	typeIdStr := this.GetString("typeId")
	if typeIdStr == "" {
		msg := ""
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
			Msg: "typeId格式错误！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	err = message.AutoEndVote(typeId)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "自动截止投票失败！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	result := utils.ResultUtil{
		Code: 200,
		Msg: "自动截止投票成功",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}