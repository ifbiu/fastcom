package member

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

func SetAdmin(uuid int,setOpenId string) (error) {
	o := orm.NewOrm()
	_ = o.Begin()
	exec, err := o.Raw("UPDATE member SET authority=2 WHERE organize_uuid = ? AND openid= ?", uuid, setOpenId).Exec()
	if err != nil {
		_ = o.Rollback()
		return err
	}
	_ = o.Commit()
	affected, err := exec.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("更新失败")
	}
	return nil
}