package organize

import (
	"github.com/astaxie/beego/orm"
)

func DissolutionOrganize(uuid int) (bool,error) {
	o := orm.NewOrm()
	_,err := o.Raw("UPDATE organize SET is_del=2,del_time=now() where uuid=?",uuid).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}