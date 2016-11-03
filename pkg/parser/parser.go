package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"regexp"
	"strings"
)

const (
	// facebook
	regexp_fb_profile_id         = `fb://(page|profile|group)/(\d+)`
	regexp_fb_post_id            = `posts/(\d+)`
	regexp_fb_post_create        = `data-utime="(\d+)"`
	regexp_fb_post_comment_count = `"commentcount":(\d+)`
	regexp_fb_post_like_count    = `"likecount":(\d+)`
	regexp_fb_post_share_count   = `"sharecount":(\d+)`
	regexp_fb_post_content       = `<meta name="description" content="(...+)" /><meta name="robots"`
	regexp_fb_post_pic           = `<img class="scaledImageFitWidth img" src="(...+)" alt="`
	regexp_fb_post_permalink     = `"permalink":"(...+)","permalinkcommentid"`

	// weibo
	regexp_wb_posts_id   = `itemid":"(\d+)`
	regexp_wb_posts      = `render_data (...+)mod\\/pagelist",(...+)]},'common(...+);</script><script src=`
	regexp_wb_profile_id = `uid=(\d+)`
	regexp_wb_post_link  = `(http://|)(www.|)weibo.com`
	regexp_wb_post_info  = `}}},(...+),{"mod_type":`

	// instagram
	regexp_ig_profile    = `ProfilePage": \[([\s\S]+), "nodes": ([\s\S]+)]([\s\S]+)]},`
	regexp_ig_post_info  = `_sharedData =(...+);</script>`
	regexp_ig_posts      = `ProfilePage": \[([\s\S]+), "nodes": ([\s\S]+)]([\s\S]+)]},`
	regexp_ig_profile_id = `"owner": {"id": "(\d+)`

	// wechat
	regexp_wx_profile_id = `微信号: (\S+)</p>`
	regexp_wx_logo       = `src="((http://img01.sogoucdn.com/app/a)\S+)"`
	regexp_wx_name       = `<h3>(\S+)</h3>`
	regexp_wx_feature    = `功能介绍(...+)class="sp-txt">(...+)</span>`
	regexp_wx_url        = `href="((http://mp.weixin.qq.com/profile)\S+)"`
	regexp_wx_posts      = `var msgList = '(\S+)';`

	// youtube
	regexp_ytb_profile_id = `微信号: (\S+)</p>`
	// url
	regexp_url_type = `(facebook|instagram|weixin|weibo|youtube)`
)

// weibo
func ParseWBUID(body string) []string {
	return matcher(regexp_wb_profile_id, body)
}

func ParseWBPostUrl(rawUrl string) string {
	matcher := matcher(regexp_wb_post_link, rawUrl)
	if len(matcher) > 0 {
		i := strings.Index(rawUrl, "com")
		return rawUrl[i+4:]
	}
	return ""
}

func ParseWBPostsStr(body string) string {
	matcher := matcher(regexp_wb_posts, body)
	if len(matcher) > 2 {
		return "{" + strings.Replace(matcher[2], "(MISSING)", "", -1)
	}
	return ""
}

func ParseWBPostsUrl(src string) string {
	return getMatcherValue(1, regexp_wb_posts_id, src)
}

func ParseWBPostStr(body string) string {
	return getMatcherValue(1, regexp_wb_post_info, body)
}

// facebook
func ParseFBUIDFromUrl(url string) string {
	return getMatcherValue(1, `facebook.com/(\S+)/(photos|videos|posts)`, url)
}

func ParseFBUIDFromBody(body string) []string {
	return matcher(regexp_fb_profile_id, body)
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

func ParseYTBUID(body string) string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(body))

	var uid string
	var exists bool

	if !exists {
		uid, exists = doc.Selection.Find("#watch7-user-header").Find("a").Attr("href")
	}

	if !exists {
		uid, exists = doc.Selection.Find("#appbar-nav").Find("a").Attr("href")
	}

	if !exists {
		uid, exists = doc.Selection.Find("#qualified-channel-title-text").Find("a").Attr("href")
	}

	common.Log.Info(uid)
	if len(uid) > 6 {
		return uid[6:]
	}
	return uid
}

// wechat

func ParseWXUID(body string) []string {
	return matcher(regexp_wx_profile_id, body)
}

func ParseWXName(body string) string {
	return getMatcherValue(1, regexp_wx_name, body)
}

func ParseWXWeb(body string) string {
	return getMatcherValue(1, regexp_wx_url, body)
}

func ParseWXAvatar(body string) string {
	return getMatcherValue(1, regexp_wx_logo, body)
}

func ParseWXAbout(body string) string {
	return getMatcherValue(2, regexp_wx_feature, body)
}

func ParseWXPostsStr(body string) string {
	return getMatcherValue(1, regexp_wx_posts, body)
}

func ParseWXPostsUrl(body string) string {
	return getMatcherValue(1, regexp_wx_url, body)
}

// instagram
func ParseIGProfile(body string) string {
	matcher := matcher(regexp_ig_profile, body)
	if len(matcher) > 3 {
		return matcher[1] + matcher[3]
	}
	return ""
}

func ParseIGUID(body string) []string {
	return matcher(regexp_ig_profile_id, body)
}

func ParseIGPostStr(body string) string {
	return getMatcherValue(1, regexp_ig_post_info, body)
}

func ParseIGV2UID(body string) []string {
	return matcher(regexp_ig_profile_id, body)
}

func ParseIGV2PostsStr(body string) string {
	matcher := matcher(regexp_ig_posts, body)
	if len(matcher) > 2 {
		return `{ "nodes": ` + matcher[2] + "]}"
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
	matcher := matcher(regexp_url_type, url)
	if len(matcher) > 0 {
		return matcher[0]
	}
	return ""
}
