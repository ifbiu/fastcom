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


type PublishApproveController struct {
	controllers.BaseController
}

type RequestPublishApprove struct {
	Openid string `json:"openid"`
	Uuid string `json:"uuid"`
}

func (this *PublishApproveController) Post()  {
	publishParam := &RequestPublishApprove{}
	json.Unmarshal(this.Ctx.Input.RequestBody, publishParam)
	if publishParam.Openid=="" ||
		publishParam.Uuid==""  {
		var (
			isOpenid = ""
			isUuid = ""
		)
		if publishParam.Openid=="" {
			isOpenid = "openid "
		}
		if publishParam.Uuid=="" {
			isUuid = "uuid "
		}
		this.Data["json"] = utils.ResultUtil{Code: 500,Msg: "缺少必传参数："+isOpenid+isUuid}
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
	// 是否没有管理员权限
	authOrganize, _ := organize.GetAuthOrganize(publishParam.Openid, uuid)

	if authOrganize != nil{
		result := utils.ResultUtil{
			Code: 500,
			Msg: "抱歉，您已经在该组织中！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	selectApprove, err := message.SelectApprove(publishParam.Openid, publishParam.Uuid)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "服务异常！",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}
	if selectApprove {
		result := utils.ResultUtil{
			Code: 200,
			Msg: "审核中...",
		}
		this.Data["json"] = &result
		this.ServeJSON()
		return
	}

	openids, err := message.SelectApproveOpenIds(publishParam.Uuid)
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

	_, err = message.PublishApprove(publishParam.Openid,openids,publishParam.Uuid)
	if err != nil {
		fmt.Println(err)
		result := utils.ResultUtil{
			Code: 500,
			Msg: "发布审核消息失败！",
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
		Msg: "申请成功",
	}
	this.Data["json"] = &result
	this.ServeJSON()
	return
}