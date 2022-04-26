package member

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

func TransferManager(uuid int,openId string,setOpenId string) error {
	o := orm.NewOrm()
	exec1, err := o.Raw("UPDATE member SET authority=2 WHERE organize_uuid=? AND openid=?",uuid,openId).Exec()
	if err != nil {
		return err
	}
	affected1, err := exec1.RowsAffected()
	if err != nil {
		return err
	}
	if affected1==0 {
		return errors.New("更新失败！")
	}

	exec2, err := o.Raw("UPDATE member SET authority=1 WHERE organize_uuid=? AND openid=?",uuid,setOpenId).Exec()
	if err != nil {
		return err
	}

	affected2, err := exec2.RowsAffected()
	if err != nil {
		return err
	}
	if affected2==0 {
		return errors.New("更新失败！")
	}

	return nil
}