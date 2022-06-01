package member

import "github.com/astaxie/beego/orm"

func DeleteMember(uuid int,openId,delOpenId string) (bool,error) {
	o := orm.NewOrm()
	_,err := o.Raw("UPDATE member SET is_del=3,del_admin=?,del_time=now() where organize_uuid=? and openid=?",openId,uuid,delOpenId).Exec()
	if err != nil {
		return false, err
	}
	var memberId int
	err = o.Raw("SELECT id FROM member WHERE  organize_uuid=? and openid=?",uuid,openId).QueryRow(&memberId)
	if err != nil {
		_ = o.Rollback()
		return false, err
	}
	_, err = o.Raw("INSERT INTO status (openid,organize_uuid,type,type_id,is_read,create_time) VALUES (?,?,?,?,?,now())",delOpenId,uuid,7,memberId,1).Exec()
	if err != nil {
		_ = o.Rollback()
		return false, err
	}
	return true, nil
}
