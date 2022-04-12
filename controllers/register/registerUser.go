package register

import (
	"encoding/json"
	"fastcom/db"
	"fastcom/logic/register"
	"fastcom/models"
	"fastcom/utils"
	"fmt"
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
	u := &RequestUser{}
	json.Unmarshal(this.Ctx.Input.RequestBody, u)
	fmt.Println(u)
	key := "sms:"+u.Phone
	if u.OpenId=="" || u.Phone=="" || u.Image=="" || u.Sex == 0 || u.NickName == "" || u.Code == "" {
		var (
			isOpenid string = ""
			isPhone string = ""
			isImage string = ""
			isSex string = ""
			isNickname string = ""
			isCode string = ""
		)
		if u.OpenId=="" {
			isOpenid = "openid "
		}
		if u.Phone=="" {
			isPhone = "phone "
		}
		if u.Image=="" {
			isImage = "image "
		}
		if u.Sex==0 {
			isSex = "sex "
		}
		if u.NickName=="" {
			isNickname = "nickname "
		}
		if u.Code=="" {
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
	if resCode != u.Code {
		this.Data["json"] = utils.ResultUtil{
			Code: 500,
			Msg: "验证码错误",
		}
		this.ServeJSON()
	}
	user := models.User{
		OpenId: u.OpenId,
		Phone: u.Phone,
		Image: u.Image,
		Sex: u.Sex,
		NickName: u.NickName,
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