package organize

import (
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/db"
	"fastcom/utils"
	"log"
)

type HistoryRecordController struct {
	controllers.BaseController
}

func (this *HistoryRecordController) Get()  {
	openId := this.GetString("openid")
	if openId == ""{
		msg := ""
		if openId == "" {
			msg += "openid "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少字段："+msg}
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
	key := "searchOrganize:"+openId
	rds,err := db.InitRedis()
	defer rds.Close()
	if err != nil {
		log.Panicln(err)
	}
	resArr, err := rds.LRange(key,0,8)
	if err != nil {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "查询失败",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	result := utils.ResultDataUtil{
		Code: 200,
		Data: resArr,
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return

}



