package models

import "time"

type Organize struct {
	Id int `json:"id"`
	Uuid int `json:"uuid"`
	CoverImg string `json:"coverImg"`
	OrganizeName string `json:"organizeName"`
	Introduce string `json:"introduce"`
	Maximum int `json:"maximum"`
	CreateTime time.Time `json:"createTime"`
}