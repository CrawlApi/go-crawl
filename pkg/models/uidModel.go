package models

import "github.com/llitfkitfk/cirkol/pkg/common"


type UID struct {
	Type       string `json:"type"`
	UserId     string `json:"user_id"`
	Date       int64  `json:"date"`
	Url        string `json:"url"`
	Status     bool   `json:"status"`
	Media      string `json:"media"`
	Message    string `json:"message"`
	ErrMessage string `json:"error_message"`
}

func (u *UID) FetchErr(err error) {
	if err != nil {
		u.ErrMessage = err.Error()
	} else {
		u.ErrMessage = common.ERROR_MSG_URL_NOT_MATCHED
	}
	u.Date = common.Now()
	u.Status = false
}