package register

import (
	"fastcom/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)


func AddUserInfo(user *models.User) (bool,error) {
	o := orm.NewOrm()
	r,err := o.Raw("INSERT INTO user(openid, phone, image, sex) values (?,?,?,?)", user.OpenId, user.Phone,user.Image,user.Sex).Exec()
	fmt.Println(r)
	if err != nil {
		return false, err
	}
	return true, nil
}