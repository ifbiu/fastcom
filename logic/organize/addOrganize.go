package organize

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func AddOrganize(openid string,organizeName string,coverImg string,introduce string,authorName string) (bool,error) {
	o := orm.NewOrm()
	now := time.Now().Format("2006-01-02 15:04:05")
	r1,err := o.Raw("INSERT INTO organize(cover_img, organize_name, introduce, create_time) values (?,?,?,?)",coverImg,organizeName,introduce,now).Exec()
	if err != nil {
		return false, err
	}
	organizeId, err := r1.LastInsertId()
	if err != nil {
		return false, err
	}
	if organizeId == 0 {
		return false, err
	}
	_,err = o.Raw("INSERT INTO member(organize_id, name, openid,authority, create_time) values (?,?,?,?,?)",organizeId,authorName,openid,1,now).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}