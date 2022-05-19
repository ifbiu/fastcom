package message

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
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

type approveResponse1 struct {
	Image string `json:"image"`
	NickName string `json:"nick_name"`
	OrganizeName string `json:"organize_name"`
}
type approveResponse2 struct {
	Image string `json:"image"`
	NickName string `json:"nick_name"`
	ApproveUser string `json:"approve_user"`
	IsApprove int `json:"is_approve"`
	OrganizeName string `json:"organize_name"`
}

type voteOutput1 struct {
	Title string `json:"title"`
	Content []string `json:"content"`
	VoteType int `json:"type"`
	IsEnd int `json:"isEnd"`
	MaxNum int `json:"maxNum"`
	VoteNum int `json:"voteNum"`
	IsVoteNum int `json:"isVoteNum"`
	IsAbstained int `json:"isAbstained"`
	OrganizeName string `json:"organizeName"`
	CreateUser string `json:"createUser"`
	CreateTime string `json:"createTime"`
	EndTime string `json:"endTime"`
}
type voteOutput2 struct {
	Title string `json:"title"`
	VoteType int `json:"type"`
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
	VoteType int `json:"type"`
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

type approveOutput1 struct {
	Image string `json:"image"`
	NickName string `json:"nickName"`
	OrganizeName string `json:"organizeName"`
}

type approveOutput2 struct {
	Image string `json:"image"`
	NickName string `json:"nickName"`
	ApproveUser string `json:"approveUser"`
	IsApprove int `json:"isApprove"`
	OrganizeName string `json:"organizeName"`
}

type voteResultResponse struct {
	Title string `json:"title"`
	OrganizeName string `json:"organizeName"`
	CreateUser string `json:"createUser"`
	CreateTime time.Time `json:"createTime"`
	EndTime time.Time `json:"endTime"`
}

type voteResultOutput struct {
	Title string `json:"title"`
	OrganizeName string `json:"organizeName"`
	CreateUser string `json:"createUser"`
	CreateTime string `json:"createTime"`
	EndTime string `json:"endTime"`
	VoteNumY int `json:"voteNumY"`
	VotePercentageY float64 `json:"votePercentageY"`
	VoteNumN int `json:"voteNumN"`
	VotePercentageN float64 `json:"votePercentageN"`
	VoteNumA int `json:"voteNumA"`
	VotePercentageA float64 `json:"votePercentageA"`
	VoteResult []voteResultFull `json:"voteResult"`
}

type voteResultFull struct {
	Content string `json:"content"`
	Num int `json:"num"`
	PercentageNum float64 `json:"percentageNum"`
}

func GetMessageInfo(theType int,typeId int,openId string) (interface{},error) {
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
		countVote := 0
		isVoteNum := 0
		voteNum := 0
		err = o.Raw("SELECT is_end FROM vote WHERE id=?", typeId).QueryRow(&isEnd)
		if err != nil {
			return nil, err
		}
		err = o.Raw("SELECT count(id) FROM vote_success WHERE vote_id=? AND openid = ? AND vote_item_id != 0", typeId,openId).QueryRow(&countVote)
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
			if countVote == 0 {	// 未投
				voteRes := voteResponse{}
				var voteNotEndTrueContent []string
				err = o.Raw("SELECT title,organize.organize_name as organize_name,member.name as create_user,vote.create_time as create_time,vote.end_time as end_time,vote.max_num as max_num,vote.is_abstained as is_abstained FROM vote left join organize on vote.organize_uuid = organize.uuid left join member on vote.create_user=member.openid WHERE vote.id=? AND vote.organize_uuid=member.organize_uuid", typeId).QueryRow(&voteRes)
				if err != nil {
					return nil, err
				}
				_,err = o.Raw("SELECT content FROM vote_item WHERE vote_id=?", typeId).QueryRows(&voteNotEndTrueContent)
				if err != nil {
					return nil, err
				}
				voteOut := voteOutput1{}
				voteOut.Title = voteRes.Title
				voteOut.VoteType = countVote
				voteOut.IsEnd = isEnd
				voteOut.VoteNum = voteNum
				voteOut.IsVoteNum = isVoteNum
				voteOut.Content = voteNotEndTrueContent
				voteOut.OrganizeName = voteRes.OrganizeName
				voteOut.CreateUser = voteRes.CreateUser
				voteOut.IsAbstained = voteRes.IsAbstained
				voteOut.MaxNum = voteRes.MaxNum
				voteOut.CreateTime =voteRes.CreateTime.Format("2006年01月02日 15:04")
				voteOut.EndTime =voteRes.EndTime.Format("2006年01月02日 15:04")
				return voteOut,nil
			}else{	// 已投
				voteRes := voteResponse{}
				err = o.Raw("SELECT title,organize.organize_name as organize_name,member.name as create_user,vote.create_time as create_time,vote.end_time as end_time FROM vote left join organize on vote.organize_uuid = organize.uuid left join member on vote.create_user=member.openid WHERE vote.id=? AND vote.organize_uuid=member.organize_uuid", typeId).QueryRow(&voteRes)
				if err != nil {
					return nil, err
				}
				voteOut := voteOutput2{}
				voteOut.Title = voteRes.Title
				voteOut.VoteType = countVote
				voteOut.IsEnd = isEnd
				voteOut.VoteNum = voteNum
				voteOut.IsVoteNum = isVoteNum
				voteOut.OrganizeName = voteRes.OrganizeName
				voteOut.CreateUser = voteRes.CreateUser
				voteOut.CreateTime =voteRes.CreateTime.Format("2006年01月02日 15:04")
				voteOut.EndTime =voteRes.EndTime.Format("2006年01月02日 15:04")
				return voteOut,nil
			}
		}else if isEnd==2 {	// 自动截止
			voteRes := voteResponse{}
			err = o.Raw("SELECT title,organize.organize_name as organize_name,member.name as create_user,vote.create_time as create_time,vote.end_time as end_time FROM vote left join organize on vote.organize_uuid = organize.uuid left join member on vote.create_user=member.openid WHERE vote.id=? AND vote.organize_uuid=member.organize_uuid", typeId).QueryRow(&voteRes)
			if err != nil {
				return nil, err
			}
			voteOut := voteOutput2{}
			voteOut.Title = voteRes.Title
			voteOut.VoteType = countVote
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
			voteOut.VoteType = countVote
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
	}else if theType == 3 { // 审核
		isApprove := 0

		err := o.Raw("SELECT is_approve FROM approve WHERE id=?", typeId).QueryRow(&isApprove)
		if err != nil {
			return nil, err
		}

		// 未审核
		if isApprove==0 {
			approveRes := approveResponse1{}
			approveOut := approveOutput1{}
			err = o.Raw("SELECT user.image as image,user.nick_name as nick_name,organize.organize_name as organize_name FROM approve join user on approve.start_user = user.openid join organize on approve.organize_uuid = organize.uuid WHERE approve.id=?", typeId).QueryRow(&approveRes)
			if err != nil {
				return nil, err
			}
			approveOut.Image = approveRes.Image
			approveOut.NickName = approveRes.NickName
			approveOut.OrganizeName = approveRes.OrganizeName
			return approveOut,nil
		}else if isApprove==1 || isApprove==2 {	// 已审核通过
			approveRes := approveResponse2{}
			approveOut := approveOutput2{}
			err = o.Raw("SELECT user.image as image,user.nick_name as nick_name,organize.organize_name as organize_name,approve.is_approve as is_approve,approve.approve_user as approve_user FROM approve join user on approve.start_user = user.openid join organize on approve.organize_uuid = organize.uuid WHERE approve.id=?", typeId).QueryRow(&approveRes)
			if err != nil {
				return nil, err
			}
			var approveUser string
			err := o.Raw("SELECT name FROM member WHERE openid=? AND organize_uuid=(SELECT organize_uuid FROM approve WHERE id=?)", approveRes.ApproveUser,typeId).QueryRow(&approveUser)
			if err != nil {
				return nil, err
			}
			approveOut.Image = approveRes.Image
			approveOut.NickName = approveRes.NickName
			approveOut.ApproveUser = approveUser
			approveOut.IsApprove = approveRes.IsApprove
			approveOut.OrganizeName = approveRes.OrganizeName
			return approveOut,nil
		}
	}else if theType == 4{
		var voteItemIds []int
		var voteNumY,voteNumN,voteNumA,voteNumAll int
		var votePercentageY,votePercentageN,votePercentageA float64
		var voteNumNArr []string
		var voteNumAArr []string
		voteRes := voteResultResponse{}
		voteOut := voteResultOutput{}
		err := o.Raw("SELECT title,organize.organize_name as organize_name,member.name as create_user,vote.create_time as create_time,vote.end_time as end_time FROM vote left join organize on vote.organize_uuid = organize.uuid left join member on vote.create_user=member.openid WHERE vote.id=? AND vote.organize_uuid=member.organize_uuid", typeId).QueryRow(&voteRes)
		if err != nil {
			return nil, err
		}
		_, err = o.Raw("SELECT id FROM vote_item WHERE vote_id=?",typeId).QueryRows(&voteItemIds)
		if err != nil {
			return nil,err
		}
		if len(voteItemIds)==0 {
			return nil,errors.New("投票项空值")
		}
		// 总人数
		err = o.Raw("SELECT count(DISTINCT openid) FROM vote_success WHERE vote_id=?", typeId).QueryRow(&voteNumAll)
		if err != nil {
			return nil, err
		}
		// 已投票人数
		err = o.Raw("SELECT count(DISTINCT openid) FROM vote_success WHERE vote_id=? AND vote_item_id<>0", typeId).QueryRow(&voteNumY)
		if err != nil {
			return nil, err
		}
		// 未投票人数
		_,err = o.Raw("SELECT openid FROM vote_success WHERE vote_id = ? AND vote_item_id = 0 GROUP BY openid HAVING count( openid )=(select count(DISTINCT serial_id) from vote_success where vote_id=?)", typeId,typeId).QueryRows(&voteNumNArr)
		if err != nil {
			return nil, err
		}
		voteNumN = len(voteNumNArr)

		// 弃票人数
		_,err = o.Raw("SELECT openid FROM vote_success WHERE vote_id = ? AND vote_item_id = 2 GROUP BY openid HAVING count( openid )=(select count(DISTINCT serial_id) from vote_success where vote_id=?)", typeId,typeId).QueryRows(&voteNumAArr)
		if err != nil {
			return nil, err
		}
		voteNumA = len(voteNumAArr)

		// 百分比
		votePercentageY, err = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(voteNumY) / float64(voteNumAll)), 64)
		if err != nil {
			return nil,err
		}
		votePercentageN, err = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(voteNumN) / float64(voteNumAll)), 64)
		if err != nil {
			return nil,err
		}
		votePercentageA, err = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(voteNumA) / float64(voteNumAll)), 64)
		if err != nil {
			return nil,err
		}

		voteResultAll := make([]voteResultFull,len(voteItemIds))
		for i, id := range voteItemIds {
			voteResult := voteResultFull{}
			err := o.Raw("SELECT vote_item.content as content,vote_result.vote_num as num,vote_result.vote_percentage as percentage_num FROM vote_result JOIN vote_item on vote_result.vote_item_id=vote_item.id WHERE vote_result.vote_item_id=?", id).QueryRow(&voteResult)
			if err != nil {
				return nil, err
			}
			voteResultAll[i] = voteResult
		}
		voteOut.Title = voteRes.Title
		voteOut.OrganizeName = voteRes.OrganizeName
		voteOut.CreateUser = voteRes.CreateUser
		voteOut.VoteNumY = voteNumY
		voteOut.VoteNumN = voteNumN
		voteOut.VoteNumA = voteNumA
		voteOut.VotePercentageY = votePercentageY
		voteOut.VotePercentageN = votePercentageN
		voteOut.VotePercentageA = votePercentageA
		voteOut.CreateTime = voteRes.CreateTime.Format("2006年01月02日 15:04")
		voteOut.EndTime = voteRes.EndTime.Format("2006年01月02日 15:04")
		voteOut.VoteResult = voteResultAll
		return voteOut,nil
	}

	return []string{},nil
}