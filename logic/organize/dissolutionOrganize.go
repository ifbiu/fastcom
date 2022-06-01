package organize

import (
	"github.com/astaxie/beego/orm"
)

func GetOrganizeOpenIds(uuid int) ([]string,error) {
	o := orm.NewOrm()
	var openIds []string
	_,err := o.Raw("SELECT openid FROM member WHERE organize_uuid=? AND is_del=1",uuid).QueryRows(&openIds)
	if err != nil {
		return nil, err
	}
	return openIds,nil
}

func DissolutionOrganize(uuid int) (bool,error) {
	o := orm.NewOrm()
	_,err := o.Raw("UPDATE organize SET is_del=2,del_time=now() where uuid=?",uuid).Exec()
	if err != nil {
		return false, err
	}
	var organizeId int
	err = o.Raw("SELECT id FROM organize WHERE uuid=?", uuid).QueryRow(&organizeId)
	if err != nil {
		return false, err
	}
	openIds, err := GetOrganizeOpenIds(uuid)
	if err != nil {
		return false, err
	}
	for _, openId := range openIds {
		_, err = o.Raw("INSERT INTO status (openid,organize_uuid,type,type_id,is_read,create_time) VALUES (?,?,?,?,?,now())",openId,uuid,8,organizeId,1).Exec()
		if err != nil {
			_ = o.Rollback()
			return false, err
		}
	}

	return true, nil
}