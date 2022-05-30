package message

import (
	"fastcom/common"
	"fastcom/logic/organize"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type RequestMessageMenu struct {
	Title string `json:"title"`
	OrganizeName string `json:"organizeName"`
	OrganizeUuid string `json:"organizeUuid"`
	Type int `json:"type"`
	TypeId int `json:"typeId"`
	IsRead int `json:"isRead"`
	IsOrganizeDel int `json:"isOrganizeDel"`
	ShowTime  time.Time`json:"showTime"`
}

type ResponseMessageMenu struct {
	Title string `json:"title"`
	OrganizeName string `json:"organizeName"`
	Type int `json:"type"`
	TypeId int `json:"typeId"`
	IsRead int `json:"isRead"`
	IsOutOrganize bool `json:"isOutOrganize"`
	IsOrganizeDel int `json:"isOrganizeDel"`
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
	requestMessageMenu := RequestMessageMenu{}
	responseMessageMenu := make([]ResponseMessageMenu,len(selectMessage))
	for i, message := range selectMessage {
		if message.Type==1 { // 公告
			err = o.Raw("SELECT organize.uuid as organize_uuid,notice.title as title,organize.organize_name as organize_name,status.type as type,status.type_id as type_id,status.is_read as is_read,notice.create_time as show_time,organize.is_del as is_organize_del from status left join notice on status.type_id = notice.id left join organize on status.organize_uuid=organize.uuid where status.id=?",message.Id).QueryRow(&requestMessageMenu)
			if err != nil {
				return nil,err
			}
			responseMessageMenu[i].Title = requestMessageMenu.Title
		}else if message.Type==2 { // 投票
			err = o.Raw("SELECT organize.uuid as organize_uuid,vote.title as title,organize.organize_name as organize_name,status.type as type,status.type_id as type_id,status.is_read as is_read,vote.create_time as show_time,organize.is_del as is_organize_del from status left join vote on status.type_id = vote.id left join organize on status.organize_uuid=organize.uuid where status.id=?",message.Id).QueryRow(&requestMessageMenu)
			if err != nil {
				return nil,err
			}
			responseMessageMenu[i].Title = requestMessageMenu.Title
		}else if message.Type==3 { // 审核
			err = o.Raw("SELECT organize.uuid as organize_uuid,organize.organize_name as organize_name,status.type as type,status.type_id as type_id,status.is_read as is_read,approve.create_time as show_time,organize.is_del as is_organize_del from status left join approve on status.type_id = approve.id left join organize on status.organize_uuid=organize.uuid where status.id=?",message.Id).QueryRow(&requestMessageMenu)
			if err != nil {
				return nil,err
			}
			responseMessageMenu[i].Title = "加入组织申请"
		}else if message.Type==4{ // 投票结果
			err = o.Raw("SELECT organize.uuid as organize_uuid,vote.title as title,organize.organize_name as organize_name,status.type as type,status.type_id as type_id,status.is_read as is_read,vote_result.create_time as show_time,organize.is_del as is_organize_del from status left join vote on status.type_id = vote.id left join organize on status.organize_uuid=organize.uuid left join vote_result on vote.id=vote_result.vote_id where status.id=?",message.Id).QueryRow(&requestMessageMenu)
			if err != nil {
				return nil,err
			}
			responseMessageMenu[i].Title = requestMessageMenu.Title
		}else if message.Type==5{ // 审核结果
			err = o.Raw("SELECT organize.uuid as organize_uuid,organize.organize_name as organize_name,status.type as type,status.type_id as type_id,status.is_read as is_read,approve.approve_time as show_time, organize.is_del as is_organize_del from status left join approve on status.type_id = approve.id left join organize on status.organize_uuid=organize.uuid where status.id=?",message.Id).QueryRow(&requestMessageMenu)
			if err != nil {
				return nil,err
			}
			responseMessageMenu[i].Title = "加入组织审核结果"
		}
		uuid, err := strconv.Atoi(requestMessageMenu.OrganizeUuid)
		if err != nil {
			return nil, err
		}
		_, err = organize.GetAuthOrganize(openId, uuid)
		if err != nil {
			responseMessageMenu[i].IsOutOrganize = true
		}else{
			responseMessageMenu[i].IsOutOrganize = false
		}
		responseMessageMenu[i].OrganizeName = requestMessageMenu.OrganizeName
		responseMessageMenu[i].Type = requestMessageMenu.Type
		responseMessageMenu[i].TypeId = requestMessageMenu.TypeId
		responseMessageMenu[i].IsRead = requestMessageMenu.IsRead
		responseMessageMenu[i].IsOrganizeDel = requestMessageMenu.IsOrganizeDel
		responseMessageMenu[i].ShowTime = common.FormatTime(requestMessageMenu.ShowTime)
	}
	return responseMessageMenu, err
}