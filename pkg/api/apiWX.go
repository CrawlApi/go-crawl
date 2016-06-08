package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/util"
	"time"
	"html"
	"github.com/llitfkitfk/cirkol/pkg/result"
)

func GetWXAPI(url string) string {
	_, body, errs := reqClient.Timeout(5 * time.Second).Get(url).End()
	if errs != nil {
		return ""
	}
	return body
}

func GetWXUid(c *gin.Context) {
	rawurl := c.PostForm("url")
	uidCh := make(chan result.UID)
	go func() {
		body := GetFBAPI(rawurl)
		matcher := util.Matcher(REGEXP_WEIXIN_PROFILE_ID, body)
		var result result.UID
		result.Url = rawurl
		result.Media = "wx"
		if len(matcher) > 0 {
			result.Status = true
			result.UserId = matcher[1]
		} else {
			result.Status = false
		}
		result.Date = time.Now().Unix()

		uidCh <- result
	}()
	Response(uidCh, c)
}

func GetWXProfile(c *gin.Context) {

	userId := c.Param("userId")

	profileCh := make(chan result.Profile)

	go func() {
		url := "http://weixin.sogou.com/weixin?type=1&query=" + userId + "&ie=utf8&_sug_=n&_sug_type_="
		body := GetWXAPI(url)
		logoMat := util.Matcher(REGEXP_WEIXIN_LOGO, body)
		featureMat := util.Matcher(REGEXP_WEIXIN_FEATURE, body)
		urlMat := util.Matcher(REGEXP_WEIXIN_URL, body)
		var result result.Profile
		result.UserId = userId

		if len(urlMat) > 0 {
			result.Website = urlMat[1]
		}

		if len(logoMat) > 0 {
			result.Avatar = logoMat[1]
		}

		if len(featureMat) > 1 {
			result.About = featureMat[2]
		}

		profileCh <- result
	}()
	ProfileResponse(profileCh, c)

}

func GetWXPosts(c *gin.Context) {
	userId := c.Param("userId")
	postCh := make(chan string)
	go func() {

		url := "http://weixin.sogou.com/weixin?type=1&query=" + userId + "&ie=utf8&_sug_=n&_sug_type_="
		body := GetWXAPI(url)
		urlMat := util.Matcher(REGEXP_WEIXIN_URL, body)
		if len(urlMat) > 0 {
			postBody := GetWXAPI(urlMat[1])
			postMat := util.Matcher(REGEXP_WEIXIN_POSTS, postBody)
			if len(postMat) > 0 {

				postCh <- html.UnescapeString(postMat[1])
			}
		}
	}()

	StringResponse(<-postCh, c)
}
