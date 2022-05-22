package common

import (
	"errors"
	"fastcom/db"
	"github.com/streadway/amqp"
	"log"
)

func AmqpMessage(openids []string) (error) {
	//topic := "topic:messageStr:"+uuidStr
	conn, err := db.InitAmqp()
	if err!=nil {
		return err
	}
	// 创建信道
	ch, err := conn.Channel()
	if err != nil {
		return errors.New("Failed to open a channel")
	}
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
	if err != nil {
		return errors.New("Failed to declare an exchange")
	}
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

	log.Printf("%s 发送内容 %s",openIdStr)
	return nil
}
