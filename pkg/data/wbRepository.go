package data

import (
	"errors"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/parnurzeal/gorequest"
	"strings"
)

const (
	REGEXP_WEIBO_POSTS_ID   = `itemid":"(\d+)`
	REGEXP_WEIBO_POSTS      = `render_data (...+)mod\\/pagelist",(...+)]},'common(...+);</script><script src=`
	REGEXP_WEIBO_PROFILE_ID = `uid=(\d+)`
)

const (
	URL_WEIBO_API_POSTS = "http://m.weibo.cn/page/tpl?containerid=%s_-_WEIBO_SECOND_PROFILE_WEIBO&itemid=&title=全部微博"
)

type WBRepo struct {
	Agent *gorequest.SuperAgent
	Url   string
}

func (r *WBRepo) FetchUIDApi() (string, error) {
	return getApi(r.Agent, r.Url)
}

func (r *WBRepo) FetchProfileApi() (string, error) {
	return getApi(r.Agent, r.Url)
}

func (r *WBRepo) FetchPostsApi() (string, error) {
	body, err := getApi(r.Agent, r.Url)
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

func (r *WBRepo) getPostsUrl(body string) (string, error) {
	matcher := common.Matcher(REGEXP_WEIBO_POSTS_ID, body)
	if len(matcher) > 1 {
		return common.UrlString(URL_WEIBO_API_POSTS, common.DecodeString(matcher[1])), nil
	}
	return "", errors.New(common.ERROR_MSG_REGEX_MISS_MATCHED)
}

func (r *WBRepo) ParseRawUID(body string) models.UID {
	matcher := common.Matcher(REGEXP_WEIBO_PROFILE_ID, body)

	var uid models.UID
	uid.Media = "wb"
	if len(matcher) > 1 {
		uid.UserId = matcher[1]
		uid.Status = true
	}
	uid.Date = common.Now()
	return uid
}

func (r *WBRepo) ParseRawProfile(body string) models.Profile {
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

func (r *WBRepo) ParseRawPosts(body string) models.Posts {
	var rawPosts models.WBRawPosts

	err := common.ParseJson(r.getPostsStr(body), &rawPosts)

	var posts models.Posts
	if err != nil {
		posts.FetchErr(err)
		return posts
	}
	posts.ParseWBRawPosts(rawPosts)

	return posts
}

func (r *WBRepo) getPostsStr(body string) string {
	matcher := common.Matcher(REGEXP_WEIBO_POSTS, body)
	if len(matcher) > 2 {
		return "{" + strings.Replace(matcher[2], "(MISSING)", "", -1)
	}
	return ""
}
