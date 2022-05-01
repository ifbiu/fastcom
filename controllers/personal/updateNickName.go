package personal

import (
	"encoding/json"
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/logic/personal"
	"fastcom/utils"
	"fmt"
	"log"
)

type UpdateNickNameController struct {
	controllers.BaseController
}

type RequestUpdateNickName struct {
	Openid string `json:"openid"`
	NickName string `json:"nickName"`
}

func (this *UpdateNickNameController) Post()  {
	personalParam := &RequestUpdateNickName{}
	json.Unmarshal(this.Ctx.Input.RequestBody, personalParam)
	if personalParam.Openid=="" ||
		personalParam.NickName=="" {
		var (
			isOpenid = ""
			isNickName = ""
		)
		if personalParam.Openid=="" {
			isOpenid = "openid "
		}
		if personalParam.NickName=="" {
			isNickName = "nickName "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenid+isNickName,
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

	err = personal.UpdateNickName(personalParam.Openid, personalParam.NickName)
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
