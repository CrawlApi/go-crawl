package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/util"
	"time"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"encoding/json"
)

func GetWBAPI(url string) string {
	_, body, errs := reqClient.Timeout(5 * time.Second).Get(url).End()
	if errs != nil {
		return ""
	}
	return body
}

func GetWBUid(c *gin.Context) {
	rawurl := c.PostForm("url")
	uidCh := make(chan result.UID)
	go func() {
		body := GetFBAPI(rawurl)
		matcher := util.Matcher(REGEXP_WEIBO_PROFILE_ID, body)
		var result result.UID
		result.Url = rawurl
		result.Media = "wb"
		if len(matcher) > 1 {
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

func GetWBProfile(c *gin.Context) {
	userId := c.Param("userId")

	profileCh := make(chan result.Profile)

	go func() {
		url := "http://mapi.weibo.com/2/profile?gsid=_&c=&s=&user_domain=" + userId
		body := GetWBAPI(url)
		var profile result.Profile
		var data result.WBRawProfile
		profile.Website = url
		profile.RawData = body
		err := json.Unmarshal([]byte(profile.RawData), &data)
		if err != nil {
			profile.Status = false
		} else {
			profile.MergeWBRawProfile(data)
		}
		profile.Date = time.Now().Unix()
		profileCh <- profile
	}()
	ProfileResponse(profileCh, c)
}

func GetWBPosts(c *gin.Context) {

}
