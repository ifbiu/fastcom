package message

import "github.com/astaxie/beego/orm"

func CheckApproveAuth(openid string,typeId int) (int,error) {
	approveAuth := 0
	o := orm.NewOrm()
	err := o.Raw("SELECT member.authority FROM approve join member on approve.organize_uuid = member.organize_uuid WHERE approve.id = ? AND member.openid = ?", typeId, openid).QueryRow(&approveAuth)
	if err != nil {
		return 0,err
	}
	return approveAuth,nil
}