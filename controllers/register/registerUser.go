package register

import (
	"fastcom/db"
	"fastcom/logic/register"
	"fastcom/models"
	"fastcom/utils"
	"github.com/astaxie/beego"
	"log"
	"strconv"
)

type RegisterUserController struct {
	beego.Controller
}

func (this *RegisterUserController)Post()  {
	openid := this.GetString("openid")
	phone := this.GetString("phone")
	image := this.GetString("image")
	sex, _ := strconv.Atoi(this.GetString("sex"))
	nickname := this.GetString("nickName")
	code := this.GetString("code")
	key := "sms:"+phone
	if openid=="" || phone=="" || image=="" || sex == 0 || nickname == "" || code == "" {
		var (
			isOpenid string = ""
			isPhone string = ""
			isImage string = ""
			isSex string = ""
			isNickname string = ""
			isCode string = ""
		)
		if openid=="" {
			isOpenid = "openid "
		}
		if phone=="" {
			isPhone = "phone "
		}
		if image=="" {
			isImage = "image "
		}
		if sex==0 {
			isSex = "sex "
		}
		if nickname=="" {
			isNickname = "nickname "
		}
		if code=="" {
			isCode = "code "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenid+isPhone+isImage+isSex+isNickname+isCode,
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
		NickName: nickname,
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