package organize

import "github.com/astaxie/beego/orm"

func SignOutOrganize(openId string,uuid int) (bool,error) {
	o := orm.NewOrm()
	r1,err := o.Raw("delete from member where organize_uuid=? and openid=?",uuid,openId).Exec()
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