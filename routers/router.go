package routers

import (
	"fastcom/controllers"
	"fastcom/controllers/login"
	"fastcom/controllers/member"
	"fastcom/controllers/message"
	"fastcom/controllers/openId"
	"fastcom/controllers/organize"
	"fastcom/controllers/personal"
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
	beego.Router("/editOrganize", &organize.EditOrganizeController{})
	beego.Router("/signOutOrganize", &organize.SignOutOrganizeController{})
	beego.Router("/dissolutionOrganize", &organize.DissolutionOrganizeController{})
	beego.Router("/getMemberInfo", &member.GetMemberInfoController{})
	beego.Router("/deleteMember", &member.DeleteMemberController{})
	beego.Router("/changeRemarks", &member.ChangeRemarksController{})
	beego.Router("/setAdmin", &member.SetAdminController{})
	beego.Router("/cancelAdmin", &member.CancelAdminController{})
	beego.Router("/transferManager", &member.TransferManagerController{})
	beego.Router("/publishMessage", &message.PublishMessageController{})
	beego.Router("/updateHeadPortrait", &personal.UpdateHeadPortraitController{})
	beego.Router("/updateNickName", &personal.UpdateNickNameController{})
	beego.Router("/updatePhone", &personal.UpdatePhoneController{})
	beego.Router("/updateSex", &personal.UpdateSexController{})
	beego.Router("/feedback", &personal.FeedbackController{})
	beego.Router("/messageMenu", &message.MessageMenuController{})
	beego.Router("/messageInfo", &message.MessageInfoController{})
	beego.Router("/isAuthDel", &message.IsAuthDelController{})
	beego.Router("/messageInfoDel", &message.MessageInfoDelController{})
	beego.Router("/isMessageRead", &message.IsMessageReadController{})
	beego.Router("/publishApprove", &message.PublishApproveController{})
	beego.Router("/publishVote", &message.PublishVoteController{})
	beego.Router("/angleMark", &message.AngleMarkController{})
	beego.Router("/adminApprove", &message.AdminApproveController{})
	beego.Router("/voteOperation", &message.VoteOperationController{})
	beego.Router("/voteAuth", &message.VoteAuthController{})
	beego.Router("/manualEndVote", &message.ManualEndVoteController{})
	beego.Router("/autoEndVote", &message.AutoEndVoteController{})
	beego.Router("/isPartakeVote", &organize.IsPartakeVoteController{})
	beego.Router("/historyMessage", &organize.HistoryMessageController{})
	beego.Router("/historyInfo", &organize.HistoryInfoController{})
}
