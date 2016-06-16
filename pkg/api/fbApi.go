package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"encoding/json"
)

func SearchFBProfile(c *gin.Context, ch chan <- result.Profile) {
	userId := c.Param("userId")
	url := "https://graph.facebook.com/v2.6/" + userId + "?fields=" + PAGE_PROFILE_FIELDS_ENABLE + "&access_token=" + FACEBOOK_TOKEN
	var profile result.Profile

	body, err := ReqApi(url)
	if err != nil {
		profile.ErrCode = ERROR_CODE_API_TIMEOUT
		profile.ErrMessage = err.Error()
	} else {
		var data result.FBRawProfile
		profile.RawData = body
		err = json.Unmarshal([]byte(profile.RawData), &data)
		if err != nil {
			profile.ErrCode = ERROR_CODE_JSON_ERROR
			profile.ErrMessage = err.Error()
		} else {
			profile.MergeFBRawProfile(data)
		}
	}

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