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

type UpdateSexController struct {
	controllers.BaseController
}

type RequestUpdateSex struct {
	Openid string `json:"openid"`
	Sex int `json:"sex"`
}

func (this *UpdateSexController) Post()  {
	personalParam := &RequestUpdateSex{}
	json.Unmarshal(this.Ctx.Input.RequestBody, personalParam)
	if personalParam.Openid=="" ||
		personalParam.Sex==0 {
		var (
			isOpenid = ""
			isSex = ""
		)
		if personalParam.Openid=="" {
			isOpenid = "openid "
		}
		if personalParam.Sex==0 {
			isSex = "sex "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenid+isSex,
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

	err = personal.UpdateSex(personalParam.Openid, personalParam.Sex)
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
