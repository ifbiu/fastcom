package models

import "time"

type Organize struct {
	Id int `json:"id"`
	CoverImg string `json:"coverImg"`
	OrganizeName string `json:"organizeName"`
	Introduce string `json:"introduce"`
	CreateTime time.Time `json:"createTime"`
}