package register

import (
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

func (this *RegisterUserController)Post()  {
	openid := this.GetString("openid")
	phone := this.GetString("phone")
	image := this.GetString("image")
	sex := this.GetString("sex")
	code := this.GetString("code")
	key := "sms:"+phone
	if openid=="" || phone=="" || image=="" || sex == "" || code == "" {
		var (
			isopenid string = ""
			isphone string = ""
			isimage string = ""
			issex string = ""
			iscode string = ""
		)
		if openid=="" {
			isopenid = "openid "
		}
		if phone=="" {
			isphone = "phone "
		}
		if image=="" {
			isimage = "image "
		}
		if sex=="" {
			issex = "sex "
		}
		if code=="" {
			iscode = "code "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isopenid+isphone+isimage+issex+iscode,
		}
		this.ServeJSON()
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
	}
	resCode, err := rds.Get(key)
	if resCode != code {
		this.Data["json"] = utils.ResultUtil{
			Code: 500,
			Msg: "验证码错误",
		}
		this.ServeJSON()
	}
	user := models.User{
		OpenId: openid,
		Phone: phone,
		Image: image,
		Sex: sex,
	}
	status,err := register.AddUserInfo(&user)
	if err != nil {
		log.Panicln(err)
	}
	if !status {
		this.Data["json"] = utils.ResultUtil{
			Code: 500,
			Msg: "注册失败",
		}
		this.ServeJSON()
	}

	this.Data["json"] = utils.ResultUtil{
		Code: 200,
		Msg: "注册成功",
	}
	this.ServeJSON()
}