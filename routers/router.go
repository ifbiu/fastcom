package routers

import (
	"fastcom/controllers"
	"fastcom/controllers/login"
	"fastcom/controllers/openId"
	"fastcom/controllers/register"
	"fastcom/controllers/sms"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &login.LoginController{})
    beego.Router("/getOpenId", &openId.GetOpenIdController{})
    beego.Router("/register", &register.RegisterUserController{})
    beego.Router("/isRegister", &register.IsRegisterUserController{})
    beego.Router("/seedPhoneCode", &sms.SeedSMSController{})
    beego.Router("/signOut", &login.SignOutController{})
}
