package personal

import "github.com/astaxie/beego/orm"

func UpdatePhone(openid string,phone string) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE user SET phone=? where openid=?", phone, openid).Exec()
	if err != nil {
		return err
	}
	return nil
}