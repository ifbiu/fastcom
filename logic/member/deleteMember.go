package member

import "github.com/astaxie/beego/orm"

func DeleteMember(uuid int,openId,delOpenId string) (bool,error) {
	o := orm.NewOrm()
	r1,err := o.Raw("UPDATE member SET is_del=3,del_admin=?,del_time=now() where organize_uuid=? and openid=?",openId,uuid,delOpenId).Exec()
	if err != nil {
		return false, err
	}
	organizeId, err := r1.RowsAffected()
	if err != nil {
		return false, err
	}
	if organizeId == 0 {
		return false, err
	}
	return true, nil
}
