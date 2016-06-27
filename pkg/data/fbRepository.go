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

	rawProfile, err := r.parseRawProfile(body)
	var profile models.Profile
	if err != nil {
		profile.FetchErr(err)
		return profile
	}
	profile.ParseFBProfile(rawProfile)

	return profile
}

func (r *FBRepo) parseRawProfile(body string) (models.FBRawProfile, error) {
	var data models.FBRawProfile
	err := common.ParseJson(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *FBRepo) ParseRawPosts(body string) models.Posts {
	rawPosts, err := r.parseRawPosts(body)
	var posts models.Posts
	if err != nil {
		posts.FetchErr(err)
		return posts
	}
	posts.ParseFBRawPosts(rawPosts)

	return posts
}

func (r *FBRepo) parseRawPosts(body string) (models.FBRawPosts, error) {
	var data models.FBRawPosts

	err := common.ParseJson(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
