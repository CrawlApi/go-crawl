package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"github.com/llitfkitfk/cirkol/pkg/util"
	"encoding/json"
)

func SearchWXProfile(c *gin.Context, ch chan <- result.Profile) {
	userId := c.Param("userId")
	url := "http://weixin.sogou.com/weixin?type=1&query=" + userId + "&ie=utf8&_sug_=n&_sug_type_="
	body := GetApi(url, ch)

	var profile result.Profile
	profile.UserId = userId
	profile.RawData = body
	profile.Website = util.DecodeString(util.MatchString(0, REGEXP_WEIXIN_URL, body))
	profile.Avatar = util.MatchString(0, REGEXP_WEIXIN_LOGO, body)
	profile.About = util.MatchString(1, REGEXP_WEIXIN_FEATURE, body)

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
		urlMat := util.Matcher(REGEXP_WEIXIN_URL, body)
		if len(urlMat) > 0 {
			postBody, err := ReqApi(urlMat[1])
			if err != nil {

			} else {
				postMat := util.Matcher(REGEXP_WEIXIN_POSTS, postBody)
				if len(postMat) > 0 {

					jsonStr := util.DecodeString(postMat[1])
					var data result.WXRawPosts
					err = json.Unmarshal([]byte(jsonStr), &data)
					if err != nil {
						posts.ErrCode = ERROR_CODE_JSON_ERROR
						posts.ErrMessage = err.Error()
					} else {
						posts.MergeWXRawPosts(data)
					}
				} else {
					posts.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
					posts.ErrMessage = ERROR_MSG_REGEX_MISS_MATCHED
				}
			}
		}
	}

	ch <- posts
}

func SearchWXUID(c *gin.Context, ch chan <-result.UID) {
	rawurl := c.PostForm("url")
	var uid result.UID
	body, err := ReqApi(rawurl)
	if err != nil {
		uid.ErrCode = ERROR_CODE_API_TIMEOUT
		uid.ErrMessage = err.Error()
	} else {
		matcher := util.Matcher(REGEXP_WEIXIN_PROFILE_ID, body)

		uid.Url = rawurl
		uid.Media = "wx"
		if len(matcher) > 0 {
			uid.Status = true
			uid.UserId = matcher[1]
		} else {
			uid.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
			uid.ErrMessage = ERROR_MSG_REGEX_MISS_MATCHED
		}
	}
	ch <- uid
}