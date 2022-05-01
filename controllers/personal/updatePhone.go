package personal

import (
	"encoding/json"
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/db"
	"fastcom/logic/personal"
	"fastcom/utils"
	"fmt"
	"log"
)

type UpdatePhoneController struct {
	controllers.BaseController
}

type RequestUpdatePhone struct {
	Openid string `json:"openid"`
	Phone string `json:"newPhone"`
	Code string `json:"code"`
}

func (this *UpdatePhoneController) Post()  {
	personalParam := &RequestUpdatePhone{}
	key := "sms:"+personalParam.Phone
	json.Unmarshal(this.Ctx.Input.RequestBody, personalParam)
	if personalParam.Openid=="" ||
		personalParam.Phone=="" ||
		personalParam.Code==""{
		var (
			isOpenid = ""
			isPhone = ""
			isCode = ""
		)
		if personalParam.Openid=="" {
			isOpenid = "openid "
		}
		if personalParam.Phone=="" {
			isPhone = "phone "
		}
		if personalParam.Code=="" {
			isCode = "code "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenid+isPhone+isCode,
		}
		this.ServeJSON()
		return
	}

	auth, err := common.CheckAuth(personalParam.Openid)
	if err != nil {
		log.Panicln(err)
		return
	}
	if(!auth){
		this.Redirect("/noAuth",302)
		return
	}

	rds,err := db.InitRedis()
	defer rds.Close()
	if err != nil {
		log.Panicln(err)
	}
	smsExists,err := rds.Exists(key)
	if err != nil {
		log.Panicln(err)
	}
	if (!smsExists){
		this.Data["json"] = utils.ResultUtil{
			Code: 500,
			Msg: "验证码已过期，请重新获取",
		}
		this.ServeJSON()
		return
	}
	resCode, err := rds.Get(key)
	if resCode != personalParam.Code {
		this.Data["json"] = utils.ResultUtil{
			Code: 500,
			Msg: "验证码错误",
		}
		this.ServeJSON()
		return
	}

	err = personal.UpdateNickName(personalParam.Openid, personalParam.Phone)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "修改失败！"}
		this.ServeJSON()
		return
	}
	this.Data["json"] = utils.ResultUtil{Code: 200,Msg: "修改成功！"}
	this.ServeJSON()
	return
}
