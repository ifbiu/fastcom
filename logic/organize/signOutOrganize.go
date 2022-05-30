package organize

import "github.com/astaxie/beego/orm"

func SignOutOrganize(openId string,uuid int) (bool,error) {
	o := orm.NewOrm()
	_ = o.Begin()
	r1,err := o.Raw("update member SET is_del=2 where organize_uuid=? and openid=?",uuid,openId).Exec()
	if err != nil {
		_ = o.Rollback()
		return false, err
	}
	organizeId, err := r1.RowsAffected()
	if err != nil {
		_ = o.Rollback()
		return false, err
	}
	if organizeId == 0 {
		_ = o.Rollback()
		return false, err
	}
	_ = o.Commit()
	return true, nil
}