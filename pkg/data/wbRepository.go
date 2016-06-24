package data

import (
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/parnurzeal/gorequest"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"errors"
	"strings"
)

const (
	REGEXP_WEIBO_POSTS_ID = `itemid":"(\d+)`
	REGEXP_WEIBO_POSTS = `render_data (...+)mod\\/pagelist",(...+)]},'common(...+);</script><script src=`
)

type WBRepo struct {
	Agent *gorequest.SuperAgent
	Url   string
}

func (r *WBRepo) FetchApi() (string, error) {
	return getApi(r.Agent, r.Url)
}

func (r *WBRepo) ParseRawProfile(body string) models.Profile {

	rawProfile, err := r.parseRawProfile(body)
	var profile models.Profile
	if err != nil {
		profile.FetchErr(err)
		return profile
	}
	profile.ParseWBProfile(rawProfile)

	return profile
}

func (r *WBRepo) parseRawProfile(body string) (models.WBRawProfile, error) {
	var data models.WBRawProfile
	err := common.ParseJson(body, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *WBRepo) ParseRawPosts(body string) models.Posts {
	rawPosts, err := r.parseRawPosts(body)
	var posts models.Posts
	if err != nil {
		posts.FetchErr(err)
		return posts
	}
	posts.ParseWBRawPosts(rawPosts)

	return posts
}
func (r *WBRepo)  getPostsUrl(body string) (string, error) {
	matcher := common.Matcher(REGEXP_WEIBO_POSTS_ID, body)
	if len(matcher) > 1 {
		return "http://m.weibo.cn/page/tpl?containerid=" + common.DecodeString(matcher[1]) + "_-_WEIBO_SECOND_PROFILE_WEIBO&itemid=&title=全部微博", nil
	}
	return "", errors.New(common.ERROR_MSG_REGEX_MISS_MATCHED)
}

func (r *WBRepo) getPostsStr(body string) string {
	matcher := common.Matcher(REGEXP_WEIBO_POSTS, body)
	if len(matcher) > 2 {
		return "{" + strings.Replace(matcher[2], "(MISSING)", "", -1)
	}
	return ""
}

func (r *WBRepo) parseRawPosts(body string) (models.WBRawPosts, error) {
	var data models.WBRawPosts
	urlStr, err := r.getPostsUrl(body)
	if err != nil {
		return data, err
	}

	postsBody, err := getApi(r.Agent, urlStr)
	if err != nil {
		return data, err
	}

	postsStr := r.getPostsStr(postsBody)

	err = common.ParseJson(postsStr, &data)

	if err != nil {
		return data, err
	}
	return data, nil
}
