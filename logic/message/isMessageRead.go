package message

import "github.com/astaxie/beego/orm"

func IsMessageRead(theType int,typeId int,openId string) (bool,error) {
	o := orm.NewOrm()
	_, err := o.Raw("update status SET is_read=2,read_time=now() WHERE openid=? AND type=? AND type_id=? ", openId, theType, typeId).Exec()
	if err != nil {
		return false, err
	}
	return true,nil
}