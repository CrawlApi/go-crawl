package data

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/parnurzeal/gorequest"
)

const (
	URL_INSTAGRAM_PROFILE_V2 = "https://i.instagram.com/api/v1/users/%s/info/"
	URL_INSTAGRAM_POSTS_V2   = "https://i.instagram.com/api/v1/users/%s/info/"
)

const (
	URL_INSTAGRAM_API_POSTS    = "https://www.instagram.com/%s/"
	REGEX_INSTAGRAM_POSTS      = `ProfilePage": \[([\s\S]+), "nodes": ([\s\S]+)]([\s\S]+)]},`
	REGEX_INSTAGRAM_PROFILE_ID = `"owner": {"id": "(\d+)`
)

type IGV2Repo struct {
	Agent  *gorequest.SuperAgent
	UserId string
	RawUrl string
}

func NewIGV2RepoWithUid(userId string) *IGV2Repo {
	return &IGV2Repo{
		Agent:  common.GetAgent(),
		UserId: userId,
	}
}

func NewIGV2RepoWithUrl(rawUrl string) *IGV2Repo {
	return &IGV2Repo{
		Agent:  common.GetAgent(),
		RawUrl: rawUrl,
	}
}

func (r *IGV2Repo) FetchUIDApi() (string, error) {
	return getApi(r.Agent, r.RawUrl)
}

func (r *IGV2Repo) FetchProfileApi() (string, error) {
	return getApi(r.Agent, common.UrlString(URL_INSTAGRAM_PROFILE_V2, r.UserId))
}

func (r *IGV2Repo) FetchPostsApi() (string, error) {
	body, err := getApi(r.Agent, common.UrlString(URL_INSTAGRAM_POSTS_V2, r.UserId))
	if err != nil {
		return body, err
	}
	urlStr, err := r.getPostsUrl(body)
	if err != nil {
		return urlStr, err
	}
	postsBody, err := getApi(r.Agent, urlStr)
	if err != nil {
		return postsBody, err
	}
	return postsBody, nil
}

func (r *IGV2Repo) getPostsUrl(body string) (string, error) {
	var data models.IGV2RawProfile
	err := common.ParseJson(body, &data)
	postsUrl := common.UrlString(URL_INSTAGRAM_API_POSTS, data.User.Username)
	if err != nil {
		return postsUrl, err
	}
	return postsUrl, nil
}

func (r *IGV2Repo) ParseRawUID(body string) models.UID {
	matcher := common.Matcher(REGEX_INSTAGRAM_PROFILE_ID, body)

	var uid models.UID
	uid.Media = "ig"
	if len(matcher) > 1 {
		uid.UserId = matcher[1]
		uid.Status = true
	}
	uid.Date = common.Now()
	return uid
}

func (r *IGV2Repo) ParseRawProfile(body string) models.Profile {

	var profile models.Profile
	var data models.IGV2RawProfile
	err := common.ParseJson(body, &data)
	if err != nil {
		profile.FetchErr(err)
		return profile
	}
	profile.ParseIGV2Profile(data)

	return profile
}

func (r *IGV2Repo) ParseRawPosts(body string) models.Posts {
	var data models.IGV2RawPosts

	err := common.ParseJson(r.getRawPostsStr(body), &data)

	var posts models.Posts
	if err != nil {
		posts.FetchErr(err)
		return posts
	}
	posts.ParseIGV2RawPosts(data)

	return posts
}

func (r *IGV2Repo) getRawPostsStr(body string) string {
	matcher := common.Matcher(REGEX_INSTAGRAM_POSTS, body)
	if len(matcher) > 2 {
		return `{ "nodes": ` + matcher[2] + "]}"
	}
	return ""
}
