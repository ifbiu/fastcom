package models

import "time"

type Notice struct {
	Id int `json:"id"`
	OrganizeUuid int `json:"organize_uuid"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreateUser string `json:"create_user"`
	CreateTime time.Time `json:"create_time"`
}