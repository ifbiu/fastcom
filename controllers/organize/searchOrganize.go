package organize

import (
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/logic/organize"
	"fastcom/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
)

type SearchOrganizeController struct {
	controllers.BaseController
}

type RequestSearchOrganize struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

func (this *SearchOrganizeController) Get()  {
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

	searchOrganize, err := organize.SearchOrganize(uuid)
	if err != nil {
		if err == orm.ErrNoRows{
			result := RequestSearchOrganize{
				Code: 200,
				Data: new(interface{}),
			}
			this.Data["json"] = &result
			this.ServeJSON()
			return
		}
		result := utils.ResultUtil{
			Code: 500,
			Msg: "服务异常！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	result := RequestSearchOrganize{
		Code: 200,
		Data: searchOrganize,
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}