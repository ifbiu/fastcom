package organize

import (
	"encoding/json"
	"fastcom/common"
	"fastcom/logic/organize"
	"fastcom/utils"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"strconv"
)

type AddOrganizeController struct {
	beego.Controller
}

type RequestOrganize struct {
	Openid string `json:"openid"`
	OrganizeName string `json:"organizeName"`
	CoverImg string `json:"coverImg"`
	Introduce string `json:"introduce"`
	AuthorName string `json:"authorName"`
}

func (this *AddOrganizeController) Post()  {
	organizeParam := &RequestOrganize{}
	json.Unmarshal(this.Ctx.Input.RequestBody, organizeParam)
	if organizeParam.Openid=="" ||
		organizeParam.OrganizeName=="" ||
		organizeParam.CoverImg=="" ||
		organizeParam.Introduce == "" ||
		organizeParam.AuthorName == "" {
		var (
			isOpenid string = ""
			isOrganizeName string = ""
			isCoverImg string = ""
			isIntroduce string = ""
			isAuthorName string = ""
		)
		if organizeParam.Openid=="" {
			isOpenid = "openid "
		}
		if organizeParam.OrganizeName=="" {
			isOrganizeName = "organizeName "
		}
		if organizeParam.CoverImg=="" {
			isCoverImg = "coverImg "
		}
		if organizeParam.Introduce=="" {
			isIntroduce = "introduce "
		}
		if organizeParam.AuthorName=="" {
			isAuthorName = "authorName "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenid+isOrganizeName+isCoverImg+isIntroduce+isAuthorName,
		}
		this.ServeJSON()
		return
	}
	auth, err := common.CheckAuth(organizeParam.Openid)
	if err != nil {
		log.Panicln(err)
		return
	}
	if(!auth){
		this.Redirect("/noAuth",302)
		return
	}
	uuid, err := strconv.Atoi(utils.GenerateNum(10))
	if err != nil {
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "生成uuid失败！"}
		this.ServeJSON()
		return
	}
	// 默认最大承载人数
	maximum := 200
	addOrganize, err := organize.AddOrganize(uuid,maximum,organizeParam.Openid, organizeParam.OrganizeName, organizeParam.CoverImg, organizeParam.Introduce, organizeParam.AuthorName)
	if err != nil || !addOrganize {
		fmt.Println(err)
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "服务异常，请联系管理员！"}
		this.ServeJSON()
		return
	}

	this.Data["json"] = utils.ResultUtil{Code: 200,Msg: "success",
	}
	this.ServeJSON()
	return
}
