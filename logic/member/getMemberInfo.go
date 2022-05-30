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
	IsMe bool `json:"isMe"`
}

type responseData2 struct {
	Openid string `json:"openid"`
	Image string `json:"image"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	IsMe bool `json:"isMe"`
}

func GetMemberInfo(authOrganize interface{},uuid int,openId string) (interface{},error) {
	o := orm.NewOrm()
	if authOrganize == 3 {
		var users memberInfo1
		var Data3,Data2,Data1 []responseData1
		_,err := o.Raw("select member.openid as openid,image,name from member join user on member.openid = user.openid where authority=1 and organize_uuid=? and member.is_del=1",uuid).QueryRows(&Data1)
		if err != nil {
			return nil,err
		}
		_,err = o.Raw("select member.openid as openid,image,name from member join user on member.openid = user.openid where authority=2 and organize_uuid=? and member.is_del=1",uuid).QueryRows(&Data2)
		if err != nil {
			return nil,err
		}
		_,err = o.Raw("select member.openid as openid,image,name from member join user on member.openid = user.openid where authority=3 and organize_uuid=? and member.is_del=1",uuid).QueryRows(&Data3)
		if err != nil {
			return nil,err
		}
		for i1, data1 := range Data1 {
			if data1.Openid == openId {
				Data1[i1].IsMe = true
			}else{
				Data1[i1].IsMe = false
			}
		}
		for i2, data2 := range Data2 {
			if data2.Openid == openId {
				Data2[i2].IsMe = true
			}else{
				Data2[i2].IsMe = false
			}
		}
		for i3, data3 := range Data3 {
			if data3.Openid == openId {
				Data3[i3].IsMe = true
			}else{
				Data3[i3].IsMe = false
			}
		}
		users.Manager = Data1
		users.Admin = Data2
		users.Member = Data3
		return users,nil
	}else{
		var users memberInfo2
		var Data3,Data2,Data1 []responseData2
		_,err := o.Raw("select member.openid as openid,image,name,phone from member join user on member.openid = user.openid where authority=1 and organize_uuid=? and member.is_del=1",uuid).QueryRows(&Data1)
		if err != nil {
			return nil,err
		}
		_,err = o.Raw("select member.openid as openid,image,name,phone from member join user on member.openid = user.openid where authority=2 and organize_uuid=? and member.is_del=1",uuid).QueryRows(&Data2)
		if err != nil {
			return nil,err
		}
		_,err = o.Raw("select member.openid as openid,image,name,phone from member join user on member.openid = user.openid where authority=3 and organize_uuid=? and member.is_del=1",uuid).QueryRows(&Data3)
		if err != nil {
			return nil,err
		}
		for i1, data1 := range Data1 {
			if data1.Openid == openId {
				Data1[i1].IsMe = true
			}else{
				Data1[i1].IsMe = false
			}
		}
		for i2, data2 := range Data2 {
			if data2.Openid == openId {
				Data2[i2].IsMe = true
			}else{
				Data2[i2].IsMe = false
			}
		}
		for i3, data3 := range Data3 {
			if data3.Openid == openId {
				Data3[i3].IsMe = true
			}else{
				Data3[i3].IsMe = false
			}
		}
		users.Manager = Data1
		users.Admin = Data2
		users.Member = Data3
		return users,nil
	}
}