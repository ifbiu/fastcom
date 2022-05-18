package message

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

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
	VoteResult []voteResultFull `json:"voteResult"`
}

type voteResultFull struct {
	Content string `json:"content"`
	Num int `json:"num"`
	PercentageNum float64 `json:"percentageNum"`
}


func VoteResult(typeId int) (interface{},error) {
	o := orm.NewOrm()
	var voteItemIds []int
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
	voteOut.CreateTime = voteRes.CreateTime.Format("2006年01月02日 15:04")
	voteOut.EndTime = voteRes.EndTime.Format("2006年01月02日 15:04")
	voteOut.VoteResult = voteResultAll
	return voteOut,nil
}