package controllers

import (
	"time"
	"github.com/llitfkitfk/cirkol/pkg/models"
)

const (
	ERROR_CODE_API_FETCH = 4001
	ERROR_CODE_API_TIMEOUT = 4002
	//ERROR_CODE_API_MISS_MATCHED = 4001

	//ERROR_CODE_JSON_ERROR = 4003
	//ERROR_CODE_TIMEOUT = 4004
	//ERROR_CODE_REGEX_MISS_MATCHED = 4005
	//ERROR_CODE_URL_TYPE_NOT_FOUND = 4006

	//ERROR_MSG_API_FETCH = "request api timeout"
	//ERROR_MSG_API_MISS_MATCHED = "no api matched"
	ERROR_MSG_API_TIMEOUT = "request api timeout"
	//ERROR_MSG_JSON_ERROR = "json parse error"
	//ERROR_MSG_TIMEOUT = "request timeout"
	//ERROR_MSG_REGEX_MISS_MATCHED = "regex miss matched"
	//ERROR_MSG_URL_MISS_MATCHED = "url miss matched"
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
	p.ErrCode = ERROR_CODE_API_FETCH
	p.ErrMessage = err.Error()
	p.Date = time.Now().Unix()
	p.Status = false
}

func TimeOutProfile() Profile {
	var p Profile
	p.ErrCode = ERROR_CODE_API_TIMEOUT
	p.ErrMessage = ERROR_MSG_API_TIMEOUT
	p.Date = time.Now().Unix()
	p.Status = false

	return p
}

func (p *Profile) ParseRawProfile(data models.FBRawProfile) {
	p.UserId = data.ID
	p.Name = data.Name
	p.Avatar = data.Cover.Source

	p.Fans = data.Engagement.Count
	p.About = data.About
	p.Website = data.Link
	p.Date = time.Now().Unix()
	p.Status = true
}


