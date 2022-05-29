package organize

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type delMessageResponse struct {
	DelUser string `json:"delUser"`
}
type messageResponse struct {
	Title string `json:"title"`
	Content string `json:"content"`
	OrganizeName string `json:"organizeName"`
	CreateUser string `json:"createUser"`
	CreateTime time.Time `json:"createTime"`
}

type voteResponse struct {
	Title string `json:"title"`
	OrganizeName string `json:"organizeName"`
	CreateUser string `json:"createUser"`
	IsAbstained int `json:"isAbstained"`
	MaxNum int `json:"maxNum"`
	CreateTime time.Time `json:"createTime"`
	EndTime time.Time `json:"endTime"`
}


type voteOutput1 struct {
	Title string `json:"title"`
	Content []string `json:"content"`
	IsEnd int `json:"isEnd"`
	MaxNum int `json:"maxNum"`
	VoteNum int `json:"voteNum"`
	IsVoteNum int `json:"isVoteNum"`
	OrganizeName string `json:"organizeName"`
	CreateUser string `json:"createUser"`
	CreateTime string `json:"createTime"`
	EndTime string `json:"endTime"`
}
type voteOutput2 struct {
	Title string `json:"title"`
	IsEnd int `json:"isEnd"`
	VoteNum int `json:"voteNum"`
	IsVoteNum int `json:"isVoteNum"`
	OrganizeName string `json:"organizeName"`
	CreateUser string `json:"createUser"`
	CreateTime string `json:"createTime"`
	EndTime string `json:"endTime"`
}

type voteOutput3 struct {
	Title string `json:"title"`
	IsEnd int `json:"isEnd"`
	VoteNum int `json:"voteNum"`
	IsVoteNum int `json:"isVoteNum"`
	ManualUser string `json:"manualUser"`
	ManualTime string `json:"manualTime"`
	OrganizeName string `json:"organizeName"`
	CreateUser string `json:"createUser"`
	CreateTime string `json:"createTime"`
	EndTime string `json:"endTime"`
}

type endManualResponse struct {
	ManualUser string `json:"manualUser"`
	ManualTime time.Time `json:"manualTime"`
}

type messageOutput struct {
	Title string `json:"title"`
	Content string `json:"content"`
	OrganizeName string `json:"organizeName"`
	CreateUser string `json:"createUser"`
	ReadCount int `json:"readCount"`
	CreateTime string `json:"createTime"`
}

func GetHistoryInfo(theType int,typeId int) (interface{},error) {
	o := orm.NewOrm()
	delId := 0
	readCount := 0
	err := o.Raw("SELECT count(*) as read_count FROM status WHERE type=? AND type_id=? AND is_read=2", theType,typeId).QueryRow(&readCount)
	if err != nil {
		return nil, err
	}
	// 公告
	if theType == 1 {
		err := o.Raw("SELECT is_del FROM notice WHERE id=?", typeId).QueryRow(&delId)
		if err != nil {
			return nil, err
		}
		if delId == 2 {
			delMessage := delMessageResponse{}
			err := o.Raw("SELECT member.name as del_user FROM notice join member on notice.del_user = member.openid WHERE notice.id=? and notice.organize_uuid=member.organize_uuid", typeId).QueryRow(&delMessage)
			if err != nil {
				return nil, err
			}
			return delMessage,nil
		}
		messageRes := messageResponse{}
		err = o.Raw("SELECT title,content,organize.organize_name as organize_name,member.name as create_user,notice.create_time as create_time FROM notice left join organize on notice.organize_uuid = organize.uuid left join member on notice.create_user=member.openid WHERE notice.id=? AND notice.organize_uuid=member.organize_uuid", typeId).QueryRow(&messageRes)
		if err != nil {
			return nil, err
		}
		messageOut := messageOutput{}
		messageOut.Title = messageRes.Title
		messageOut.Content = messageRes.Content
		messageOut.OrganizeName = messageRes.OrganizeName
		messageOut.CreateUser = messageRes.CreateUser
		messageOut.ReadCount = readCount
		messageOut.CreateTime =messageRes.CreateTime.Format("2006年01月02日 15:04")
		return messageOut,nil
	}else if theType == 2 {	// 投票
		err := o.Raw("SELECT is_del FROM vote WHERE id=?", typeId).QueryRow(&delId)
		if err != nil {
			return nil, err
		}
		if delId == 2 {
			delMessage := delMessageResponse{}
			err := o.Raw("SELECT member.name as del_user FROM vote join member on vote.del_user = member.openid WHERE vote.id=? and vote.organize_uuid=member.organize_uuid", typeId).QueryRow(&delMessage)
			if err != nil {
				return nil, err
			}
			return delMessage,nil
		}
		isEnd :=0
		isVoteNum := 0
		voteNum := 0
		err = o.Raw("SELECT is_end FROM vote WHERE id=?", typeId).QueryRow(&isEnd)
		if err != nil {
			return nil, err
		}
		err = o.Raw("SELECT count(DISTINCT openid) FROM vote_success WHERE vote_item_id<>0 AND vote_id=?", typeId).QueryRow(&isVoteNum)
		if err != nil {
			return nil, err
		}
		err = o.Raw("SELECT count(DISTINCT openid) FROM vote_success WHERE vote_id=?", typeId).QueryRow(&voteNum)
		if err != nil {
			return nil, err
		}
		// 未截止
		if isEnd==1 {
			voteRes := voteResponse{}
			var voteNotEndTrueContent []string
			err = o.Raw("SELECT title,organize.organize_name as organize_name,member.name as create_user,vote.create_time as create_time,vote.end_time as end_time,vote.max_num as max_num FROM vote left join organize on vote.organize_uuid = organize.uuid left join member on vote.create_user=member.openid WHERE vote.id=? AND vote.organize_uuid=member.organize_uuid", typeId).QueryRow(&voteRes)
			if err != nil {
				return nil, err
			}
			_,err = o.Raw("SELECT content FROM vote_item WHERE vote_id=?", typeId).QueryRows(&voteNotEndTrueContent)
			if err != nil {
				return nil, err
			}
			voteOut := voteOutput1{}
			voteOut.Title = voteRes.Title
			voteOut.IsEnd = isEnd
			voteOut.VoteNum = voteNum
			voteOut.IsVoteNum = isVoteNum
			voteOut.Content = voteNotEndTrueContent
			voteOut.OrganizeName = voteRes.OrganizeName
			voteOut.CreateUser = voteRes.CreateUser
			voteOut.MaxNum = voteRes.MaxNum
			voteOut.CreateTime =voteRes.CreateTime.Format("2006年01月02日 15:04")
			voteOut.EndTime =voteRes.EndTime.Format("2006年01月02日 15:04")
			return voteOut,nil
		}else if isEnd==2 {	// 自动截止
			voteRes := voteResponse{}
			err = o.Raw("SELECT title,organize.organize_name as organize_name,member.name as create_user,vote.create_time as create_time,vote.end_time as end_time FROM vote left join organize on vote.organize_uuid = organize.uuid left join member on vote.create_user=member.openid WHERE vote.id=? AND vote.organize_uuid=member.organize_uuid", typeId).QueryRow(&voteRes)
			if err != nil {
				return nil, err
			}
			voteOut := voteOutput2{}
			voteOut.Title = voteRes.Title
			voteOut.IsEnd = isEnd
			voteOut.VoteNum = voteNum
			voteOut.IsVoteNum = isVoteNum
			voteOut.OrganizeName = voteRes.OrganizeName
			voteOut.CreateUser = voteRes.CreateUser
			voteOut.CreateTime =voteRes.CreateTime.Format("2006年01月02日 15:04")
			voteOut.EndTime =voteRes.EndTime.Format("2006年01月02日 15:04")
			return voteOut,nil
		}else if isEnd==3{ // 手动截止
			voteRes := voteResponse{}
			err = o.Raw("SELECT title,organize.organize_name as organize_name,member.name as create_user,vote.create_time as create_time,vote.end_time as end_time FROM vote left join organize on vote.organize_uuid = organize.uuid left join member on vote.create_user=member.openid WHERE vote.id=? AND vote.organize_uuid=member.organize_uuid", typeId).QueryRow(&voteRes)
			if err != nil {
				return nil, err
			}
			endManualRes := endManualResponse{}
			err = o.Raw("SELECT member.name as manual_user,manual_time from vote join member on vote.manual_user = member.openid  WHERE vote.id = ? AND member.organize_uuid = vote.organize_uuid", typeId).QueryRow(&endManualRes)
			if err != nil {
				return nil, err
			}
			voteOut := voteOutput3{}
			voteOut.Title = voteRes.Title
			voteOut.IsEnd = isEnd
			voteOut.VoteNum = voteNum
			voteOut.IsVoteNum = isVoteNum
			voteOut.ManualUser = endManualRes.ManualUser
			voteOut.ManualTime = endManualRes.ManualTime.Format("2006年01月02日 15:04")
			voteOut.OrganizeName = voteRes.OrganizeName
			voteOut.CreateUser = voteRes.CreateUser
			voteOut.CreateTime =voteRes.CreateTime.Format("2006年01月02日 15:04")
			voteOut.EndTime =voteRes.EndTime.Format("2006年01月02日 15:04")
			return voteOut,nil
		}
	}
	return []string{},nil
}