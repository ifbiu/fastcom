package openId

import (
	"fastcom/controllers"
	"fastcom/utils"
	"github.com/astaxie/beego/httplib"
	"log"
)

type GetOpenIdController struct {
	controllers.BaseController
}

type GetJsCode struct {
	Openid string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid string `json:"unionid"`
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

type GetJsCodeSuccess struct {
	Openid string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid string `json:"unionid"`
}

type GetJsCodeErr struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

func (this *GetOpenIdController) Get(){
	jscode := this.GetString("jscode")
	if jscode == "" {
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少jscode字段"}
		this.ServeJSON()
		return
	}
	APPID := "wx729d7fc79ac8c86f"
	SECRET := "48231588a3adeb2adeb07cf70e4db02a"
	getJsCode := GetJsCode{}
	err :=httplib.Get("https://api.weixin.qq.com/sns/jscode2session?appid="+APPID+"&secret="+SECRET+"&js_code="+jscode+"&grant_type=authorization_code").ToJSON(&getJsCode)
	if err!=nil {
		log.Panicln(err)
	}
	if getJsCode.Openid != "" {
		getJsCodeSuccess := GetJsCodeSuccess{
			Openid: getJsCode.Openid,
			SessionKey: getJsCode.SessionKey,
			Unionid: getJsCode.Unionid,
		}
		this.Data["json"] = getJsCodeSuccess
	}else{
		getJsCodeErr := GetJsCodeErr{
			ErrCode: getJsCode.ErrCode,
			ErrMsg: getJsCode.ErrMsg,
		}
		this.Data["json"] = getJsCodeErr
	}
	this.ServeJSON()
	return
}