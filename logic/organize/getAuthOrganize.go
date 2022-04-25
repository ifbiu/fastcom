package organize

import "github.com/astaxie/beego/orm"


func GetAuthOrganize(openId string,uuid int) (interface{},error) {
	o := orm.NewOrm()
	var authority int

	err := o.Raw("select authority from member where openid=? and organize_uuid=?",openId,uuid).QueryRow(&authority)
	if err != nil {
		return nil,err
	}
	return authority,nil
}