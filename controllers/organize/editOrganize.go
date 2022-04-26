package organize

import (
	"encoding/json"
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/logic/organize"
	"fastcom/utils"
	"fmt"
	"log"
	"strconv"
)

type EditOrganizeController struct {
	controllers.BaseController
}

type RequestEditOrganizeInfo struct {
	Openid string `json:"openid"`
	OrganizeName string `json:"organizeName"`
	CoverImg string `json:"coverImg"`
	Introduce string `json:"introduce"`
	Uuid string `json:"uuid"`
}

func (this *EditOrganizeController) Post() {
	organizeParam := &RequestEditOrganizeInfo{}
	json.Unmarshal(this.Ctx.Input.RequestBody, organizeParam)
	if organizeParam.Openid=="" ||
		organizeParam.OrganizeName=="" ||
		organizeParam.CoverImg=="" ||
		organizeParam.Introduce == "" {
		var (
			isOpenid string = ""
			isOrganizeName string = ""
			isCoverImg string = ""
			isIntroduce string = ""
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
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenid+isOrganizeName+isCoverImg+isIntroduce,
		}
		this.ServeJSON()
		return
	}
	uuid, err := strconv.Atoi(organizeParam.Uuid)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "uuid格式错误",
		}
		this.Data["json"] = &result
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

	// 是否没有管理员权限
	authOrganize, err := organize.GetAuthOrganize(organizeParam.Openid, uuid)
	if err != nil {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "服务异常！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	if authOrganize != 1 && authOrganize != 2  {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉，您没有管理员权限！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	infoStatus, err := organize.EditOrganizeInfo(
		organizeParam.OrganizeName,
		organizeParam.CoverImg,
		organizeParam.Introduce,
		uuid,
	)
	if err != nil || !infoStatus {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "修改失败",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	result := utils.ResultUtil{
		Code: 200,
		Msg: "修改成功",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}