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
	REGEXP_WEIXIN_LOGO       = `src="((http://img01.sogoucdn.com/app/a)\S+)"`
	REGEXP_WEIXIN_NAME       = `<h3>(\S+)</h3>`
	REGEXP_WEIXIN_FEATURE    = `功能介绍(...+)class="sp-txt">(...+)</span>`
	REGEXP_WEIXIN_URL        = `href="((http://mp.weixin.qq.com/profile)\S+)"`
	REGEXP_WEIXIN_POSTS      = `var msgList = '(\S+)';`
)

type WXRepo struct {
	Agent          *gorequest.SuperAgent
	Url            string
	UserId         string
	ProfileRawData string
}

func (r *WXRepo) FetchApi() (string, error) {
	return getApi(r.Agent, r.Url)
}

func (r *WXRepo) ParseRawProfile(body string) models.Profile {
	rawProfile, err := r.parseRawProfile(body)
	var profile models.Profile
	if err != nil {
		profile.FetchErr(err)
		return profile
	}
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
	data.UserId = r.UserId
	data.RawData = body
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
	return "", errors.New(common.ERROR_MSG_JSON_ERROR)
}
func (r *WXRepo) getPostsStr(body string) string {
	matcher := common.Matcher(REGEXP_WEIXIN_POSTS, body)
	if len(matcher) > 1 {
		return common.DecodeString(matcher[1])
	}
	return ""
}
func (r *WXRepo) parseRawPosts(body string) (models.WXRawPosts, error) {
	urlStr, err := r.getPostsUrl(body)
	var data models.WXRawPosts
	if err != nil {
		return data, errors.New(common.ERROR_MSG_REGEX_MISS_MATCHED)
	}

	postsBody, err := getApi(r.Agent, urlStr)
	if err != nil {
		return data, errors.New(common.ERROR_MSG_WX_POSTS_API_FETCH)
	}

	postsStr := r.getPostsStr(postsBody)
	err = common.ParseJson(postsStr, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
