package organize

import (
	"fastcom/common"
	"github.com/astaxie/beego/orm"
	"time"
)

type RequestMessageMenu struct {
	Title string `json:"title"`
	ShowTime  time.Time`json:"showTime"`
}

type ResponseMessageMenu struct {
	Id int `json:"id"`
	Title string `json:"title"`
	ShowTime  string`json:"showTime"`
}

func HistoryMessage(uuid int,theType int,page int,pageSize int) (interface{},error) {
	o := orm.NewOrm()
	var typeIds []int
	pageRes := 0
	if page > 1 {
		pageRes = pageSize * (page - 1)
	}

	_, err := o.Raw("SELECT DISTINCT type_id from status where organize_uuid=? AND type=? order by create_time desc limit ?,?", uuid,theType, pageRes, pageSize).QueryRows(&typeIds)
	if err != nil {
		return nil, err
	}

	responseMessageMenu := make([]ResponseMessageMenu, len(typeIds))
	for i, id := range typeIds {
		requestMessageMenu := RequestMessageMenu{}
		if theType == 1 { // 公告
			err = o.Raw("SELECT notice.title as title,notice.create_time as show_time from notice where id=?", id).QueryRow(&requestMessageMenu)
			if err != nil {
				return nil, err
			}
		} else if theType == 2 { // 投票
			err = o.Raw("SELECT vote.title as title,vote.create_time as show_time from vote where id=?", id).QueryRow(&requestMessageMenu)
			if err != nil {
				return nil, err
			}
		}
		responseMessageMenu[i].Id = id
		responseMessageMenu[i].Title = requestMessageMenu.Title
		responseMessageMenu[i].ShowTime = common.FormatTime(requestMessageMenu.ShowTime)
	}
	return responseMessageMenu, err
}