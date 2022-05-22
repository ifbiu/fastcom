package message

import (
	"errors"
	"fastcom/common"
	"fmt"
	"github.com/astaxie/beego/orm"
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
		var percentageNum float64
		if alreadyVoteNum == 0 {
			percentageNum = 0
		}else{
			percentageNum, err = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(num) / float64(alreadyVoteNum)), 64)
			if err != nil {
				return err
			}
		}
		_, err = o.Raw("INSERT INTO vote_result (vote_id,vote_item_id,vote_num,vote_percentage,create_time) VALUES (?,?,?,?,now())",typeId,voteItemIds[i],num,percentageNum).Exec()
		if err != nil {
			return err
		}
	}
	_,err = o.Raw("SELECT DISTINCT openid FROM vote_success WHERE vote_item_id<>0 AND vote_id=?",typeId).QueryRows(&members)
	if err != nil {
		return err
	}
	err = common.AmqpMessage(members)
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