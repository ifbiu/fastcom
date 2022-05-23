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

type HistoryMessageController struct {
	controllers.BaseController
}

func (this *HistoryMessageController) Get()  {
	openId := this.GetString("openid")
	uuidStr := this.GetString("uuid")
	typeStr := this.GetString("type")
	pageStr := this.GetString("page")
	pageSizeStr := this.GetString("pageSize")
	if openId == "" || uuidStr == "" || typeStr == "" || pageStr == "" || pageSizeStr == "" {
		msg := ""
		if openId == "" {
			msg += "openid "
		}
		if uuidStr == "" {
			msg += "uuid "
		}
		if typeStr == "" {
			msg += "type "
		}
		if pageStr == "" {
			msg += "page "
		}
		if pageSizeStr == "" {
			msg += "pageSize "
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
	theType, err := strconv.Atoi(typeStr)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "type格式错误！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "page格式错误",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "pageSize格式错误",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	if page<0 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "page页数不能小于0！",
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
	if authOrganize!=1 && authOrganize!=2 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉您无权限！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	message, err := organize.HistoryMessage(uuid, theType, page, pageSize)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "查询失败！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	result := utils.ResultDataUtil{
		Code: 200,
		Data: message,
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}