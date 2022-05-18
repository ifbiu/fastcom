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

type SelectMessage struct {
	Id int `json:"id"`
	Type int `json:"type"`
	TypeId int `json:"typeId"`
}

func GetMessageMenu(openId string,page int,pageSize int) (interface{},error) {
	o := orm.NewOrm()
	pageRes := 0
	if page>1 {
		pageRes = pageSize*(page-1)
	}


	selectMessage := []SelectMessage{}
	_,err := o.Raw("SELECT id,type,type_id from status where openid=? order by create_time desc limit ?,?", openId,pageRes,pageSize).QueryRows(&selectMessage)
	if err != nil {
		return nil,err
	}

	responseMessageMenu := make([]ResponseMessageMenu,len(selectMessage))
	for i, message := range selectMessage {
		if message.Type==1 {
			requestMessageMenu := RequestMessageMenu{}
			err = o.Raw("SELECT notice.title as title,organize.organize_name as organize_name,status.type as type,status.type_id as type_id,status.is_read as is_read,notice.create_time as show_time from status left join notice on status.type_id = notice.id left join organize on status.organize_uuid=organize.uuid where status.id=?",message.Id).QueryRow(&requestMessageMenu)
			if err != nil {
				return nil,err
			}
			responseMessageMenu[i].Title = requestMessageMenu.Title
		    responseMessageMenu[i].OrganizeName = requestMessageMenu.OrganizeName
		    responseMessageMenu[i].Type = requestMessageMenu.Type
		    responseMessageMenu[i].TypeId = requestMessageMenu.TypeId
		    responseMessageMenu[i].IsRead = requestMessageMenu.IsRead
			responseMessageMenu[i].ShowTime = common.FormatTime(requestMessageMenu.ShowTime)
		}else if message.Type==2 {
			requestMessageMenu := RequestMessageMenu{}
			err = o.Raw("SELECT vote.title as title,organize.organize_name as organize_name,status.type as type,status.type_id as type_id,status.is_read as is_read,vote.create_time as show_time from status left join vote on status.type_id = vote.id left join organize on status.organize_uuid=organize.uuid where status.id=?",message.Id).QueryRow(&requestMessageMenu)
			if err != nil {
				return nil,err
			}
			responseMessageMenu[i].Title = requestMessageMenu.Title
			responseMessageMenu[i].OrganizeName = requestMessageMenu.OrganizeName
			responseMessageMenu[i].Type = requestMessageMenu.Type
			responseMessageMenu[i].TypeId = requestMessageMenu.TypeId
			responseMessageMenu[i].IsRead = requestMessageMenu.IsRead
			responseMessageMenu[i].ShowTime = common.FormatTime(requestMessageMenu.ShowTime)
		}else if message.Type==3 {
			requestMessageMenu := RequestMessageMenu{}
			err = o.Raw("SELECT organize.organize_name as organize_name,status.type as type,status.type_id as type_id,status.is_read as is_read,approve.create_time as show_time from status left join approve on status.type_id = approve.id left join organize on status.organize_uuid=organize.uuid where status.id=?",message.Id).QueryRow(&requestMessageMenu)
			if err != nil {
				return nil,err
			}
			responseMessageMenu[i].Title = "加入组织审核"
			responseMessageMenu[i].OrganizeName = requestMessageMenu.OrganizeName
			responseMessageMenu[i].Type = requestMessageMenu.Type
			responseMessageMenu[i].TypeId = requestMessageMenu.TypeId
			responseMessageMenu[i].IsRead = requestMessageMenu.IsRead
			responseMessageMenu[i].ShowTime = common.FormatTime(requestMessageMenu.ShowTime)
		}else if message.Type==4{
			requestMessageMenu := RequestMessageMenu{}
			err = o.Raw("SELECT vote.title as title,organize.organize_name as organize_name,status.type as type,status.type_id as type_id,status.is_read as is_read,vote.create_time as show_time from status left join vote on status.type_id = vote.id left join organize on status.organize_uuid=organize.uuid where status.id=?",message.Id).QueryRow(&requestMessageMenu)
			if err != nil {
				return nil,err
			}
			responseMessageMenu[i].Title = requestMessageMenu.Title
			responseMessageMenu[i].OrganizeName = requestMessageMenu.OrganizeName
			responseMessageMenu[i].Type = requestMessageMenu.Type
			responseMessageMenu[i].TypeId = requestMessageMenu.TypeId
			responseMessageMenu[i].IsRead = requestMessageMenu.IsRead
			responseMessageMenu[i].ShowTime = common.FormatTime(requestMessageMenu.ShowTime)
		}

	}
	return responseMessageMenu, err
}