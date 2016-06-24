package models

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"time"
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
	p.Date = time.Now().Unix()
	p.Status = false
}

func (p *Profile) ParseFBProfile(data FBRawProfile) {
	p.UserId = data.ID
	p.Name = data.Name
	p.Avatar = data.Cover.Source

	p.Fans = data.Engagement.Count
	p.About = data.About
	p.Website = data.Link
	p.Date = time.Now().Unix()
	p.Status = true
}

func (p *Profile) ParseWXProfile(data WXRawProfile) {
	p.UserId = data.UserId
	p.Name = data.Name
	p.Avatar = data.Avatar
	p.About = data.About
	p.Website = data.Website
	p.Date = time.Now().Unix()
	p.Status = true
}
