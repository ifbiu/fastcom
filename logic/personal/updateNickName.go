package personal

import "github.com/astaxie/beego/orm"

func UpdateNickName(openid string,nickName string) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE user SET nick_name=? where openid=?", nickName, openid).Exec()
	if err != nil {
		return err
	}
	return nil
}