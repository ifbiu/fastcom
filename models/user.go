package models

import "time"

type User struct {
	Id int `json:"id"`
	Openid string `json:"openid"`
	Phone string `json:"phone"`
	Image string `json:"image"`
	Sex int `json:"sex"`
	NickName string `json:"nickName"`
	CreateTime time.Time `json:"createTime"`
}