package message

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

func IsVote(openId string,typeId int) (bool) {
	o := orm.NewOrm()
	resId := 0
	err := o.Raw("SELECT vote_item_id FROM vote_success WHERE openid=?,vote_id=?",openId,typeId).QueryRow(&resId)
	if err != nil {
		return false
	}
	return true
}

func VoteOperation(openId string,vote int,typeId int,serialIds []int) (error) {
	o := orm.NewOrm()
	serialIdAll := ""
	for i, v := range serialIds {
		if i==0 {
			serialIdAll = strconv.Itoa(v)
		}else{
			serialIdAll += ","+strconv.Itoa(v)
		}
	}
	if vote==1 {
		for _, serialId := range serialIds {
			_,err := o.Raw("UPDATE vote_success SET vote_item_id=? WHERE openid=? AND vote_id=? AND serial_id=?",1,openId,typeId,serialId).Exec()
			if err != nil {
				return err
			}
		}
	}else if vote==2 {
		_,err := o.Raw("UPDATE vote_success SET vote_item_id=? WHERE openid=? AND vote_id=?",2,openId,typeId).Exec()
		if err != nil {
			return err
		}
	}

	return nil
}