package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"github.com/llitfkitfk/cirkol/pkg/util"
	"html"
	"encoding/json"
	"log"
)

func SearchWXProfile(c *gin.Context, ch chan <- result.Profile) {
	userId := c.Param("userId")
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
	} else {
		profile.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
		profile.ErrMessage = ERROR_MSG_REGEX_MISS_MATCHED
	}

	if len(logoMat) > 0 {
		profile.Avatar = logoMat[1]
	}

	if len(featureMat) > 1 {
		profile.About = featureMat[2]
	}
	profile.Status = true

	ch <- profile
}

func SearchWXPosts(c *gin.Context, ch chan <- result.Posts) {
	userId := c.Param("userId")
	url := "http://weixin.sogou.com/weixin?type=1&query=" + userId + "&ie=utf8&_sug_=n&_sug_type_="
	var posts result.Posts
	body, err := ReqApi(url)
	if err != nil {
		posts.ErrCode = ERROR_CODE_API_TIMEOUT
		posts.ErrMessage = err.Error()
	} else {

	}
	urlMat := util.Matcher(REGEXP_WEIXIN_URL, body)
	if len(urlMat) > 0 {
		postBody, err := ReqApi(urlMat[1])
		if err != nil {

		} else {
			postMat := util.Matcher(REGEXP_WEIXIN_POSTS, postBody)
			if len(postMat) > 0 {

				jsonStr := html.UnescapeString(postMat[1])
				var data result.WXRawPosts
				err = json.Unmarshal([]byte(jsonStr), &data)
				if err != nil {
					posts.ErrCode = ERROR_CODE_JSON_ERROR
					posts.ErrMessage = err.Error()
				} else {
					posts.MergeWXRawPosts(data)
				}
			}
		}
	}
	ch <- posts
}

func SearchWXUID(c *gin.Context, ch chan <-result.UID) {
	rawurl := c.PostForm("url")
	log.Println(rawurl)

}