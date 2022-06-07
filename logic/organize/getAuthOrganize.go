package organize

import (
	"github.com/astaxie/beego/orm"
)


func GetAuthOrganize(openId string,uuid int) (interface{},error) {
	o := orm.NewOrm()
	var authority int

	err := o.Raw("select authority from member where openid=? and organize_uuid=? and is_del=1",openId,uuid).QueryRow(&authority)
	if err != nil {
		return nil,err
	}
	return authority,nil
}

func GetIsDelOrganize(openId string,uuid int) (int,error) {
	o := orm.NewOrm()
	var isDel int

	err := o.Raw("select is_del from member where openid=? and organize_uuid=? order by id desc limit 1",openId,uuid).QueryRow(&isDel)
	if err != nil {
		return 0,err
	}
	return isDel,nil
}