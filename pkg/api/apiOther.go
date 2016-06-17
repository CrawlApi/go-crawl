package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"time"
)

func SearchOtherProfile(c *gin.Context, ch chan <- result.Profile) {
	var profile result.Profile
	profile.ErrCode = ERROR_CODE_API_MISS_MATCHED
	profile.ErrMessage = ERROR_MSG_API_MISS_MATCHED
	profile.Date = time.Now().Unix()
	ch <- profile
}
func SearchOtherPosts(c *gin.Context, ch chan <- result.Posts) {
	var posts result.Posts
	posts.ErrCode = ERROR_CODE_API_MISS_MATCHED
	posts.ErrMessage = ERROR_MSG_API_MISS_MATCHED
	posts.Date = time.Now().Unix()
	ch <- posts
}

func SearchOtherUID(c *gin.Context, ch chan <- result.UID) {
	var uid result.UID
	uid.ErrCode = ERROR_CODE_API_MISS_MATCHED
	uid.ErrMessage = ERROR_MSG_API_MISS_MATCHED
	uid.Date = time.Now().Unix()
	ch <- uid
}
