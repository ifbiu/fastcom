package message

import (
	"errors"
	"fastcom/db"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/streadway/amqp"
	"log"
	"strconv"
)



func ManualEndVote(openId string,typeId int) (error) {
	o := orm.NewOrm()
	var members []string
	_, err := o.Raw("UPDATE vote SET is_end=3,manual_user=?,manual_time=now() WHERE id=?", openId,typeId).Exec()
	if err != nil {
		return err
	}
	var voteItemIds []int
	var alreadyVoteNum int
	_, err = o.Raw("SELECT id FROM vote_item WHERE vote_id=?",typeId).QueryRows(&voteItemIds)
	if err != nil {
		return err
	}
	if len(voteItemIds)==0 {
		return errors.New("投票项空值")
	}
	err = o.Raw("SELECT count(id) FROM vote_success WHERE vote_id=? AND vote_item_id=1 ORDER BY serial_id",typeId).QueryRow(&alreadyVoteNum)
	if err != nil {
		return err
	}
	for i := 0; i < len(voteItemIds); i++ {
		var num int
		err = o.Raw("SELECT count(id) FROM vote_success WHERE vote_id=? AND vote_item_id=1 AND serial_id=? ORDER BY serial_id",typeId,i+1).QueryRow(&num)
		if err != nil {
			return err
		}
		percentageNum, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(num) / float64(alreadyVoteNum)), 64)
		if err != nil {
			return err
		}
		_, err = o.Raw("INSERT INTO vote_result (vote_item_id,vote_num,vote_percentage,create_time) VALUES (?,?,?,now())",voteItemIds[i],num,percentageNum).Exec()
		if err != nil {
			return err
		}
	}
	_,err = o.Raw("SELECT DISTINCT openid FROM vote_success WHERE vote_item_id<>0").QueryRows(&members)
	if err != nil {
		return err
	}
	err = publishVoteResult(members)
	if err != nil {
		return err
	}
	for _, openid := range members {
		_, err := o.Raw("INSERT INTO status (openid,organize_uuid,type,type_id,is_read,create_time) VALUES (?,(SELECT organize_uuid FROM vote WHERE id=?),4,?,1,now())",openid,typeId,typeId).Exec()
		if err != nil {
			return err
		}
	}
	return nil
}

func publishVoteResult(members []string) error {
	//topic := "topic:messageStr:"+uuidStr
	conn, err := db.InitAmqp()
	if err!=nil {
		return err
	}
	// 创建信道
	ch, err := conn.Channel()
	if err!=nil {
		return err
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
	if err!=nil {
		return err
	}

	openIdStr:=""
	if len(members)>0 {
		for i, openid := range members {
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

	log.Printf("%s 发送内容",openIdStr)
	return nil
}