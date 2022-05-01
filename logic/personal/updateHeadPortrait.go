package personal

import "github.com/astaxie/beego/orm"

func UpdateHeadPortrait(openid string,image string) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE user SET image=? where openid=?", image, openid).Exec()
	if err != nil {
		return err
	}
	return nil
}