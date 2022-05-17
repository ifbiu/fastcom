package message

import (
	"encoding/json"
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/logic/message"
	"fastcom/utils"
	"fmt"
	"log"
	"strconv"
)

type VoteOperationController struct {
	controllers.BaseController
}

type RequestVoteOperation struct {
	OpenId string `json:"openid"`
	Vote string `json:"vote"`
	TypeId string `json:"typeId"`
	SerialIds []int `json:"serialIds"`
}

func (this *VoteOperationController) Post()  {
	publishParam := &RequestVoteOperation{}
	json.Unmarshal(this.Ctx.Input.RequestBody, publishParam)
	if publishParam.OpenId=="" ||
		publishParam.Vote=="" ||
		publishParam.TypeId=="" ||
		(len(publishParam.SerialIds)==0 && publishParam.Vote == "1")  {
		var (
			isOpenId = ""
			isVote = ""
			isTypeId = ""
			isSerialIds = ""
		)
		if publishParam.OpenId=="" {
			isOpenId = "openid "
		}
		if publishParam.Vote=="" {
			isVote = "vote "
		}
		if publishParam.TypeId=="" {
			isTypeId = "typeId "
		}
		if len(publishParam.SerialIds)==0 && publishParam.Vote == "1"{
			isSerialIds = "serialIds "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenId+isVote+isTypeId+isSerialIds}
		this.ServeJSON()
		return
	}
	vote, err := strconv.Atoi(publishParam.Vote)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "vote格式错误！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	typeId, err := strconv.Atoi(publishParam.TypeId)
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
	if vote==0 {
		isVote := message.IsVote(publishParam.OpenId, typeId)
		if isVote {
			// 投过票了
			result := utils.ResultDataUtil{
				Code: 200,
				Data: true,
			}
			this.Data["json"] = &result
			this.ServeJSON()
			return
		}else{
			// 没投过票
			result := utils.ResultDataUtil{
				Code: 200,
				Data: false,
			}
			this.Data["json"] = &result
			this.ServeJSON()
			return
		}
	}
	if vote != 1 && vote != 2 {	// 1、投票 2、弃票
		result := utils.ResultUtil{
			Code: 500,
			Msg: "vote输入有误！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	auth, err := common.CheckAuth(publishParam.OpenId)
	if err != nil {
		log.Panicln(err)
		return
	}
	if(!auth){
		this.Redirect("/noAuth",302)
		return
	}

	err = message.VoteOperation(publishParam.OpenId, vote, typeId,publishParam.SerialIds)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "投票失败！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	if vote==1 {
		result := utils.ResultUtil{
			Code: 200,
			Msg: "投票成功",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}else if vote==2 {
		result := utils.ResultUtil{
			Code: 200,
			Msg: "弃票成功",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}






}