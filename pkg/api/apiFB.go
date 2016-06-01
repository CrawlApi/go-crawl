package api

import (
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
	"github.com/llitfkitfk/cirkol/pkg/util"
)

func FBResponse(body string, c *gin.Context) {
	if len(body) > 0 {
		c.String(http.StatusOK, body)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Request Facebook API Failed",
			"status":  false,
			"date": time.Now().Unix(),
		})
	}
}

func GetFBAPI(url string) string {
	_, body, errs := gorequest.New().Timeout(5 * time.Second).Get(url).Set("accept-language", "en-US").End()
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
	FBResponse(<-profileCh, c)
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

	FBResponse(<-postCh, c)
}

// FaceBook user_id
func GetFBUid(c *gin.Context) {
	rawurl := c.Query("url")
	realUrl := util.Matcher(REGEXP_URI, rawurl)
	if len(realUrl) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "not real url",
			"status": false,
			"date": time.Now().Unix(),
		})
	} else {
		uidCh := make(chan string)
		go func() {
			body := GetFBAPI(rawurl)
			logCh <- body
			matcher := util.Matcher(REGEXP_FACEBOOK_PROFILE_ID, body)
			if len(matcher) > 2 {
				logCh <- matcher
				uidCh <- `{"profile": {"type": "` + matcher[1] + `", "user_id": "` + matcher[2] + `"}, "date": ` + util.ToString(time.Now().Unix()) + `}`
			} else {
				uidCh <- ""
			}
		}()
		FBResponse(<-uidCh, c)
	}

}
