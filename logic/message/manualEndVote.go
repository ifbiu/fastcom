package message

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)



func ManualEndVote(openId string,typeId int) (error) {
	o := orm.NewOrm()
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
	return nil
}