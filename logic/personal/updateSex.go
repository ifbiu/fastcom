package personal

import "github.com/astaxie/beego/orm"

func UpdateSex(openid string,sex int) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE user SET sex=? where openid=?", sex, openid).Exec()
	if err != nil {
		return err
	}
	return nil
}