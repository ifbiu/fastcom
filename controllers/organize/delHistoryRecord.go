package organize

import (
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/db"
	"fastcom/utils"
	"log"
)

type DelHistoryRecordController struct {
	controllers.BaseController
}

func (this *DelHistoryRecordController) Get()  {
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
	boolRes, err := rds.Del(key)
	if err != nil || !boolRes {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "删除失败",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	result := utils.ResultUtil{
		Code: 200,
		Msg: "删除成功",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return

}