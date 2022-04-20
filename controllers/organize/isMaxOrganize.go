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

type IsMaxOrganizeController struct {
	controllers.BaseController
}

type MaxOrganize struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	State int `json:"state"`
}

func (this *IsMaxOrganizeController) Get()  {
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

	isHave,resBool, err := organize.IsMaxOrganize(uuid,openId)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "服务器异常！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	if !isHave {
		result := MaxOrganize{
			Code: 200,
			Msg: "用户已加入该组织，不可重复加入！",
			State: 3,
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	if !resBool {
		result := MaxOrganize{
			Code: 200,
			Msg: "该组织已满！",
			State: 2,
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	result := MaxOrganize{
		Code: 200,
		Msg: "可以加入该组织",
		State: 1,
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}