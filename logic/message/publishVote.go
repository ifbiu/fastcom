package message

import (
	"fastcom/db"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
	"time"
)

func PublishVote(openid string,openids []string,uuid string,title string,maxNum int,isAbstained int,endTime int64,items []string) (bool,error) {
	o := orm.NewOrm()
	now := time.Now().Format("2006-01-02 15:04:05")
	end := time.Unix(endTime,0).Format("2006-01-02 15:04:05")

	exec, err := o.Raw("INSERT INTO vote (organize_uuid,title,max_num,is_abstained,create_user,create_time,end_time) VALUES (?,?,?,?,?,?,?)",uuid,title,maxNum,isAbstained,openid,now,end).Exec()
	if err != nil {
		return false, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return false, err
	}
	for i, item := range items {
		index := i+1
		_, err := o.Raw("INSERT INTO vote_item (vote_id,serial_id,content) VALUES (?,?,?)",id,index,item).Exec()
		if err != nil {
			return false, err
		}
	}
	for _, openidOne := range openids {
		_, err := o.Raw("INSERT INTO status (openid,organize_uuid,type,type_id,is_read,create_time) VALUES (?,?,?,?,?,?)",openidOne,uuid,2,id,1,now).Exec()
		if err != nil {
			return false, err
		}
		for i, _ := range items {
			index := i+1
			_, err := o.Raw("INSERT INTO vote_success (openid,vote_id,serial_id,vote_item_id) VALUES (?,?,?,?)",openidOne,id,index,0).Exec()
			if err != nil {
				return false, err
			}
		}
	}
	key := "vote:"+strconv.Itoa(int(id))
	rds, err := db.InitRedis()
	if err != nil {
		log.Panicln(err)
		return false,err
	}
	exists, err := rds.Exists(key)
	if err != nil {
		return false, err
	}
	if exists {
		return false,nil
	}
	t := time.Unix(endTime,0).Sub(time.Now())
	err = rds.Set(key,id, t)
	if err != nil {
		return false, err
	}
	return true,nil
}