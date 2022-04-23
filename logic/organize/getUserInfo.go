package organize

import (
	"github.com/astaxie/beego/orm"
)

type responseResult struct {
	Image string `json:"image"`
	Sex int `json:"sex"`
	NickName string `json:"nickName"`
}

func GetUserInfo(openid string) (interface{},error) {
	o := orm.NewOrm()
	var user responseResult

	err := o.Raw("select nick_name,image,sex from user where openid=?",openid).QueryRow(&user)
	if err != nil {
		return nil,err
	}
	return user,nil
}