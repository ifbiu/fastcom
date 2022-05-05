package message

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func SelectOpenids(uuid string)([]string,error){

	var openids []string
	o := orm.NewOrm()
	_,err := o.Raw("SELECT openid from member where organize_uuid = ?", uuid).QueryRows(&openids)
	if err != nil {
		return openids,err
	}
	return openids,nil
}

func PublishMessage(openid string,openids []string,uuid string,title string,content string) (bool,error) {
	o := orm.NewOrm()
	now := time.Now().Format("2006-01-02 15:04:05")
	exec, err := o.Raw("INSERT INTO notice (organize_uuid,title,content,create_user,create_time) VALUES (?,?,?,?,?)",uuid,title,content,openid,now).Exec()
	if err != nil {
		return false, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return false, err
	}
	for _, openid := range openids {
		_, err := o.Raw("INSERT INTO status (openid,organize_uuid,type,type_id,is_read,read_time) VALUES (?,?,?,?,?,?)",openid,uuid,1,id,1,now).Exec()
		if err != nil {
			return false, err
		}
	}
	return true,nil
}
