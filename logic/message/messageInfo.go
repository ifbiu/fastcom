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

func GetMessageInfo(theType int,typeId int,openId string) (interface{},error) {
	o := orm.NewOrm()
	delId := 0
	// 公告
	if theType == 1 {
		err := o.Raw("SELECT is_del FROM notice WHERE id=?", typeId).QueryRow(&delId)
		if err != nil {
			return nil, err
		}
		if delId == 2 {
			delMessage := delMessageResponse{}
			err := o.Raw("SELECT del_user FROM notice WHERE id=?", typeId).QueryRow(&delMessage)
			if err != nil {
				return nil, err
			}
			return delMessage,nil
		}
		err = o.Raw("SELECT is_del FROM notice WHERE id=?", typeId).QueryRow(&delId)
		if err != nil {
			return nil, err
		}
	}

	return []string{},nil
}