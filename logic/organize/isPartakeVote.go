package organize

import "github.com/astaxie/beego/orm"

func GetAuthVote(openId string,typeId int) (int,error) {
	o := orm.NewOrm()
	var uuid,authority int
	err := o.Raw("SELECT organize_uuid FROM vote WHERE id=?",typeId).QueryRow(&uuid)
	if err != nil {
		return 0, err
	}
	err = o.Raw("select authority from member where openid=? and organize_uuid=?",openId,uuid).QueryRow(&authority)
	if err != nil {
		return 0,err
	}
	return authority,nil
}

func IsPartakeVote(typeId int,openId string) (bool,error) {
	o := orm.NewOrm()
	var openIds []string
	_, err := o.Raw("SELECT distinct openid FROM vote_success WHERE vote_id=?",typeId).QueryRows(&openIds)
	if err != nil {
		return false, err
	}
	if len(openIds)==0 {
		return false,nil
	}
	for _, opId := range openIds {
		if opId==openId {
			return true,nil
		}
	}
	return false, nil
}
