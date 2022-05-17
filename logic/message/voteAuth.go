package message

import "github.com/astaxie/beego/orm"

func IsAuthVote(openId string,typeId int) (int,error)  {
	o := orm.NewOrm()
	var uuid int
	var authority int
	err := o.Raw("SELECT organize_uuid FROM vote WHERE id=?", typeId).QueryRow(&uuid)
	if err != nil {
		return 0,err
	}
	err = o.Raw("SELECT authority FROM member WHERE organize_uuid=? AND openid=?",uuid, openId).QueryRow(&authority)
	if err != nil {
		return 0,err
	}
	return authority,nil
}
