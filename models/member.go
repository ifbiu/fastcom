package models

import "time"

type Member struct {
	Id int `json:"id"`
	Uuid int `json:"uuid"`
	OrganizeId int `json:"organizeId"`
	Name string `json:"name"`
	Openid string `json:"openid"`
	Authority int `json:"authority"`
	CreateTime time.Time `json:"createTime"`
}