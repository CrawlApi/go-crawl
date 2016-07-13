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
	return models.Profile{}, nil
}

func (r *Result) GetIGV2Profile() (models.Profile, error) {
	return models.Profile{}, nil
}

func (r *Result) GetWXProfile() (models.Profile, error) {
	return models.Profile{}, nil
}

func (r *Result) GetYTBProfile() (models.Profile, error) {
	return models.Profile{}, nil
}