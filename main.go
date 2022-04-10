package main

import (
	_ "fastcom/routers"
	_ "fastcom/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

