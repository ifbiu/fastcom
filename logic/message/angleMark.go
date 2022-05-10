package message

import "github.com/astaxie/beego/orm"

func GetAngleMark(openId string) (int,error) {
	o := orm.NewOrm()
	num := 0
	err := o.Raw("SELECT count(id) FROM status WHERE openId = ? AND is_read = 1", openId).QueryRow(&num)
	if err != nil {
		return 0, err
	}
	return num,nil
}