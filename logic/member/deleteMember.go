package member

import "github.com/astaxie/beego/orm"

func DeleteMember(uuid int,delOpenId string) (bool,error) {
	o := orm.NewOrm()
	r1,err := o.Raw("UPDATE member SET is_del=2 where organize_uuid=? and openid=?",uuid,delOpenId).Exec()
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
