package client

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/llitfkitfk/cirkol/pkg/parser"
)

func (r *Result) GetFBProfile() (models.Profile, error) {
	var profile models.Profile

	if r.err != nil {
		return profile, r.err
	}

	var rawData models.FBRawProfile
	err := common.ParseJson(r.Body, &rawData)
	if err != nil {
		return profile, err
	}

	profile.ParseFBProfile(rawData)

	return profile, nil
}

func (r *Result) GetWBProfile() (models.Profile, error) {
	var profile models.Profile
	if r.err != nil {
		return profile, r.err
	}

	var data models.WBRawProfile
	err := common.ParseJson(r.Body, &data)
	if err != nil {
		return profile, err
	}

	profile.ParseWBProfile(data)

	return profile, nil
}

func (r *Result) GetIGProfile() (models.Profile, error) {
	var profile models.Profile
	if r.err != nil {
		return profile, r.err
	}

	var data models.IGRawProfile
	dStr := parser.ParseIGProfile(r.Body)
	err := common.ParseJson(dStr, &data)
	if err != nil {
		return profile, err
	}

	profile.ParseIGProfile(data)

	return profile, nil
}

func (r *Result) GetIGV2Profile() (models.Profile, error) {
	var profile models.Profile
	if r.err != nil {
		return profile, r.err
	}

	var data models.IGV2RawProfile
	err := common.ParseJson(r.Body, &data)
	if err != nil {
		return profile, err
	}
	profile.ParseIGV2Profile(data)

	return profile, nil
}

func (r *Result) GetWXProfile() (models.Profile, error) {
	var profile models.Profile
	if r.err != nil {
		return profile, r.err
	}

	data, _ := r.parseWXRawProfile(r.Body)

	profile.ParseWXProfile(data)
	return profile, nil
}

func (r *Result) parseWXRawProfile(body string) (models.WXRawProfile, error) {
	var data models.WXRawProfile
	data.Name =  parser.ParseWXName(body)
	data.Website = common.DecodeString(parser.ParseWXWeb(body))
	data.Avatar = parser.ParseWXAvatar(body)
	data.About = parser.ParseWXAbout(body)
	return data, nil
}

func (r *Result) GetYTBProfile() (models.Profile, error) {
	//var rawProfile models.WBRawProfile
	//
	//err := common.ParseJson(body, &rawProfile)
	//
	//var profile models.Profile
	//if err != nil {
	//	profile.FetchErr(err)
	//	return profile
	//}
	//profile.ParseWBProfile(rawProfile)
	//
	//return profile
	return models.Profile{}, nil
}
