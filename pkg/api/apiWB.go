package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"log"
	"strings"
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
	userId := c.Param("userId")
	log.Println(userId)
	querySrc := c.Query("q")

	url := "http://m.weibo.cn/d/" + querySrc
	body := GetPostsApi(url, ch)
	userPostsId := MatchStrPostCh(0, REGEXP_WEIBO_POSTS_ID, body, ch)
	log.Printf(userPostsId)
	urlPosts := "http://m.weibo.cn/page/tpl?containerid=" + userPostsId + "_-_WEIBO_SECOND_PROFILE_WEIBO&itemid=&title=全部微博"
	postsRawDataBody := GetPostsApi(urlPosts, ch)

	postsRawData := MatchStrPostCh(1, REGEXP_WEIBO_POSTS, postsRawDataBody, ch)

	postsRawData = "{" + strings.Replace(postsRawData, "(MISSING)", "", -1)
	var data result.WBRawPosts
	ParsePostsJson(postsRawData, &data, ch)
	var posts result.Posts
	posts.MergeWBRawPosts(data)

	ch <- posts

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

