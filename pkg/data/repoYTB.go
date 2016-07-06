package data

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/parnurzeal/gorequest"
)

const (
	URL_YOUTUBE_PROFILE = "https://www.youtube.com/user/%s/about"
	URL_YOUTUBE_POSTS   = "https://www.youtube.com/user/%s/videos"
)

type YTBRepo struct {
	Agent  *gorequest.SuperAgent
	UserId string
	RawUrl string
}

func NewYTBRepoWithUid(userId string) *YTBRepo {
	return &YTBRepo{
		Agent:  common.GetAgent(),
		UserId: userId,
	}
}

func NewYTBRepoWithUrl(rawUrl string) *YTBRepo {
	return &YTBRepo{
		Agent:  common.GetAgent(),
		RawUrl: rawUrl,
	}
}

func (r *YTBRepo) FetchUIDApi() (string, error) {
	return getApi(r.Agent, r.RawUrl)
}

func (r *YTBRepo) FetchProfileApi() (string, error) {
	return getApi(r.Agent, common.UrlString(URL_YOUTUBE_PROFILE, r.UserId))
}

func (r *YTBRepo) FetchPostsApi() (string, error) {
	body, err := getApi(r.Agent, common.UrlString(URL_YOUTUBE_POSTS, r.UserId))
	if err != nil {
		return body, err
	}

	return body, nil

}

func (r *YTBRepo) ParseRawUID(body string) models.UID {
	//matcher := common.Matcher(REGEXP_WEIBO_PROFILE_ID, body)

	var uid models.UID
	//uid.Media = "wb"
	//if len(matcher) > 1 {
	//	uid.UserId = matcher[1]
	//	uid.Status = true
	//}
	//uid.Date = common.Now()
	return uid
}

func (r *YTBRepo) ParseRawProfile(body string) models.Profile {
	var rawProfile models.WBRawProfile

	err := common.ParseJson(body, &rawProfile)

	var profile models.Profile
	if err != nil {
		profile.FetchErr(err)
		return profile
	}
	profile.ParseWBProfile(rawProfile)

	return profile
}

func (r *YTBRepo) ParseRawPosts(body string) models.Posts {
	var rawPosts models.WBRawPosts

	err := common.ParseJson(body, &rawPosts)

	var posts models.Posts
	if err != nil {
		posts.FetchErr(err)
		return posts
	}
	posts.ParseWBRawPosts(rawPosts)

	return posts
}
