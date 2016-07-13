package parser

import (
	"regexp"
	"strings"
)

const (
	REGEXP_FACEBOOK_PROFILE_ID         = `fb://(page|profile|group)/(\d+)`
	REGEXP_FACEBOOK_POST_ID            = `posts/(\d+)`
	REGEXP_FACEBOOK_POST_CREATED_AT    = `data-utime="(\d+)"`
	REGEXP_FACEBOOK_POST_COMMENT_COUNT = `"commentcount":(\d+)`
	REGEXP_FACEBOOK_POST_LIKE_COUNT    = `"likecount":(\d+)`
	REGEXP_FACEBOOK_POST_SHARE_COUNT   = `"sharecount":(\d+)`
	REGEXP_FACEBOOK_POST_CONTENT       = `<meta name="description" content="(...+)" /><meta name="robots"`
	REGEXP_FACEBOOK_POST_PIC           = `<img class="scaledImageFitWidth img" src="(...+)" alt="`
	REGEXP_FACEBOOK_POST_PERMALINK     = `"permalink":"(...+)","permalinkcommentid"`
)

const (
	REGEXP_WEIBO_POSTS_ID   = `itemid":"(\d+)`
	REGEXP_WEIBO_POSTS      = `render_data (...+)mod\\/pagelist",(...+)]},'common(...+);</script><script src=`
	REGEXP_WEIBO_PROFILE_ID = `uid=(\d+)`

	REGEXP_WEIBO_POST_LINK = `(http://|)(www.|)weibo.com`
	REGEXP_WEIBO_POST_INFO = `}}},(...+),{"mod_type":`
)

const REGEX_URL_TYPE = `(facebook|instagram|weixin|weibo)`

func ParseWBPostUrl(rawUrl string) string {
	matcher := matcher(REGEXP_WEIBO_POST_LINK, rawUrl)
	if len(matcher) > 0 {
		i := strings.Index(rawUrl, "com")
		return rawUrl[i+4:]
	}
	return ""
}

func ParseWBPostsUrl(src string) string {
	return getMatcherValue(1, REGEXP_WEIBO_POSTS_ID, src)
}

func ParseFBUIDFromUrl(url string) string {
	return getMatcherValue(1, `facebook.com/(\S+)/(photos|videos|posts)`, url)
}

func ParseFBUIDFromBody(body string) []string {
	return matcher(REGEXP_FACEBOOK_PROFILE_ID, body)
}

func ParseWBUID(body string) []string {
	return matcher(REGEXP_WEIBO_PROFILE_ID, body)
}

func ParseFBPostSuffId(url string) string {
	switch getMatcherValue(2, `facebook.com/(\S+)/(photos|videos|posts)/`, url) {
	case "videos":
		return getMatcherValue(3, `facebook.com/(\S+)/videos/(...+)/(\d+)`, url)
	case "photos":
		return getMatcherValue(3, `facebook.com/(\S+)/photos/(...+)/(\d+)`, url)
	case "posts":
		return getMatcherValue(2, `facebook.com/(\S+)/posts/(\d+)`, url)
	}
	return ""

}

func getMatcherValue(length int, expr, body string) string {
	matcher := matcher(expr, body)
	if len(matcher) > length {
		return matcher[length]
	}
	return ""
}

func matcher(expr string, s string) []string {
	r, _ := regexp.Compile(expr)
	return r.FindStringSubmatch(s)

}

func CheckUrl(url string) string {
	matcher := matcher(REGEX_URL_TYPE, url)
	if len(matcher) > 0 {
		return matcher[0]
	}
	return ""
}
