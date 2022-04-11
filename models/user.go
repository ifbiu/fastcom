package models

type User struct {
	Id int `json:"id" orm:"id"`
	OpenId string `json:"openid" orm:"openid"`
	Phone string `json:"phone" orm:"phone"`
	Image string `json:"image" orm:"image"`
	Sex string `json:"sex" orm:"sex"`
}