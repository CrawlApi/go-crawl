package data

import (
	"errors"
	"fmt"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/parnurzeal/gorequest"
	"strings"
)

const (
	URL_WEIBO_PROFILE    = "http://mapi.weibo.com/2/profile?gsid=_&c=&s=&user_domain=%s"
	URL_WEIBO_POSTS      = "http://m.weibo.cn/%s"
	URL_WEIBO_API_POSTS  = "http://m.weibo.cn/page/tpl?containerid=%s_-_WEIBO_SECOND_PROFILE_WEIBO&itemid=&title=全部微博"
	WEIBO_POST_LINK_PREF = "http://m.weibo.cn/%s"
)

const (
	REGEXP_WEIBO_POSTS_ID   = `itemid":"(\d+)`
	REGEXP_WEIBO_POSTS      = `render_data (...+)mod\\/pagelist",(...+)]},'common(...+);</script><script src=`
	REGEXP_WEIBO_PROFILE_ID = `uid=(\d+)`

	REGEXP_WEIBO_POST_LINK = `(http://|)(www.|)weibo.com`
	REGEXP_WEIBO_POST_INFO = `}}},(...+),{"mod_type":`
)

type WBRepo struct {
	Agent  *gorequest.SuperAgent
	UserId string
	RawUrl string
}

func NewWBRepoWithUid(userId string) *WBRepo {
	return &WBRepo{
		Agent:  common.GetAgent(),
		UserId: userId,
	}
}

func NewWBRepoWithUrl(rawUrl string) *WBRepo {
	return &WBRepo{
		Agent:  common.GetAgent(),
		RawUrl: rawUrl,
	}
}

func (r *WBRepo) FetchUIDApi() (string, error) {
	return getApi(r.Agent, r.RawUrl)
}

func (r *WBRepo) FetchProfileApi() (string, error) {
	return getApi(r.Agent, common.UrlString(URL_WEIBO_PROFILE, r.UserId))
}

func (r *WBRepo) FetchPostsApi() (string, error) {
	body, err := getApi(r.Agent, common.UrlString(URL_WEIBO_POSTS, r.UserId))
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

func (r *WBRepo) FetchPostInfo() (string, error) {
	url, err := r.getPostInfoUrl(r.RawUrl)
	if err != nil {
		return "", err
	}
	return getApi(r.Agent, url)
}

func (r *WBRepo) getPostInfoUrl(rawUrl string) (string, error) {
	matcher := common.Matcher(REGEXP_WEIBO_POST_LINK, rawUrl)
	if len(matcher) > 0 {
		i := strings.Index(rawUrl, "com")
		url := fmt.Sprintf(WEIBO_POST_LINK_PREF, rawUrl[i+4:])
		return url, nil
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

func (r *WBRepo) ParsePostInfo(body string) models.Post {

	data := r.parseRawPost(body)
	var post models.Post
	post.ParseWBRawPost(data)
	return post
}

func (r *WBRepo) parseRawPost(body string) models.WBRawPost {
	str := common.GetMatcherValue(1, REGEXP_WEIBO_POST_INFO, body)
	var result models.WBRawPost
	common.ParseJson(str, &result)
	return result
}
