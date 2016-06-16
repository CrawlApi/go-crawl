package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"github.com/llitfkitfk/cirkol/pkg/util"
)

func SearchWXProfile(userId string, c *gin.Context, ch chan<- result.Profile) {
	url := "http://weixin.sogou.com/weixin?type=1&query=" + userId + "&ie=utf8&_sug_=n&_sug_type_="
	var profile result.Profile
	body, err := ReqApi(url)
	if err != nil {
		profile.ErrCode = ERROR_CODE_API_TIMEOUT
		profile.ErrMessage = err.Error()
	}
	logoMat := util.Matcher(REGEXP_WEIXIN_LOGO, body)
	featureMat := util.Matcher(REGEXP_WEIXIN_FEATURE, body)
	urlMat := util.Matcher(REGEXP_WEIXIN_URL, body)

	profile.UserId = userId

	if len(urlMat) > 0 {
		profile.Website = urlMat[1]
	}

	if len(logoMat) > 0 {
		profile.Avatar = logoMat[1]
	}

	if len(featureMat) > 1 {
		profile.About = featureMat[2]
	}

	ch <- profile
}

func SearchWXPosts(userId string, c *gin.Context, ch chan<- result.Posts) {
	//url := "http://weixin.sogou.com/weixin?type=1&query=" + userId + "&ie=utf8&_sug_=n&_sug_type_="
	//body, err := ReqApi(url)
	//
	//urlMat := util.Matcher(REGEXP_WEIXIN_URL, body)
	//if len(urlMat) > 0 {
	//	postBody := ReqApi(urlMat[1])
	//	postMat := util.Matcher(REGEXP_WEIXIN_POSTS, postBody)
	//	if len(postMat) > 0 {
	//
	//		ch <- html.UnescapeString(postMat[1])
	//	}
	//}
}
