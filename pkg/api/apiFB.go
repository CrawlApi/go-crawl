package api

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/llitfkitfk/cirkol/pkg/util"
)

func GetFBAPI(url string) string {
	_, body, errs := reqClient.Timeout(8 * time.Second).Get(url).Set("accept-language", "en-US").End()
	if errs != nil {
		return ""
	}
	return body
}

// FaceBook Profile
func GetFBProfile(c *gin.Context) {
	userId := c.Param("userId")

	profileCh := make(chan string)

	go func() {
		url := "https://graph.facebook.com/v2.6/" + userId + "?fields=" + PAGE_PROFILE_FIELDS_ENABLE + "&access_token=" + FACEBOOK_TOKEN
		body := GetFBAPI(url)
		profileCh <- body
	}()
	StringResponse(<-profileCh, c)
}

// FaceBook Posts
func GetFBPosts(c *gin.Context) {
	userId := c.Param("userId")
	limit := c.DefaultQuery("limit", "10")
	postCh := make(chan string)
	go func() {
		url := "https://graph.facebook.com/v2.6/" + userId + "/feed?fields=" + PAGE_FEED_FIELDS_ENABLE + "," + PAGE_FEED_CONNECTIONS + "&limit=" + limit + "&access_token=" + FACEBOOK_TOKEN
		body := GetFBAPI(url)
		postCh <- body
	}()

	StringResponse(<-postCh, c)
}

// FaceBook user_id
func GetFBUid(c *gin.Context) {
	rawurl := c.PostForm("url")
	uidCh := make(chan UID)
	go func() {
		body := GetFBAPI(rawurl)
		matcher := util.Matcher(REGEXP_FACEBOOK_PROFILE_ID, body)
		var result UID
		result.Url = rawurl
		result.Media = "fb"
		if len(matcher) > 2 {
			result.Status = true
			result.Type = matcher[1]
			result.UserId = matcher[2]
		} else {
			result.Status = false
		}
		result.Date = time.Now().Unix()

		uidCh <- result
	}()
	Response(uidCh, c)
}
