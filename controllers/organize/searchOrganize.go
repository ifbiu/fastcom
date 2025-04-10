package organize

import (
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/db"
	"fastcom/logic/organize"
	"fastcom/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
	"time"
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
	isSearchStr := this.GetString("isSearch")
	if openId == "" || uuidStr == "" || isSearchStr == "" {
		msg := ""
		if openId == "" {
			msg += "openid "
		}
		if uuidStr == "" {
			msg += "uuid "
		}
		if isSearchStr == "" {
			msg += "isSearch "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少字段："+msg}
		this.ServeJSON()
		return
	}
	if isSearchStr!="1" && isSearchStr!="2"{
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "isSearch 参数有误！"}
		this.ServeJSON()
		return
	}
	// 走搜索 存搜索记录
	if isSearchStr=="1" {
		key := "searchOrganize:"+openId
		rds,err := db.InitRedis()
		defer rds.Close()
		if err != nil {
			log.Panicln(err)
		}

		_,err = rds.ZRem(key,uuidStr)
		if err != nil {
			result := utils.ResultUtil{
				Code: 500,
				Msg: "服务异常！",
			}
			this.Data["json"] = &result
			this.ServeJSON()
			return
		}
		_, err = rds.ZIncrBy(key, float64(time.Now().Unix()),uuidStr)
		if err != nil {
			result := utils.ResultUtil{
				Code: 500,
				Msg: "服务异常！",
			}
			this.Data["json"] = &result
			this.ServeJSON()
			return
		}
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

	searchOrganize, err := organize.SearchOrganize(uuid)
	if err != nil {
		fmt.Println(err)
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