package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//切记：导入驱动包
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func init() {
	driverName := beego.AppConfig.String("mysqlDriverName")

	//注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	//数据库连接
	user := beego.AppConfig.String("mysqlUser")
	pwd := beego.AppConfig.String("mysqlPassword")
	host := beego.AppConfig.String("mysqlHost")
	port := beego.AppConfig.String("mysqlPort")
	dbname := beego.AppConfig.String("mysqlDbname")

	//dbConn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	err := orm.RegisterDataBase("default", driverName, dbConn)
	if err != nil {
		log.Println("mysql connect error!")
		return
	}
	log.Println("mysql connect success...")
}