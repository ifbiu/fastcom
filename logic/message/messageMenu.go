package message

import (
	"fastcom/common"
	"github.com/astaxie/beego/orm"
	"time"
)

type RequestMessageMenu struct {
	Title string `json:"title"`
	OrganizeName string `json:"organizeName"`
	Type int `json:"type"`
	TypeId int `json:"typeId"`
	IsRead int `json:"isRead"`
	ShowTime  time.Time`json:"showTime"`
}

type ResponseMessageMenu struct {
	Title string `json:"title"`
	OrganizeName string `json:"organizeName"`
	Type int `json:"type"`
	TypeId int `json:"typeId"`
	IsRead int `json:"isRead"`
	ShowTime  string`json:"showTime"`

}

func GetMessageMenu(openId string,page int,pageSize int) (interface{},error) {
	o := orm.NewOrm()
	requestMessageMenu := []RequestMessageMenu{}
	pageRes := 0
	if page>1 {
		pageRes = pageSize*(page-1)
	}
	_,err := o.Raw("SELECT notice.title as title,organize.organize_name as organize_name,status.type as type,status.type_id as type_id,status.is_read as is_read,notice.create_time as show_time from status left join notice on status.type_id = notice.id left join organize on status.organize_uuid=organize.uuid where status.openid=? order by notice.create_time desc limit ?,?", openId,pageRes,pageSize).QueryRows(&requestMessageMenu)
	if err != nil {
		return nil,err
	}

	responseMessageMenu := make([]ResponseMessageMenu,len(requestMessageMenu))
	for i, v := range requestMessageMenu {
        responseMessageMenu[i].Title = v.Title
        responseMessageMenu[i].OrganizeName = v.OrganizeName
        responseMessageMenu[i].Type = v.Type
        responseMessageMenu[i].TypeId = v.TypeId
        responseMessageMenu[i].IsRead = v.IsRead
		responseMessageMenu[i].ShowTime = common.FormatTime(v.ShowTime)
	}
	return responseMessageMenu,nil
}