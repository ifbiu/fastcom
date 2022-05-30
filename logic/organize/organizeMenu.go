package organize

import (
	"fastcom/models"
	"github.com/astaxie/beego/orm"
)

type OrganizeMenu struct {
	Uuid int `json:"uuid"`
	CoverImg string `json:"coverImg"`
	OrganizeName string `json:"organizeName"`
	Introduce string `json:"introduce"`
	CreateTime string `json:"createTime"`
	AdminCount int `json:"adminCount"`
	MemberCount int `json:"memberCount"`
	SuperAdminName string `json:"superAdminName"`
}

func GetMenu(openId string,status string) ([]OrganizeMenu,error) {
	o := orm.NewOrm()
	var organizeIdAll = []string{};
	organizeAll := []models.Organize{}
	var err error
	if status == "admin" {
		_,err = o.Raw("SELECT organize_uuid FROM member WHERE openid =? and authority in (1,2) and is_del=1 order by id desc", openId).QueryRows(&organizeIdAll)
		if err != nil {
			return nil, err
		}
	}

	if status == "member" {
		_,err = o.Raw("SELECT organize_uuid FROM member WHERE openid =? and authority = 3 and is_del=1 order by id desc", openId).QueryRows(&organizeIdAll)
		if err != nil {
			return nil, err
		}
	}

	if len(organizeIdAll) == 0 {
		return []OrganizeMenu{},nil
	}

	wen := ""
	for i := 0; i < len(organizeIdAll); i++ {
		if i==0 {
			wen += "?"
		}else{
			wen += ",?"
		}
	}
	_,err = o.Raw("SELECT * FROM organize WHERE uuid in ("+wen+")  and is_del=1 order by create_time desc", organizeIdAll).QueryRows(&organizeAll)
	if err != nil {
		return nil, err
	}
	organizeMenu := make([]OrganizeMenu,len(organizeAll))
	for i, organize := range organizeAll {
		var adminIds int
		var memberIds int
		var superAdminName string
		err := o.Raw("SElECT count(id) from member where organize_uuid = ? and authority in (1,2)  and is_del=1", organize.Uuid).QueryRow(&adminIds)
		if err != nil {
			return nil, err
		}
		err = o.Raw("SElECT count(id) from member where organize_uuid = ? and authority = 3  and is_del=1", organize.Uuid).QueryRow(&memberIds)
		if err != nil {
			return nil, err
		}
		err = o.Raw("SElECT name from member where organize_uuid = ? and authority = 1  and is_del=1", organize.Uuid).QueryRow(&superAdminName)
		if err != nil {
			return nil, err
		}
		organizeMenu[i].Uuid = organize.Uuid
		organizeMenu[i].CoverImg = organize.CoverImg
		organizeMenu[i].OrganizeName = organize.OrganizeName
		organizeMenu[i].Introduce = organize.Introduce
		organizeMenu[i].CreateTime = organize.CreateTime.Format("2006-01-02")
		organizeMenu[i].AdminCount = adminIds
		organizeMenu[i].MemberCount = memberIds
		organizeMenu[i].SuperAdminName = superAdminName
	}
	return organizeMenu,nil
}
