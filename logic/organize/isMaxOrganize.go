package organize

import (
	"github.com/astaxie/beego/orm"
)

func IsOrganizeDel(uuid int) (int,error) {
	o := orm.NewOrm()
	var isDel int
	err := o.Raw("SElECT is_del from organize where uuid = ?", uuid).QueryRow(&isDel)
	if err != nil {
		return 0,err
	}
	return isDel,nil
}

func IsMaxOrganize(uuid int,openid string) (bool,bool,error) {
	o := orm.NewOrm()
	var uuidCount int
	var maximum int
	var isHave int
	err := o.Raw("SElECT count(id) from member where openid = ? and organize_uuid = ? and is_del = 1", openid,uuid).QueryRow(&isHave)
	if err != nil {
		return false,false,err
	}
	if isHave != 0 {
		return false,false,err
	}
	err = o.Raw("SElECT maximum from organize where uuid = ? and is_del = 1", uuid).QueryRow(&maximum)
	if err != nil {
		return false,false,err
	}
	err = o.Raw("SElECT count(id) from member where organize_uuid = ?  and is_del = 1", uuid).QueryRow(&uuidCount)
	if err != nil {
		return false,false,err
	}
	if uuidCount < maximum{
		return true,true,err
	}else{
		return true,false,err
	}
}