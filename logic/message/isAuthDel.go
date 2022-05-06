package message

import "github.com/astaxie/beego/orm"

func IsAuthDel(theType int,typeId int,openId string) (int,error) {
	o := orm.NewOrm()
	res := 0
	if theType==1 {
		err := o.Raw("SELECT member.authority FROM notice left join member on notice.organize_uuid=member.organize_uuid WHERE member.openid=? and notice.id=?", openId, typeId).QueryRow(&res)
		if err != nil {
			return 0, err
		}
	}
	return res,nil
}