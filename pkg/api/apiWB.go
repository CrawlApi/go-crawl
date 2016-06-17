package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"github.com/llitfkitfk/cirkol/pkg/util"
)

func SearchWBProfile(c *gin.Context, ch chan <- result.Profile) {
	userId := c.Param("userId")
	var profile result.Profile
	url := "http://mapi.weibo.com/2/profile?gsid=_&c=&s=&user_domain=" + userId
	body := GetApi(url, ch)
	profile.Website = url
	profile.RawData = body
	var data result.WBRawProfile
	ParseJson(profile.RawData, &data, ch)
	profile.MergeWBRawProfile(data)
	ch <- profile
}

func SearchWBPosts(c *gin.Context, ch chan <- result.Posts) {

	//url := "https://www.instagram.com/" + "" + "/"
	//var posts result.Posts
	//urlCh := make(chan string)
	//SearchApiData(url, &posts)

	//select {
	//case body := <- urlCh:
	//
	//}
	//postsMat := util.Matcher(REGEX_INSTAGRAM_POSTS, body)
	//if len(postsMat) > 2 {
	//
	//} else {
	//	posts.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
	//	posts.ErrMessage = ERROR_MSG_REGEX_MISS_MATCHED
	//}

}

func SearchWBUID(c *gin.Context, ch chan <-result.UID) {
	rawurl := c.PostForm("url")
	var uid result.UID
	body, err := ReqApi(rawurl)
	if err != nil {
		uid.ErrCode = ERROR_CODE_API_TIMEOUT
		uid.ErrMessage = err.Error()
	} else {
		matcher := util.Matcher(REGEXP_WEIBO_PROFILE_ID, body)
		uid.Url = rawurl
		uid.Media = "wb"
		if len(matcher) > 1 {
			uid.Status = true
			uid.UserId = matcher[1]
		} else {
			uid.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
			uid.ErrMessage = ERROR_MSG_REGEX_MISS_MATCHED
		}
	}
	ch <- uid
}

