package models

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/util"
)

type Profile struct {
	UserId     string `json:"user_id"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	PostNum    int    `json:"post_num"`
	FollowNum  int    `json:"follow_num"`
	Fans       int    `json:"fans"`
	Birthday   string `json:"birthday"`
	Website    string `json:"website"`
	About      string `json:"about"`
	RawData    string `json:"raw_data"`
	Status     bool   `json:"status"`
	Date       int64  `json:"date"`
	ErrCode    int    `json:"error_code"`
	ErrMessage string `json:"error_message"`
}

func (p *Profile) FetchErr(err error) {
	p.ErrCode = common.ERROR_CODE_API_FETCH
	p.ErrMessage = err.Error()
	p.Date = common.Now()
	p.Status = false
}

func (p *Profile) ParseFBProfile(data FBRawProfile) {
	p.UserId = data.ID
	p.Name = data.Name
	p.Avatar = data.Cover.Source

	p.Fans = data.Engagement.Count
	p.About = data.About
	p.Website = data.Link
	p.RawData = common.Interface2String(data)
	p.Date = common.Now()
	p.Status = true
}

func (p *Profile) ParseWBProfile(data WBRawProfile) {
	p.UserId = util.Int2Str(data.UserInfo.ID)
	p.Name = data.UserInfo.Name
	p.Avatar = data.UserInfo.ProfileImageURL
	p.PostNum = data.UserInfo.StatusesCount
	p.FollowNum = data.UserInfo.FriendsCount
	p.Fans = data.UserInfo.FollowersCount
	p.About = data.UserInfo.Description

	p.RawData = common.Interface2String(data)
	p.Date = common.Now()
	p.Status = true
}

func (p *Profile) ParseWXProfile(data WXRawProfile) {
	p.UserId = data.UserId
	p.Name = data.Name
	p.Avatar = data.Avatar
	p.About = data.About
	p.Website = data.Website
	p.RawData = common.Interface2String(data)
	p.Date = common.Now()
	p.Status = true
}
