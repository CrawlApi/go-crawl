package client

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
)

func (r *Result) GetFBProfile() (models.Profile, error) {
	var profile models.Profile

	if r.err == nil {
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
	//var rawProfile models.WBRawProfile
	//err := common.ParseJson(result, &rawProfile)
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

func (r *Result) GetIGProfile() (models.Profile, error) {
	//var data models.IGRawProfile
	//err := common.ParseJson(r.getRawProfileStr(body), &data)
	//
	//var profile models.Profile
	//if err != nil {
	//	profile.FetchErr(err)
	//	return profile
	//}
	//profile.ParseIGProfile(data)
	//
	//return profile
	return models.Profile{}, nil
}

//func (r *IGRepo) getRawProfileStr(body string) string {
//	matcher := common.Matcher(REGEX_INSTAGRAM_PROFILE, body)
//	if len(matcher) > 3 {
//		return matcher[1] + matcher[3]
//	}
//	return ""
//}

func (r *Result) GetIGV2Profile() (models.Profile, error) {
	//var profile models.Profile
	//var data models.IGV2RawProfile
	//err := common.ParseJson(body, &data)
	//if err != nil {
	//	profile.FetchErr(err)
	//	return profile
	//}
	//profile.ParseIGV2Profile(data)
	//
	//return profile
	return models.Profile{}, nil
}

func (r *Result) GetWXProfile() (models.Profile, error) {
	return models.Profile{}, nil
}

func (r *Result) GetYTBProfile() (models.Profile, error) {
	return models.Profile{}, nil
}