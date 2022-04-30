package models

import "time"

type Status struct {
	Id int `json:"id"`
	Openid string `json:"openid"`
	OrganizeUuid int `json:"organize_uuid"`
	Type int `json:"type"`
	TypeId int `json:"type_id"`
	IsRead int `json:"is_read"`
	ReadTime time.Time `json:"read_time"`
}