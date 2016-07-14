package client

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
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
	//rawProfile, _ := r.parseRawProfile(body)
	//
	//var profile models.Profile
	//profile.ParseWXProfile(rawProfile)
	//return profile
	return models.Profile{}, nil

}

//func (r *WXRepo) parseRawProfile(body string) (models.WXRawProfile, error) {
//	var data models.WXRawProfile
//	data.Name = r.getName(body)
//	data.Website = r.getWebsite(body)
//	data.Avatar = r.getAvatar(body)
//	data.About = r.getAbout(body)
//	return data, nil
//}
//
//func (r *WXRepo) getName(body string) string {
//	matcher := common.Matcher(REGEXP_WEIXIN_NAME, body)
//	if len(matcher) > 1 {
//		return matcher[1]
//	}
//	return ""
//}
//
//func (r *WXRepo) getWebsite(body string) string {
//	matcher := common.Matcher(REGEXP_WEIXIN_URL, body)
//	if len(matcher) > 1 {
//		return common.DecodeString(matcher[1])
//	}
//	return ""
//}
//
//func (r *WXRepo) getAvatar(body string) string {
//	matcher := common.Matcher(REGEXP_WEIXIN_LOGO, body)
//	if len(matcher) > 1 {
//		return matcher[1]
//	}
//	return ""
//}
//
//func (r *WXRepo) getAbout(body string) string {
//	matcher := common.Matcher(REGEXP_WEIXIN_FEATURE, body)
//	if len(matcher) > 2 {
//		return matcher[2]
//	}
//	return ""
//}

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
