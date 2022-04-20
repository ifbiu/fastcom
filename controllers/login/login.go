package login

import (
	"encoding/json"
	"fastcom/controllers"
	"fastcom/db"
	"fastcom/utils"
	"log"
	"time"
)

type LoginController struct {
	controllers.BaseController
}

type RequestLogin struct {
	OpenId string `json:"openid"`
}

func (this *LoginController) Post()  {
	loginParam := &RequestLogin{}
	json.Unmarshal(this.Ctx.Input.RequestBody,loginParam)
	if loginParam.OpenId == "" {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "缺少参数： openid",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	key := "login:"+loginParam.OpenId
	rds, err := db.InitRedis()
	if err != nil {
		log.Panicln(err)
	}
	err = rds.Set(key, "login success", time.Hour)
	if err != nil {
		result := utils.ResultUtil{
			Code: 200,
			Msg: "登录失败",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	result := utils.ResultUtil{
		Code: 200,
		Msg: "登录成功",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}