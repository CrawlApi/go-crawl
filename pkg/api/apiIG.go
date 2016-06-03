package api

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/llitfkitfk/cirkol/pkg/util"
)

func GetIGAPI(url string) string {
	_, body, errs := reqClient.Timeout(5 * time.Second).Get(url).End()
	if errs != nil {
		return ""
	}
	return body
}

func GetIGUid(c *gin.Context) {
	rawurl := c.PostForm("url")
	uidCh := make(chan UID)
	go func() {
		body := GetIGAPI(rawurl)
		matcher := util.Matcher(REGEX_INSTAGRAM_PROFILE_ID, body)
		var result UID
		result.Url = rawurl
		result.Media = "ig"
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

func GetIGProfile(c *gin.Context) {
	userId := c.Param("userId")

	profileCh := make(chan string)

	go func() {
		//url := "https://i.instagram.com/api/v1/users/" + userId + "/info/"
		url := "https://api.instagram.com/v1/users/" + userId + "/?access_token=" + INSTAGRAM_TOKEN
		body := GetIGAPI(url)
		profileCh <- body
	}()
	StringResponse(<-profileCh, c)
}

func GetIGPosts(c *gin.Context) {
	userId := c.Param("userId")
	postCh := make(chan string)
	go func() {
		url := "https://api.instagram.com/v1/users/" + userId + "/media/recent?access_token=" + INSTAGRAM_TOKEN
		body := GetIGAPI(url)
		postCh <- body
	}()

	StringResponse(<-postCh, c)
}
