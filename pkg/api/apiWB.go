package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
)

func SearchWBProfile(c *gin.Context, ch chan <- result.Profile) {
	userId := c.Param("userId")

	url := "http://mapi.weibo.com/2/profile?gsid=_&c=&s=&user_domain=" + userId
	body := GetProfileApi(url, ch)
	var profile result.Profile
	profile.UserId = userId
	profile.Website = url
	profile.RawData = body

	var data result.WBRawProfile
	ParseProfileJson(profile.RawData, &data, ch)
	profile.MergeWBRawProfile(data)
	ch <- profile
}

func SearchWBPosts(c *gin.Context, ch chan <- result.Posts) {

	//url := "https://www.instagram.com/" + "" + "/"
	//var posts result.Posts
	//urlCh := make(chan string)
	//SearchApiData(url, &posts)

	//select {
	//case body := <- urlCh:
	//
	//}
	//postsMat := util.Matcher(REGEX_INSTAGRAM_POSTS, body)
	//if len(postsMat) > 2 {
	//
	//} else {
	//	posts.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
	//	posts.ErrMessage = ERROR_MSG_REGEX_MISS_MATCHED
	//}

}

func SearchWBUID(c *gin.Context, ch chan <-result.UID) {
	rawurl := c.PostForm("url")
	var uid result.UID
	body := GetUidApi(rawurl, ch)

	uid.Url = rawurl
	uid.Media = "wb"
	uid.UserId = MatchStrUidCh(0, REGEXP_WEIBO_PROFILE_ID, body, ch)
	uid.Status = true
	ch <- uid
}

