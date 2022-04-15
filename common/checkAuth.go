package common

import (
	"fastcom/db"
	"log"
)

func CheckAuth(openid string) (bool,error) {
	key := "login:"+openid
	rds, err := db.InitRedis()
	if err != nil {
		log.Panicln(err)
		return false,err
	}
	exists, err := rds.Exists(key)
	if err != nil {
		log.Panicln(err)
		return false,err
	}
	if !exists {
		return false,nil
	}
	return true,nil
}
