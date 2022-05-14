package message

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func MessageInfoDel(theType int,typeId int,openId string) (bool,error) {
	o := orm.NewOrm()
	now := time.Now().Format("2006-01-02 15:04:05")
	if theType==1 {
		_, err := o.Raw("update notice SET is_del=2,del_user=?,del_time=? WHERE id=? ", openId ,now, typeId).Exec()
		if err != nil {
			return false, err
		}
		return true,nil
	}else if theType==2 {
		_, err := o.Raw("update vote SET is_del=2,del_user=?,del_time=? WHERE id=? ", openId ,now, typeId).Exec()
		if err != nil {
			return false, err
		}
		return true,nil
	}
	return false,nil
}