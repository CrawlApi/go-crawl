package data

import (
	"errors"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/parnurzeal/gorequest"
)

const (
	// WEIXIN CONST
	REGEXP_WEIXIN_PROFILE_ID = `微信号: (\S+)</p>`
	REGEXP_WEIXIN_LOGO = `src="((http://img01.sogoucdn.com/app/a)\S+)"`
	REGEXP_WEIXIN_NAME = `<h3>(\S+)</h3>`
	REGEXP_WEIXIN_FEATURE = `功能介绍(...+)class="sp-txt">(...+)</span>`
	REGEXP_WEIXIN_URL = `href="((http://mp.weixin.qq.com/profile)\S+)"`
	REGEXP_WEIXIN_POSTS = `var msgList = '(\S+)';`
)

type WXRepo struct {
	Agent          *gorequest.SuperAgent
	Url            string
	ProfileRawData string
}

func (r *WXRepo) FetchUIDApi() (string, error) {
	return getApi(r.Agent, r.Url)
}

func (r *WXRepo) FetchProfileApi() (string, error) {
	return getApi(r.Agent, r.Url)
}

func (r *WXRepo) FetchPostsApi() (string, error) {
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

func (r *WXRepo) ParseRawUID(body string) models.UID {
	matcher := common.Matcher(REGEXP_WEIXIN_PROFILE_ID, body)

	var uid models.UID
	uid.Media = "wx"
	if len(matcher) > 1 {
		uid.UserId = matcher[1]
		uid.Status = true
	}
	uid.Date = common.Now()
	return uid
}

func (r *WXRepo) ParseRawProfile(body string) models.Profile {
	rawProfile, _ := r.parseRawProfile(body)

	var profile models.Profile
	profile.ParseWXProfile(rawProfile)
	return profile
}

func (r *WXRepo) parseRawProfile(body string) (models.WXRawProfile, error) {
	var data models.WXRawProfile

	data = r.getWXRawProfile(body)

	return data, nil
}

func (r *WXRepo) getWXRawProfile(body string) models.WXRawProfile {
	var data models.WXRawProfile
	data.Name = r.getName(body)
	data.Website = r.getWebsite(body)
	data.Avatar = r.getAvatar(body)
	data.About = r.getAbout(body)
	return data
}

func (r *WXRepo) getName(body string) string {
	matcher := common.Matcher(REGEXP_WEIXIN_NAME, body)
	if len(matcher) > 1 {
		return matcher[1]
	}
	return ""
}

func (r *WXRepo) getWebsite(body string) string {
	matcher := common.Matcher(REGEXP_WEIXIN_URL, body)
	if len(matcher) > 1 {
		return common.DecodeString(matcher[1])
	}
	return ""
}

func (r *WXRepo) getAvatar(body string) string {
	matcher := common.Matcher(REGEXP_WEIXIN_LOGO, body)
	if len(matcher) > 1 {
		return matcher[1]
	}
	return ""
}

func (r *WXRepo) getAbout(body string) string {
	matcher := common.Matcher(REGEXP_WEIXIN_FEATURE, body)
	if len(matcher) > 2 {
		return matcher[2]
	}
	return ""
}

func (r *WXRepo) ParseRawPosts(body string) models.Posts {
	rawPosts, err := r.parseRawPosts(body)
	var posts models.Posts
	if err != nil {
		posts.FetchErr(err)
		return posts
	}
	posts.ParseWXRawPosts(rawPosts)

	return posts
}

func (r *WXRepo) getPostsUrl(body string) (string, error) {
	matcher := common.Matcher(REGEXP_WEIXIN_URL, body)
	if len(matcher) > 1 {
		return matcher[1], nil
	}
	return "", errors.New(common.ERROR_MSG_REGEX_MISS_MATCHED)
}

func (r *WXRepo) getPostsStr(body string) string {
	matcher := common.Matcher(REGEXP_WEIXIN_POSTS, body)
	if len(matcher) > 1 {
		return common.DecodeString(matcher[1])
	}
	return ""
}

func (r *WXRepo) parseRawPosts(body string) (models.WXRawPosts, error) {
	var data models.WXRawPosts

	err := common.ParseJson(r.getPostsStr(body), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
