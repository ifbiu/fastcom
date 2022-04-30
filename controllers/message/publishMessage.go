package message

import (
	"encoding/json"
	"fastcom/common"
	"fastcom/controllers"
	"fastcom/db"
	"fastcom/logic/message"
	"fastcom/logic/organize"
	"fastcom/utils"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type PublishMessageController struct {
	controllers.BaseController
}

type RequestPublishMessage struct {
	Openid string `json:"openid"`
	Uuid string `json:"uuid"`
	Title string `json:"title"`
	Content string `json:"content"`
}

func (this *PublishMessageController) Post()  {
	publishParam := &RequestPublishMessage{}
	json.Unmarshal(this.Ctx.Input.RequestBody, publishParam)
	if publishParam.Openid=="" ||
		publishParam.Uuid=="" ||
		publishParam.Title=="" ||
		publishParam.Content == "" {
		var (
			isOpenid string = ""
			isUuid string = ""
			isTitle string = ""
			isContent string = ""
		)
		if publishParam.Openid=="" {
			isOpenid = "openid "
		}
		if publishParam.Uuid=="" {
			isUuid = "uuid "
		}
		if publishParam.Title=="" {
			isTitle = "title "
		}
		if publishParam.Content=="" {
			isContent = "content "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenid+isUuid+isTitle+isContent,
		}
		this.ServeJSON()
		return
	}
	uuid, err := strconv.Atoi(publishParam.Uuid)
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

	auth, err := common.CheckAuth(publishParam.Openid)
	if err != nil {
		log.Panicln(err)
		return
	}
	if(!auth){
		this.Redirect("/noAuth",302)
		return
	}

	// 是否没有管理员权限
	authOrganize, err := organize.GetAuthOrganize(publishParam.Openid, uuid)
	if err != nil {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "组织与成员无联系！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	if authOrganize == 3 {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉，您没有发布的权限！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	openids, err := message.SelectOpenids(publishParam.Uuid)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉，搜索不到成员！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	_, err = message.PublishMessage(openids,publishParam.Uuid,publishParam.Title,publishParam.Content)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "公告发布失败！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}


	//topic := "topic:messageStr:"+uuidStr
	conn, err := db.InitAmqp()
	if err != nil {
		result := utils.ResultUtil{
			Code: 500,
			Msg: "服务异常！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	// 创建信道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//for _, openid := range openids {
	//
	//}
	// 声明交换机
	q, err := ch.QueueDeclare(
		"fastcom", // 队列名字
		false,   // 消息是否持久化
		false,   // 不使用的时候删除队列
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	openIdStr:=""
	if len(openids)>0 {
		for i, openid := range openids {
			if i==0 {
				openIdStr = openid
			}else{
				openIdStr = openIdStr + ","+openid
			}
		}
	}

	// 推送消息
	err = ch.Publish(
		"",     // exchange（交换机名字），这里忽略
		q.Name, // 路由参数，这里使用队列名字作为路由参数
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(openIdStr),  // 消息内容
		})

	log.Printf("%s 发送内容 %s",publishParam.Openid,openIdStr)

	result := utils.ResultUtil{
		Code: 200,
		Msg: "发布公告成功！",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}