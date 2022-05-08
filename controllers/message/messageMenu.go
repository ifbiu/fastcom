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

type MessageMenuController struct {
	controllers.BaseController
}

func (this *MessageMenuController) Get()  {
	openId := this.GetString("openid")
	pageStr := this.GetString("page")
	pageSizeStr := this.GetString("pageSize")
	if openId == "" || pageStr == "" || pageSizeStr == "" {
		msg := ""
		if openId == "" {
			msg += "openid "
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
	menu, err := message.GetMessageMenu(openId,page,pageSize)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "查询失败！"}
		this.ServeJSON()
		return
	}
	this.Data["json"] = utils.ResultDataUtil{Code: 200,Data: menu}
	this.ServeJSON()
	return

}
