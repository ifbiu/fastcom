package message

import (
	"github.com/astaxie/beego/orm"
)

func SelectApprove(openid string,uuid string) (bool,error) {
	o := orm.NewOrm()
	var count = 0
	err := o.Raw("SELECT count(id) from approve where organize_uuid = ? AND start_user=? AND is_approve=0",uuid,openid).QueryRow(&count)
	if err != nil {
		return false,err
	}
	if count==0 {
		return false,nil
	}
	return true,nil
}


func SelectApproveOpenIds(uuid string)([]string,error){

	var openids []string
	o := orm.NewOrm()
	_,err := o.Raw("SELECT openid from member where organize_uuid = ? AND authority in (1,2)  and is_del=1", uuid).QueryRows(&openids)
	if err != nil {
		return openids,err
	}
	return openids,nil
}

func PublishApprove(openid string,openids []string,uuid string) (bool,error) {
	o := orm.NewOrm()
	exec, err := o.Raw("INSERT INTO approve (organize_uuid,start_user,create_time) VALUES (?,?,now())",uuid,openid).Exec()
	if err != nil {
		return false, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return false, err
	}
	for _, opid := range openids {
		_, err := o.Raw("INSERT INTO status (openid,organize_uuid,type,type_id,is_read,create_time) VALUES (?,?,?,?,?,now())",opid,uuid,3,id,1).Exec()
		if err != nil {
			return false, err
		}
	}
	return true,nil
}