package routers

import (
	"fastcom/controllers"
	"fastcom/controllers/login"
	"fastcom/controllers/openId"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &login.LoginController{})
    beego.Router("/getOpenId", &openId.GetOpenIdController{})
}
