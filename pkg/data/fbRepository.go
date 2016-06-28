package data

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/parnurzeal/gorequest"
)

const REGEXP_FACEBOOK_PROFILE_ID = `fb://(page|profile|group)/(\d+)`

type FBRepo struct {
	Agent *gorequest.SuperAgent
	Url   string
}

func (r *FBRepo) FetchUIDApi() (string, error) {
	return getApi(r.Agent, r.Url)
}

func (r *FBRepo) FetchProfileApi() (string, error) {
	return getApi(r.Agent, r.Url)
}

func (r *FBRepo) FetchPostsApi() (string, error) {
	return getApi(r.Agent, r.Url)
}

func (r *FBRepo) FetchReactionsApi() (string, error) {
	return getApi(r.Agent, r.Url)
}

func (r *FBRepo) ParseRawUID(body string) models.UID {

	matcher := common.Matcher(REGEXP_FACEBOOK_PROFILE_ID, body)

	var uid models.UID
	uid.Media = "fb"
	if len(matcher) > 2 {
		uid.Type = matcher[1]
		uid.UserId = matcher[2]
		uid.Status = true
	}
	uid.Date = common.Now()
	return uid
}

func (r *FBRepo) ParseRawProfile(body string) models.Profile {
	var rawProfile models.FBRawProfile
	err := common.ParseJson(body, &rawProfile)

	var profile models.Profile
	if err != nil {
		profile.FetchErr(err)
		return profile
	}
	profile.ParseFBProfile(rawProfile)

	return profile
}

func (r *FBRepo) ParseRawPosts(body string) models.Posts {
	var rawPosts models.FBRawPosts
	err := common.ParseJson(body, &rawPosts)
	var posts models.Posts
	if err != nil {
		posts.FetchErr(err)
		return posts
	}
	posts.ParseFBRawPosts(rawPosts)

	return posts
}

func (r *FBRepo) ParseRawReactions(body string)  models.FBReactions {
	var data models.FBRawReactions
	err := common.ParseJson(body, &data)

	var reactions models.FBReactions
	if err != nil {
		reactions.FetchErr(err)
		return reactions
	}
	reactions.ParseFBReactions(data)

	return reactions
}
