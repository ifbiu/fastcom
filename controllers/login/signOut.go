package login

import (
	"fastcom/controllers"
	"fastcom/db"
	"fastcom/utils"
	"log"
)

type SignOutController struct {
	controllers.BaseController
}

func (this *SignOutController) Get() {
	openId := this.GetString("openid")
	if openId == "" {
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少参数：openid"}
		this.ServeJSON()
		return
	}
	key := "login:"+openId
	rds, err := db.InitRedis()
	if err != nil {
		log.Fatalln(err)
		return
	}
	exists, err := rds.Exists(key)
	if exists {
		del, err := rds.Del(key)
		if err != nil {
			return
		}
		if del {
			this.Data["json"] = utils.ResultUtil{Code: 200,Msg: "已退出登录"}
			this.ServeJSON()
			return
		}else{
			this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "退出失败"}
			this.ServeJSON()
			return
		}
	}else{
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "当前用户还没有登录"}
		this.ServeJSON()
		return
	}
}