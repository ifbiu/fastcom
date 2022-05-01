package personal

import (
	"encoding/json"
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/utils"
	"log"
)

type FeedbackController struct {
	controllers.BaseController
}

type RequestFeedback struct {
	Openid string `json:"openid"`
	Title string `json:"title"`
	Content string `json:"content"`
}

func (this *FeedbackController) Post()  {
	feedbackParam := &RequestFeedback{}
	json.Unmarshal(this.Ctx.Input.RequestBody, feedbackParam)
	if feedbackParam.Openid=="" ||
		feedbackParam.Title=="" ||
		feedbackParam.Content==""{
		var (
			isOpenid = ""
			isTitle = ""
			isContent = ""
		)
		if feedbackParam.Openid=="" {
			isOpenid = "openid "
		}
		if feedbackParam.Title=="" {
			isTitle = "title "
		}
		if feedbackParam.Content=="" {
			isContent = "content "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenid+isTitle+isContent,
		}
		this.ServeJSON()
		return
	}

	auth, err := common.CheckAuth(feedbackParam.Openid)
	if err != nil {
		log.Panicln(err)
		return
	}
	if(!auth){
		this.Redirect("/noAuth",302)
		return
	}



	this.Data["json"] = utils.ResultUtil{Code: 200,Msg: "反馈成功！"}
	this.ServeJSON()
	return
}