package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"github.com/llitfkitfk/cirkol/pkg/util"
	"time"
	"log"
)

func GetFBAPI(url string) string {
	_, body, errs := reqClient.Timeout(8 * time.Second).Get(url).Set("accept-language", "en-US").End()
	if errs != nil {
		log.Println(errs)
		return ""
	}
	return body
}

// FaceBook Profile
func GetFBProfile(c *gin.Context) {
	userId := c.Param("userId")

	profileCh := make(chan result.Profile)

	go func() {
		url := "https://graph.facebook.com/v2.6/" + userId + "?fields=" + PAGE_PROFILE_FIELDS_ENABLE + "&access_token=" + FACEBOOK_TOKEN
		body := GetFBAPI(url)
		var profile result.Profile
		var data result.FBRawProfile
		profile.RawData = body
		err := json.Unmarshal([]byte(profile.RawData), &data)
		if err != nil {
			profile.Status = false
		} else {
			profile.MergeFBRawProfile(data)
		}
		profile.Date = time.Now().Unix()
		profileCh <- profile
	}()
	ProfileResponse(profileCh, c)
}

// FaceBook Posts
func GetFBPosts(c *gin.Context) {
	userId := c.Param("userId")
	limit := c.DefaultQuery("limit", "10")
	postCh := make(chan result.Posts)
	go func() {
		url := "https://graph.facebook.com/v2.6/" + userId + "/feed?fields=" + PAGE_FEED_FIELDS_ENABLE + "," + PAGE_FEED_CONNECTIONS + "&limit=" + limit + "&access_token=" + FACEBOOK_TOKEN
		body := GetFBAPI(url)
		var posts result.Posts
		var data result.FBRawPosts
		posts.RawData = body
		err := json.Unmarshal([]byte(posts.RawData), &data)
		if err != nil {
			posts.Status = false
		} else {
			posts.MergeFBRawPosts(data)
		}
		posts.Date = time.Now().Unix()
		postCh <- posts

	}()

	PostsResponse(postCh, c)
}

// FaceBook user_id
func GetFBUid(c *gin.Context) {
	rawurl := c.PostForm("url")
	uidCh := make(chan result.UID)
	go func() {
		body := GetFBAPI(rawurl)
		matcher := util.Matcher(REGEXP_FACEBOOK_PROFILE_ID, body)
		var result result.UID
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
