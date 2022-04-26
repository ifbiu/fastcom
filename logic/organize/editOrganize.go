package organize

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func EditOrganizeInfo(organizeName string,coverImg string,introduce string,uuid int)(bool,error){
	o := orm.NewOrm()
	now := time.Now().Format("2006-01-02 15:04:05")
	r1,err := o.Raw("UPDATE organize SET organize_name=?,introduce=?,cover_img=?,update_time=? where uuid=?",organizeName,introduce,coverImg,now,uuid).Exec()
	if err != nil {
		return false, err
	}
	organizeId, err := r1.RowsAffected()
	if err != nil {
		return false, err
	}
	if organizeId == 0 {
		return false, err
	}
	return true, nil
}