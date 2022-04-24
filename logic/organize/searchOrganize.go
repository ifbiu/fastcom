package organize

import (
	"github.com/astaxie/beego/orm"
)

type SearchOrganizeAll struct {
	Uuid int `json:"uuid"`
	CoverImg string `json:"coverImg"`
	OrganizeName string `json:"organizeName"`
	Introduce string `json:"introduce"`
	CreateTime string `json:"createTime"`
	AdminCount int `json:"adminCount"`
	MemberCount int `json:"memberCount"`
	SuperAdminName string `json:"superAdminName"`
}

func SearchOrganize(uuid int) (SearchOrganizeAll,error) {
	o := orm.NewOrm()
	var searchOrganizeAll SearchOrganizeAll
	var adminIds int
	var memberIds int
	var superAdminName string

	err := o.Raw("SELECT * FROM organize WHERE uuid = ? order by create_time desc", uuid).QueryRow(&searchOrganizeAll)
	if err != nil {
		return SearchOrganizeAll{}, err
	}
	err = o.Raw("SElECT count(id) from member where organize_uuid = ? and authority in (1,2)", uuid).QueryRow(&adminIds)
	if err != nil {
		return SearchOrganizeAll{}, err
	}
	err = o.Raw("SElECT count(id) from member where organize_uuid = ? and authority = 3", uuid).QueryRow(&memberIds)
	if err != nil {
		return SearchOrganizeAll{}, err
	}
	err = o.Raw("SElECT name from member where organize_uuid = ? and authority = 1", uuid).QueryRow(&superAdminName)
	if err != nil {
		return SearchOrganizeAll{}, err
	}
	searchOrganizeAll.AdminCount = adminIds
	searchOrganizeAll.MemberCount = memberIds
	searchOrganizeAll.SuperAdminName = superAdminName
	return searchOrganizeAll,nil
}