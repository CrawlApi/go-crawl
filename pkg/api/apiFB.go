package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"encoding/json"
	"github.com/llitfkitfk/cirkol/pkg/util"
)

func SearchFBProfile(c *gin.Context, ch chan <- result.Profile) {
	userId := c.Param("userId")

	url := "https://graph.facebook.com/v2.6/" + userId + "?fields=" + PAGE_PROFILE_FIELDS_ENABLE + "&access_token=" + FACEBOOK_TOKEN
	body := GetProfileApi(url, ch)

	var profile result.Profile
	var data result.FBRawProfile
	profile.UserId = userId
	profile.Website = url
	profile.RawData = body
	ParseProfileJson(profile.RawData, &data, ch)
	profile.MergeFBRawProfile(data)
	ch <- profile
}

func SearchFBPosts(c *gin.Context, ch chan <- result.Posts) {
	userId := c.Param("userId")
	limit := c.DefaultQuery("limit", "10")
	url := "https://graph.facebook.com/v2.6/" + userId + "/feed?fields=" + PAGE_FEED_FIELDS_ENABLE + "," + PAGE_FEED_CONNECTIONS + "&limit=" + limit + "&access_token=" + FACEBOOK_TOKEN
	var posts result.Posts

	body, err := ReqApi(url)
	if err != nil {
		posts.ErrCode = ERROR_CODE_API_TIMEOUT
		posts.ErrMessage = err.Error()
	} else {
		var data result.FBRawPosts
		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			posts.ErrCode = ERROR_CODE_JSON_ERROR
			posts.ErrMessage = err.Error()
		} else {
			posts.MergeFBRawPosts(data)
		}
	}
	ch <- posts
}

func SearchFBUID(c *gin.Context, ch chan <-result.UID) {
	rawurl := c.PostForm("url")
	var uid result.UID
	body, err := ReqApi(rawurl)
	if err != nil {
		uid.ErrCode = ERROR_CODE_API_TIMEOUT
		uid.ErrMessage = err.Error()
	} else {
		matcher := util.Matcher(REGEXP_FACEBOOK_PROFILE_ID, body)
		uid.Url = rawurl
		uid.Media = "fb"
		if len(matcher) > 2 {
			uid.Status = true
			uid.Type = matcher[1]
			uid.UserId = matcher[2]
		} else {
			uid.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
			uid.ErrMessage = ERROR_MSG_REGEX_MISS_MATCHED
		}
	}
	ch <- uid
}