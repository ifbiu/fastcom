package member

import (
	"github.com/astaxie/beego/orm"
)

type memberInfo1 struct {
	Manager []responseData1 `json:"manager"`
	Admin []responseData1 `json:"admin"`
	Member []responseData1 `json:"member"`
}

type memberInfo2 struct {
	Manager []responseData2 `json:"manager"`
	Admin []responseData2 `json:"admin"`
	Member []responseData2 `json:"member"`
}
type responseData1 struct {
	Openid string `json:"openid"`
	Image string `json:"image"`
	Name string `json:"name"`
}

type responseData2 struct {
	Openid string `json:"openid"`
	Image string `json:"image"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

func GetMemberInfo(authOrganize interface{},uuid int) (interface{},error) {
	o := orm.NewOrm()
	if authOrganize == 3 {
		var users memberInfo1
		var Data3,Data2,Data1 []responseData1
		_,err := o.Raw("select member.openid as openid,image,name from member join user on member.openid = user.openid where authority=1 and organize_uuid=?",uuid).QueryRows(&Data1)
		if err != nil {
			return nil,err
		}
		_,err = o.Raw("select member.openid as openid,image,name from member join user on member.openid = user.openid where authority=2 and organize_uuid=?",uuid).QueryRows(&Data2)
		if err != nil {
			return nil,err
		}
		_,err = o.Raw("select member.openid as openid,image,name from member join user on member.openid = user.openid where authority=3 and organize_uuid=?",uuid).QueryRows(&Data3)
		if err != nil {
			return nil,err
		}
		users.Manager = Data1
		users.Admin = Data2
		users.Member = Data3
		return users,nil
	}else{
		var users memberInfo2
		var Data3,Data2,Data1 []responseData2
		_,err := o.Raw("select member.openid as openid,image,name,phone from member join user on member.openid = user.openid where authority=1 and organize_uuid=?",uuid).QueryRows(&Data1)
		if err != nil {
			return nil,err
		}
		_,err = o.Raw("select member.openid as openid,image,name,phone from member join user on member.openid = user.openid where authority=2 and organize_uuid=?",uuid).QueryRows(&Data2)
		if err != nil {
			return nil,err
		}
		_,err = o.Raw("select member.openid as openid,image,name,phone from member join user on member.openid = user.openid where authority=3 and organize_uuid=?",uuid).QueryRows(&Data3)
		if err != nil {
			return nil,err
		}
		users.Manager = Data1
		users.Admin = Data2
		users.Member = Data3
		return users,nil
	}
}