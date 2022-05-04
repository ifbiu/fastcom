package db

import (
	"github.com/astaxie/beego"
	"github.com/streadway/amqp"
	"log"
)

func init()  {
	conn, err := InitAmqp()
	if err != nil {
		log.Panicln("rabbitmq connect error!")
		return
	}
	defer conn.Close()
	log.Println("rabbitmq connect success...")
}

func InitAmqp() (*amqp.Connection,error) {
	user := beego.AppConfig.String("amqpUser")
	pwd := beego.AppConfig.String("amqpPassword")
	host := beego.AppConfig.String("amqpHost")
	port := beego.AppConfig.String("amqpPort")

	conn, err := amqp.Dial("amqp://"+user+":"+pwd+"@"+host+":"+port+"/")
	if err != nil {
		return nil,err
	}
	return conn,nil
}

func CloseAmqp(conn *amqp.Connection)  {
	conn.Close()
}