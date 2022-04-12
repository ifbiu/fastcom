package models

type User struct {
	Id int `json:"id" orm:"id"`
	OpenId string `json:"openid" orm:"openid"`
	Phone string `json:"phone" orm:"phone"`
	Image string `json:"image" orm:"image"`
	Sex int `json:"sex" orm:"sex"`
	NickName string `json:"nickName" orm:"nickname"`
}