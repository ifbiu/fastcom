package routers

import (
	"fastcom/controllers"
	"fastcom/controllers/login"
	"fastcom/controllers/openId"
	"fastcom/controllers/organize"
	"fastcom/controllers/register"
	"fastcom/controllers/sms"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &login.LoginController{})
	beego.Router("/noAuth", &login.NoAuthController{})
	beego.Router("/getOpenId", &openId.GetOpenIdController{})
	beego.Router("/register", &register.RegisterUserController{})
	beego.Router("/isRegister", &register.IsRegisterUserController{})
	beego.Router("/seedPhoneCode", &sms.SeedSMSController{})
	beego.Router("/signOut", &login.SignOutController{})
	beego.Router("/getOrganizeMenu", &organize.OrganizeMenuController{})
	beego.Router("/addOrganize", &organize.AddOrganizeController{})
	beego.Router("/isMaxOrganize", &organize.IsMaxOrganizeController{})
	beego.Router("/searchOrganize", &organize.SearchOrganizeController{})
	beego.Router("/historyRecord", &organize.HistoryRecordController{})
	beego.Router("/delHistoryRecord", &organize.DelHistoryRecordController{})
	beego.Router("/getUserInfo", &login.GetUserInfoController{})
	beego.Router("/getAuthOrganize", &organize.GetAuthOrganize{})
}
