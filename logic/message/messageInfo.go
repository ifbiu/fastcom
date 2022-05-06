package message

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type delMessageResponse struct {
	DelUser string `json:"delUser"`
}
type messageResponse struct {
	Title string `json:"title"`
	Content string `json:"content"`
	OrganizeName string `json:"organizeName"`
	CreateUser string `json:"createUser"`
	CreateTime time.Time `json:"createTime"`
}

type messageOutput struct {
	Title string `json:"title"`
	Content string `json:"content"`
	OrganizeName string `json:"organizeName"`
	CreateUser string `json:"createUser"`
	ReadCount int `json:"readCount"`
	CreateTime string `json:"createTime"`
}

func GetMessageInfo(theType int,typeId int) (interface{},error) {
	o := orm.NewOrm()
	delId := 0
	readCount := 0
	err := o.Raw("SELECT count(*) as read_count FROM status WHERE type=? AND type_id=? AND is_read=2", theType,typeId).QueryRow(&readCount)
	if err != nil {
		return nil, err
	}
	// 公告
	if theType == 1 {
		err := o.Raw("SELECT is_del FROM notice WHERE id=?", typeId).QueryRow(&delId)
		if err != nil {
			return nil, err
		}
		if delId == 2 {
			delMessage := delMessageResponse{}
			err := o.Raw("SELECT member.name as del_user FROM notice join member on notice.del_user = member.openid WHERE notice.id=? and notice.organize_uuid=member.organize_uuid", typeId).QueryRow(&delMessage)
			if err != nil {
				return nil, err
			}
			return delMessage,nil
		}
		messageRes := messageResponse{}
		err = o.Raw("SELECT title,content,organize.organize_name as organize_name,member.name as create_user,notice.create_time as create_time FROM notice left join organize on notice.organize_uuid = organize.uuid left join member on notice.create_user=member.openid WHERE notice.id=? AND notice.organize_uuid=member.organize_uuid", typeId).QueryRow(&messageRes)
		if err != nil {
			return nil, err
		}
		messageOut := messageOutput{}
		messageOut.Title = messageRes.Title
		messageOut.Content = messageRes.Content
		messageOut.OrganizeName = messageRes.OrganizeName
		messageOut.CreateUser = messageRes.CreateUser
		messageOut.ReadCount = readCount
		messageOut.CreateTime =messageRes.CreateTime.Format("2006年01月02日 15:04")
		return messageOut,nil
	}

	return []string{},nil
}