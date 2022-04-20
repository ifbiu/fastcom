package sms

import (
	"fastcom/controllers"
	"fastcom/db"
	"fastcom/utils"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type SeedSMSController struct {
	controllers.BaseController
}

func (this *SeedSMSController)Get()  {
	phone := this.GetString("phone")
	key := "sms:"+phone
	if phone == "" {
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必须参数phone"}
		this.ServeJSON()
		return
	}
	rds,err := db.InitRedis()
	defer rds.Close()
	if err != nil {
		log.Panicln(err)
	}
	code := generateCode()
	utils.SeedSMS(phone,code)
	rds.Set(key,code,time.Minute*10)
	this.Data["json"] = utils.ResultUtil{
		Code: 200,
		Msg: "短信已成功发送，注意查收",
	}
	this.ServeJSON()
	return
}
func generateCode()  string{
	return fmt.Sprintf("%04v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))//这里面前面的04v是和后面的1000相对应的
}