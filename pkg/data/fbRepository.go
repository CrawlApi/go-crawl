package data

import (
	"encoding/json"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/parnurzeal/gorequest"
)

const (
	ERROR_MSG_API_TIMEOUT = "request api timeout"
)

type FBRepo struct {
	Agent *gorequest.SuperAgent
	Url   string
}

func (r *FBRepo) FetchApi() (string, error) {
	return getApi(r.Agent, r.Url)
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

func (r *FBRepo) parseRawProfile(src string) (models.FBRawProfile, error) {
	var data models.FBRawProfile
	err := json.Unmarshal([]byte(src), &data)
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

func (r *FBRepo) parseRawPosts(src string) (models.FBRawPosts, error) {
	var data models.FBRawPosts
	err := json.Unmarshal([]byte(src), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
