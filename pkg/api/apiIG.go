package api

import (
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"net/url"
	"regexp"
	"time"
	"github.com/llitfkitfk/cirkol/pkg/util"
)

func GetIGAPI(url string) string {
	_, body, errs := gorequest.New().Timeout(5 * time.Second).Get(url).End()
	if errs != nil {
		return ""
	}
	return body
}

func IGResponse(body string, c *gin.Context) {
	if len(body) > 0 {
		c.String(http.StatusOK, body)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Request Instagram API Failed",
			"status":  false,
			"date": time.Now().Unix(),
		})
	}
}

func GetIGUid(c *gin.Context) {
	rawurl := c.Query("url")
	url, err := url.Parse(rawurl)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "not real url",
			"user_id": nil,
		})
	}
	uidCh := make(chan string)
	go func() {
		body := GetIGAPI(url.String())
		r, _ := regexp.Compile(REGEX_INSTAGRAM_PROFILE_ID)
		matcher := r.FindStringSubmatch(body)
		if len(matcher) > 0 {
			logCh <- matcher
			uidCh <- `{"profile": {"user_id": "` + matcher[1] + `"}, "date": ` + util.ToString(time.Now().Unix()) + `}`
		} else {
			uidCh <- ""
		}
	}()

	IGResponse(<-uidCh, c)
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
	IGResponse(<-profileCh, c)
}

func GetIGPosts(c *gin.Context) {
	userId := c.Param("userId")
	postCh := make(chan string)
	go func() {
		url := "https://api.instagram.com/v1/users/" + userId + "/media/recent?access_token=" + INSTAGRAM_TOKEN
		body := GetIGAPI(url)
		postCh <- body
	}()

	IGResponse(<-postCh, c)
}
