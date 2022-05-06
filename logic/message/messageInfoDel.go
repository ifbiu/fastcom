package message

import "github.com/astaxie/beego/orm"

func MessageInfoDel(theType int,typeId int,openId string) (bool,error) {
	o := orm.NewOrm()
	if theType==1 {
		_, err := o.Raw("update notice SET is_del=2,del_user=? WHERE id=? ", openId, typeId).Exec()
		if err != nil {
			return false, err
		}
		return true,nil
	}
	return false,nil
}