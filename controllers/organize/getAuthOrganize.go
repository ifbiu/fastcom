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

type GetAuthOrganize struct {
	controllers.BaseController
}

func (this *GetAuthOrganize) Get()  {
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

	organizeDel, err := organize.IsOrganizeDel(uuid)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "找不到该组织！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	if organizeDel == 2 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "该组织已解散！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	authOrganize, err := organize.GetAuthOrganize(openId,uuid)
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
	result := utils.ResultDataUtil{
		Code: 200,
		Data: authOrganize,
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}