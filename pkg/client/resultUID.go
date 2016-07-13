package client

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/llitfkitfk/cirkol/pkg/parser"
)

func (r *Result) GetWBUID() (models.UID, error) {
	matcher := parser.ParseWBUID(r.Body)

	var uid models.UID
	uid.Media = "wb"
	if len(matcher) > 1 {
		uid.UserId = matcher[1]
		uid.Status = true
	}
	uid.Date = common.Now()
	return uid, nil
}
func (r *Result) GetFBUID() (models.UID, error) {
	matcher := parser.ParseFBUIDFromBody(r.Body)
	var uid models.UID
	uid.Media = "fb"
	uid.Date = common.Now()
	if len(matcher) > 2 {
		uid.Type = matcher[1]
		uid.UserId = matcher[2]
		uid.Status = true
		return uid, nil
	}
	return uid, common.MissMatchError()
}


func (r *Result) GetIGUID() (models.UID, error) {
	//matcher := common.Matcher(REGEX_INSTAGRAM_PROFILE_ID, body)
	//
	//var uid models.UID
	//uid.Media = "ig"
	//if len(matcher) > 1 {
	//	uid.UserId = matcher[1]
	//	uid.Status = true
	//}
	//uid.Date = common.Now()
	//return uid
	return models.UID{}, nil
}

func (r *Result) GetIGV2UID() (models.UID, error) {
	//matcher := common.Matcher(REGEX_INSTAGRAM_PROFILE_ID, body)
	//
	//var uid models.UID
	//uid.Media = "ig"
	//if len(matcher) > 1 {
	//	uid.UserId = matcher[1]
	//	uid.Status = true
	//}
	//uid.Date = common.Now()
	//return uid
	return models.UID{}, nil
}

func (r *Result) GetWXUID() (models.UID, error) {
	//matcher := common.Matcher(REGEXP_WEIXIN_PROFILE_ID, body)
	//
	//var uid models.UID
	//uid.Media = "wx"
	//if len(matcher) > 1 {
	//	uid.UserId = matcher[1]
	//	uid.Status = true
	//}
	//uid.Date = common.Now()
	//return uid
	return models.UID{}, nil
}
func (r *Result) GetYTBUID() (models.UID, error) {
	////matcher := common.Matcher(REGEXP_WEIBO_PROFILE_ID, body)
	//
	//var uid models.UID
	////uid.Media = "wb"
	////if len(matcher) > 1 {
	////	uid.UserId = matcher[1]
	////	uid.Status = true
	////}
	////uid.Date = common.Now()
	//return uid
	return models.UID{}, nil
}