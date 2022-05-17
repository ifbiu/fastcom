package message

import "github.com/astaxie/beego/orm"

func ManualEndVote(openId string,typeId int) (error) {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE vote SET is_end=3,manual_user=?,manual_time=now() WHERE id=?", openId,typeId).Exec()
	if err != nil {
		return err
	}
	return nil
}