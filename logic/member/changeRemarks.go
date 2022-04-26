package member

import "github.com/astaxie/beego/orm"

func ChangeRemarks(uuid int,openId string,newName string) (bool,error) {
	o := orm.NewOrm()
	result, err := o.Raw("UPDATE member SET name = ? WHERE organize_uuid = ? AND openid = ?",newName,uuid,openId).Exec()
	if err != nil {
		return false,err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false,err
	}
	if affected == 0 {
		return false,err
	}
	return true,nil
}