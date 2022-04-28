package message

import (
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/db"
	"fastcom/logic/organize"
	"fastcom/utils"
	"fmt"
	"log"
	"strconv"
)

type SubscribeMessageController struct {
	controllers.BaseController
}

func (this *SubscribeMessageController) Get()  {
	openId := this.GetString("openid")
	uuidStr := this.GetString("uuid")
	if openId == "" || uuidStr == ""{
		msg := ""
		if openId == "" {
			msg += "openid "
		}
		if uuidStr == "" {
			msg += "uuid "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少字段："+msg}
		this.ServeJSON()
		return
	}
	uuid, err := strconv.Atoi(uuidStr)
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

	auth, err := common.CheckAuth(openId)
	if err != nil {
		log.Panicln(err)
		return
	}
	if(!auth){
		this.Redirect("/noAuth",302)
		return
	}

	// 是否没有管理员权限
	_, err = organize.GetAuthOrganize(openId, uuid)
	if err != nil {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "组织与成员无联系！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}


	topic := "topic:message:"+uuidStr
	rds, err := db.InitRedis()
	if err != nil {
		log.Fatalln(err)
		return
	}
	res := rds.Subscribe(topic)

	for msg := range res.Channel() {
		// 打印收到的消息
		fmt.Println(msg.Channel)
		fmt.Println(msg.Payload)
	}
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "发布消息失败！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	result := utils.ResultUtil{
		Code: 200,
		Msg: "发布消息成功！",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}