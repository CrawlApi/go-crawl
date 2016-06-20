package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"github.com/llitfkitfk/cirkol/pkg/util"
)

func SearchFBProfile(c *gin.Context, ch chan <- result.Profile) {
	userId := c.Param("userId")

	url := "https://graph.facebook.com/v2.6/" + userId + "?fields=" + PAGE_PROFILE_FIELDS_ENABLE + "&access_token=" + FACEBOOK_TOKEN
	body := GetProfileApi(url, ch)

	var profile result.Profile
	var data result.FBRawProfile
	profile.UserId = userId
	profile.Website = url
	profile.RawData = body
	ParseProfileJson(profile.RawData, &data, ch)
	profile.MergeFBRawProfile(data)
	ch <- profile
}

func SearchFBPosts(c *gin.Context, ch chan <- result.Posts) {
	userId := c.Param("userId")
	limit := c.DefaultQuery("limit", "10")
	url := "https://graph.facebook.com/v2.6/" + userId + "/feed?fields=" + PAGE_FEED_FIELDS_ENABLE + "," + PAGE_FEED_CONNECTIONS + "&limit=" + limit + "&access_token=" + FACEBOOK_TOKEN

	body := GetPostsApi(url, ch)

	var data result.FBRawPosts
	ParsePostsJson(body, &data, ch)

	var posts result.Posts
	posts.MergeFBRawPosts(data)
	ch <- posts
}

func SearchFBUID(c *gin.Context, ch chan <-result.UID) {


	rawurl := c.PostForm("url")
	//name := c.PostForm("name")

	middleCh := make(chan result.UID)
	//go SearchFBUIDForName(name, middleCh)
	go SearchFBUIDForUrl(rawurl, middleCh)

	for i := 0; i < 1; i++ {
		select {
		case item := <-middleCh:
			if i == 1 {
				ch <- item
			} else {
				if item.Status {
					ch <- item
				}
			}

		}
	}

}
func SearchFBUIDForUrl(rawurl string, ch chan result.UID) {
	body := GetUidApi(rawurl, ch)
	matcher := util.Matcher(REGEXP_FACEBOOK_PROFILE_ID, body)

	var uid result.UID
	uid.Url = rawurl
	uid.Media = "fb"
	if len(matcher) > 2 {
		uid.Status = true
		uid.Type = matcher[1]
		uid.UserId = matcher[2]
	} else {
		uid.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
		uid.ErrMessage = ERROR_MSG_REGEX_MISS_MATCHED
	}

	ch <- uid
}
func SearchFBUIDForName(userName string, ch chan result.UID)  {
	//url := "https://graph.facebook.com/v2.6/" + userName + "?fields=id"
	//
	//body := GetUidApi(url, ch)
	//
	//var data result.FBRawUid
	//
	//
	//var uid result.UID
	//ch <- uid
}