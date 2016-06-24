package data

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/parnurzeal/gorequest"
)

type IGRepo struct {
	Agent *gorequest.SuperAgent
	Url   string
}

func (r *IGRepo) FetchApi() (string, error) {
	return getApi(r.Agent, r.Url)
}

func (r *IGRepo) ParseRawProfile(body string) models.Profile {

	rawProfile, err := r.parseRawProfile(body)
	var profile models.Profile
	if err != nil {
		profile.FetchErr(err)
		return profile
	}
	profile.ParseFBProfile(rawProfile)

	return profile
}

func (r *IGRepo) parseRawProfile(body string) (models.FBRawProfile, error) {
	var data models.FBRawProfile
	err := common.ParseJson(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *IGRepo) ParseRawPosts(body string) models.Posts {
	rawPosts, err := r.parseRawPosts(body)
	var posts models.Posts
	if err != nil {
		posts.FetchErr(err)
		return posts
	}
	posts.ParseFBRawPosts(rawPosts)

	return posts
}

func (r *IGRepo) parseRawPosts(body string) (models.FBRawPosts, error) {
	var data models.FBRawPosts

	err := common.ParseJson(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
