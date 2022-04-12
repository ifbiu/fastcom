package register

import (
	"fastcom/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
)


func AddUserInfo(user *models.User) (bool,bool,error) {
	o := orm.NewOrm()
	var users []models.User
	_,err := o.Raw("SELECT id FROM user WHERE openid =?", user.OpenId).QueryRows(&users)
	if err != nil {
		return true,false, err
	}
	if len(users) > 0 {
		return false,true, nil
	}
	r,err := o.Raw("INSERT INTO user(openid, phone, image, sex, nickname) values (?,?,?,?,?)", user.OpenId, user.Phone,user.Image,user.Sex,user.NickName).Exec()
	fmt.Println(r.LastInsertId())
	if err != nil {
		return true,false, err
	}
	return true,true, nil
}

func IsRegisterUser(openid string) (bool,error) {
	o := orm.NewOrm()
	var users []models.User
	_,err := o.Raw("SELECT id FROM user WHERE openid =?", openid).QueryRows(&users)
	if err != nil {
		log.Fatalln(err)
	}
	if len(users) > 0 {
		return false, nil
	}else{
		return true, nil
	}
}