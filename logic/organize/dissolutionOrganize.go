package organize

import (
	"github.com/astaxie/beego/orm"
)

func DissolutionOrganize(uuid int) (bool,error) {
	o := orm.NewOrm()
	_,err := o.Raw("delete from organize where uuid=?",uuid).Exec()
	if err != nil {
		return false, err
	}
	_,err = o.Raw("delete from member where organize_uuid=?",uuid).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}