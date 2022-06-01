package organize

import "github.com/astaxie/beego/orm"

func SignOutOrganize(openId string,uuid int) (bool,error) {
	o := orm.NewOrm()
	_ = o.Begin()
	_,err := o.Raw("update member SET is_del=2,del_time=now() where organize_uuid=? and openid=?",uuid,openId).Exec()
	if err != nil {
		_ = o.Rollback()
		return false, err
	}
	var memberId int
	err = o.Raw("SELECT id FROM member WHERE  organize_uuid=? and openid=?",uuid,openId).QueryRow(&memberId)
	if err != nil {
		_ = o.Rollback()
		return false, err
	}
	_, err = o.Raw("INSERT INTO status (openid,organize_uuid,type,type_id,is_read,create_time) VALUES (?,?,?,?,?,now())",openId,uuid,6,memberId,1).Exec()
	if err != nil {
		_ = o.Rollback()
		return false, err
	}

	_ = o.Commit()
	return true, nil
}