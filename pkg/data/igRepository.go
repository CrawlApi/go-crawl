package data

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/parnurzeal/gorequest"
)

const REGEX_INSTAGRAM_PROFILE = `ProfilePage": \[([\s\S]+), "nodes": ([\s\S]+)]([\s\S]+)]},`

type IGRepo struct {
	Agent *gorequest.SuperAgent
	Url   string
}

func (r *IGRepo) FetchProfileApi() (string, error) {
	return getApi(r.Agent, r.Url)
}

func (r *IGRepo) FetchPostsApi() (string, error) {
	return getApi(r.Agent, r.Url)
}

func (r *IGRepo) ParseRawProfile(body string) models.Profile {
	var data models.IGRawProfile
	err := common.ParseJson(r.getRawProfileStr(body), &data)

	var profile models.Profile
	if err != nil {
		profile.FetchErr(err)
		return profile
	}
	profile.ParseIGProfile(data)

	return profile
}

func (r *IGRepo) getRawProfileStr(body string) string {
	matcher := common.Matcher(REGEX_INSTAGRAM_PROFILE, body)
	if len(matcher) > 3 {
		return matcher[1] + matcher[3]
	}
	return ""
}

func (r *IGRepo) ParseRawPosts(body string) models.Posts {
	var data models.IGRawPosts
	err := common.ParseJson(body, &data)

	var posts models.Posts
	if err != nil {
		posts.FetchErr(err)
		return posts
	}
	posts.ParseIGRawPosts(data)

	return posts
}