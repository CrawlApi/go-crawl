package client

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/llitfkitfk/cirkol/pkg/parser"
)

func (r *Result) GetWBUID() (models.UID, error) {
	var uid models.UID
	if r.err != nil {
		return uid, r.err
	}

	matcher := parser.ParseWBUID(r.Body)
	uid.Media = "wb"
	uid.Date = common.Now()
	if len(matcher) > 1 {
		uid.UserId = matcher[1]
		uid.Status = true
		return uid, nil
	}

	return uid, common.MissMatchError()
}

func (r *Result) GetFBUID() (models.UID, error) {
	var uid models.UID
	if r.err != nil {
		return uid, r.err
	}

	matcher := parser.ParseFBUIDFromBody(r.Body)
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
	var uid models.UID
	if r.err != nil {
		return uid, r.err
	}

	matcher := parser.ParseIGUID(r.Body)
	uid.Media = "ig"
	uid.Date = common.Now()
	if len(matcher) > 1 {
		uid.UserId = matcher[1]
		uid.Status = true
		return uid, nil
	}
	return uid, common.MissMatchError()
}

func (r *Result) GetIGV2UID() (models.UID, error) {
	var uid models.UID
	if r.err != nil {
		return uid, r.err
	}

	matcher := parser.ParseIGV2UID(r.Body)
	uid.Media = "ig"
	uid.Date = common.Now()
	if len(matcher) > 1 {
		uid.UserId = matcher[1]
		uid.Status = true
		return uid, nil
	}

	return uid, common.MissMatchError()
}

func (r *Result) GetWXUID() (models.UID, error) {
	var uid models.UID
	if r.err != nil {
		return uid, r.err
	}

	matcher := parser.ParseWXUID(r.Body)
	uid.Media = "wx"
	uid.Date = common.Now()
	if len(matcher) > 1 {
		uid.UserId = matcher[1]
		uid.Status = true
		return uid, nil
	}
	return uid, common.MissMatchError()
}
func (r *Result) GetYTBUID() (models.UID, error) {
	return models.UID{}, nil
}
