package message

import (
	"fastcom/common"
	"github.com/astaxie/beego/orm"
)

func AdminApprove(openId string,typeId int,approve int) (error) {
	var openidRes string
	o := orm.NewOrm()
	_ , err := o.Raw("UPDATE approve SET is_approve=?,approve_user=?,approve_time=now() WHERE id=?", approve,openId,typeId).Exec()
	if err != nil {
		return err
	}
	if approve==1 {
		_ , err := o.Raw("INSERT INTO member (organize_uuid,name,openid,authority,create_time) VALUES((SELECT organize_uuid FROM approve WHERE id=?),(SELECT user.nick_name FROM approve left join user on approve.start_user=user.openid WHERE  approve.id=?),(SELECT start_user FROM approve WHERE id=?),3,now())", typeId,typeId,typeId).Exec()
		if err != nil {
			return err
		}
	}
	err = o.Raw("SELECT start_user FROM approve WHERE id=?", typeId).QueryRow(&openidRes)
	if err != nil {
		return err
	}
	_, err = o.Raw("INSERT INTO status (openid,organize_uuid,type,type_id,is_read,create_time) VALUES (?,(SELECT organize_uuid FROM approve WHERE id=?),?,?,?,now())",openidRes,typeId,5,typeId,1).Exec()
	if err != nil {
		return err
	}
	var openids []string
	openids = append(openids,openidRes)
	err = common.AmqpMessage(openids)
	if err != nil {
		return err
	}
	return nil
}