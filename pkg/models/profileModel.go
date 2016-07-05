package models

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
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
	ErrMessage string `json:"error_message"`
}

func (p *Profile) FetchErr(err error) {
	p.ErrMessage = err.Error()
	p.Date = common.Now()
	p.Status = false
}

func (p *Profile) ParseFBProfile(data FBRawProfile) {
	p.UserId = data.ID
	p.Name = data.Name
	p.Avatar = common.UrlString("https://graph.facebook.com/%s/picture?type=normal", data.ID)

	p.Fans = data.Engagement.Count
	p.About = data.About
	p.Website = data.Link
	p.RawData = common.Interface2String(data)
	p.Date = common.Now()
	p.Status = true
}

func (p *Profile) ParseWBProfile(data WBRawProfile) {
	p.UserId = common.Int2Str(data.UserInfo.ID)
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

func (p *Profile) ParseIGProfile(data IGRawProfile) {
	p.UserId = data.User.Username
	p.Name = data.User.FullName
	p.Avatar = data.User.ProfilePicURL
	p.PostNum = data.User.Media.Count
	p.FollowNum = data.User.Follows.Count
	p.Fans = data.User.FollowedBy.Count

	p.Website = "https://www.instagram.com/" + data.User.Username + "/"
	p.RawData = common.Interface2String(data)
	p.Date = common.Now()
	p.Status = true
}

func (p *Profile) ParseIGV2Profile(data IGV2RawProfile) {
	p.UserId = common.Int2Str(data.User.Pk)
	p.Name = data.User.Username
	p.Avatar = data.User.ProfilePicURL
	p.PostNum = data.User.MediaCount
	p.FollowNum = data.User.FollowingCount
	p.Fans = data.User.FollowerCount
	p.About = data.User.Biography
	p.Website = "https://www.instagram.com/" + data.User.Username + "/"
	p.RawData = common.Interface2String(data)
	p.Date = common.Now()
	p.Status = true
}
