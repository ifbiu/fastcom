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

type UpdateHeadPortraitController struct {
	controllers.BaseController
}

type RequestUpdateHeadPortrait struct {
	Openid string `json:"openid"`
	Image string `json:"image"`
}

func (this *UpdateHeadPortraitController) Post()  {
	personalParam := &RequestUpdateHeadPortrait{}
	json.Unmarshal(this.Ctx.Input.RequestBody, personalParam)
	if personalParam.Openid=="" ||
		personalParam.Image=="" {
		var (
			isOpenid = ""
			isImage = ""
		)
		if personalParam.Openid=="" {
			isOpenid = "openid "
		}
		if personalParam.Image=="" {
			isImage = "image "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenid+isImage,
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

	err = personal.UpdateHeadPortrait(personalParam.Openid, personalParam.Image)
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