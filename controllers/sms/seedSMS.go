package sms

import (
	"fastcom/db"
	"fastcom/utils"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"math/rand"
	"time"
)

type SeedSMSController struct {
	beego.Controller
}

func (this *SeedSMSController)Get()  {
	phone := this.GetString("phone")
	key := "sms:"+phone
	if phone == "" {
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必须参数phone"}
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
		code := generateCode()
		utils.SeedSMS(phone,code)
		rds.Set(key,code,time.Minute*10)
		this.Data["json"] = utils.ResultUtil{
			Code: 200,
			Msg: "短信已成功发送，注意查收",
		}
		this.ServeJSON()
	}
	this.Data["json"] = utils.ResultUtil{
		Code: 200,
		Msg: "上次获取的验证码还可以使用",
	}
	this.ServeJSON()
}
func generateCode()  string{
	return fmt.Sprintf("%04v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))//这里面前面的04v是和后面的1000相对应的
}