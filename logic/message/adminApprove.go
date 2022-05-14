package message

import "github.com/astaxie/beego/orm"

func AdminApprove(openId string,typeId int,approve int) (error) {
	o := orm.NewOrm()
	_ , err := o.Raw("UPDATE approve SET is_approve=?,approve_user=?,approve_time=now() WHERE id=?", approve,openId,typeId).Exec()
	if err != nil {
		return err
	}
	return nil
}