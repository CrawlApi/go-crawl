package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/util"
	"time"
	"github.com/llitfkitfk/cirkol/pkg/result"
)

func GetWBUid(c *gin.Context) {
	rawurl := c.PostForm("url")
	uidCh := make(chan result.UID)
	go func() {
		body := GetFBAPI(rawurl)
		matcher := util.Matcher(REGEXP_WEIBO_PROFILE_ID, body)
		var result result.UID
		result.Url = rawurl
		result.Media = "wb"
		if len(matcher) > 2 {
			result.Status = true
			result.UserId = matcher[2]
		} else {
			result.Status = false
		}
		result.Date = time.Now().Unix()

		uidCh <- result
	}()
	Response(uidCh, c)
}

func GetWBProfile(c *gin.Context) {

}

func GetWBPosts(c *gin.Context) {

}
