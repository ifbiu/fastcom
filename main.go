package main

import (
	_ "fastcom/db"
	_ "fastcom/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

