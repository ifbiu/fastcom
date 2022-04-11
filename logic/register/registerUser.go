package register

import (
	"fastcom/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)


func AddUserInfo(user *models.User) (bool,bool,error) {
	o := orm.NewOrm()
	var users []models.User
	fmt.Println(user.OpenId)
	_,err := o.Raw("SELECT id FROM user WHERE openid =?", user.OpenId).QueryRows(&users)
	if err != nil {
		return true,false, err
	}
	if len(users) > 0 {
		return false,true, nil
	}
	r,err := o.Raw("INSERT INTO user(openid, phone, image, sex) values (?,?,?,?)", user.OpenId, user.Phone,user.Image,user.Sex).Exec()
	fmt.Println(r.LastInsertId())
	if err != nil {
		return true,false, err
	}
	return true,true, nil
}