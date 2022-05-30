package organize

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/wxnacy/wgo/arrays"
	"math/rand"
	"strconv"
	"time"
)

func GenerateUuid() (string,error) {
	var resStr string
	var uuids []int64
	o := orm.NewOrm()
	res1 := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	res2 := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	resStr = "1"+res1+res2[:3]
	_, err := o.Raw("SELECT uuid FROM organize").QueryRows(&uuids)
	if err != nil {
		return "",err
	}
	resInt, err := strconv.Atoi(resStr)
	if err != nil {
		return "", err
	}
	index := arrays.ContainsInt(uuids, int64(resInt))
	if index == -1 {
		return resStr,nil
	} else {
		uuid, err := GenerateUuid()
		if err != nil {
			return "", err
		}
		return uuid,nil
	}
}

func AddOrganize(uuid int,maximum int,openid string,organizeName string,coverImg string,introduce string,authorName string) (bool,error) {
	o := orm.NewOrm()
	now := time.Now().Format("2006-01-02 15:04:05")
	r1,err := o.Raw("INSERT INTO organize(uuid, cover_img, organize_name, introduce,maximum, create_time) values (?,?,?,?,?,?)",uuid,coverImg,organizeName,introduce,maximum,now).Exec()
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
	_,err = o.Raw("INSERT INTO member(organize_uuid, name, openid,authority, create_time) values (?,?,?,?,?)",uuid,authorName,openid,1,now).Exec()
	if err != nil {
		return false, err
	}
	return true, nil
}