package openId

import (
	"fastcom/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"log"
)

type GetOpenIdController struct {
	beego.Controller
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

func (ctx *GetOpenIdController) Get(){
	jscode := ctx.GetString("jscode")
	if jscode == "" {
		ctx.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少jscode字段"}
		ctx.ServeJSON()
	}
	APPID := "wx729d7fc79ac8c86f"
	SECRET := "48231588a3adeb2adeb07cf70e4db02a"
	JSCODE := jscode
	getJsCode := GetJsCode{}
	err :=httplib.Get("https://api.weixin.qq.com/sns/jscode2session?appid="+APPID+"&secret="+SECRET+"&js_code="+JSCODE+"&grant_type=authorization_code").ToJSON(&getJsCode)
	if err!=nil {
		log.Panicln(err)
	}
	if getJsCode.Openid != "" {
		getJsCodeSuccess := GetJsCodeSuccess{
			Openid: getJsCode.Openid,
			SessionKey: getJsCode.SessionKey,
			Unionid: getJsCode.Unionid,
		}
		ctx.Data["json"] = getJsCodeSuccess
	}else{
		getJsCodeErr := GetJsCodeErr{
			ErrCode: getJsCode.ErrCode,
			ErrMsg: getJsCode.ErrMsg,
		}
		ctx.Data["json"] = getJsCodeErr
	}
	ctx.ServeJSON()
}