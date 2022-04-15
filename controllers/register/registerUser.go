package register

import (
	"encoding/json"
	"fastcom/db"
	"fastcom/logic/register"
	"fastcom/models"
	"fastcom/utils"
	"github.com/astaxie/beego"
	"log"
)

type RegisterUserController struct {
	beego.Controller
}

type RequestUser struct {
	OpenId string `json:"openid"`
	Phone string `json:"phone"`
	Image string `json:"image"`
	Sex int `json:"sex"`
	NickName string `json:"nickName"`
	Code string `json:"code"`
}

func (this *RegisterUserController)Post()  {
	userParam := &RequestUser{}
	json.Unmarshal(this.Ctx.Input.RequestBody, userParam)
	key := "sms:"+userParam.Phone
	if userParam.OpenId=="" ||
		userParam.Phone=="" ||
		userParam.Image=="" ||
		userParam.Sex == 0 ||
		userParam.NickName == "" ||
		userParam.Code == "" {
		var (
			isOpenid string = ""
			isPhone string = ""
			isImage string = ""
			isSex string = ""
			isNickname string = ""
			isCode string = ""
		)
		if userParam.OpenId=="" {
			isOpenid = "openid "
		}
		if userParam.Phone=="" {
			isPhone = "phone "
		}
		if userParam.Image=="" {
			isImage = "image "
		}
		if userParam.Sex==0 {
			isSex = "sex "
		}
		if userParam.NickName=="" {
			isNickname = "nickname "
		}
		if userParam.Code=="" {
			isCode = "code "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenid+isPhone+isImage+isSex+isNickname+isCode,
		}
		this.ServeJSON()
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
	if resCode != userParam.Code {
		this.Data["json"] = utils.ResultUtil{
			Code: 500,
			Msg: "验证码错误",
		}
		this.ServeJSON()
		return
	}
	user := models.User{
		Openid: userParam.OpenId,
		Phone: userParam.Phone,
		Image: userParam.Image,
		Sex: userParam.Sex,
		NickName: userParam.NickName,
	}
	isExist,status,err := register.AddUserInfo(&user)
	if err != nil {
		log.Panicln(err)
	}
	if !isExist {
		this.Data["json"] = utils.ResultUtil{
			Code: 500,
			Msg: "注册失败，用户已存在",
		}
		this.ServeJSON()
		return
	}
	if !status {
		this.Data["json"] = utils.ResultUtil{
			Code: 500,
			Msg: "注册失败",
		}
		this.ServeJSON()
		return
	}

	this.Data["json"] = utils.ResultUtil{
		Code: 200,
		Msg: "注册成功",
	}
	this.ServeJSON()
	return
}