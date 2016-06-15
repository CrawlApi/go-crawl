package api

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/llitfkitfk/cirkol/pkg/util"
	"encoding/json"
	"github.com/llitfkitfk/cirkol/pkg/result"
)

func GetIGAPI(url string) string {
	_, body, errs := reqClient.Timeout(8 * time.Second).Get(url).End()
	if errs != nil {
		return ""
	}
	return body
}

func GetIGUid(c *gin.Context) {
	rawurl := c.PostForm("url")
	uidCh := make(chan result.UID)
	go func() {
		body := GetIGAPI(rawurl)
		matcher := util.Matcher(REGEX_INSTAGRAM_PROFILE_ID, body)
		var uid result.UID
		uid.Url = rawurl
		uid.Media = "ig"
		if len(matcher) > 0 {
			uid.Status = true
			uid.UserId = matcher[1]
		} else {
			uid.Status = false
		}
		uid.Date = time.Now().Unix()
		uidCh <- uid
	}()

	Response(uidCh, c)
}

func GetIGProfile(c *gin.Context) {
	timer := time.After(8 * time.Second)
	//userId := c.Param("userId")
	profileCh := make(chan result.Profile)
	//querySrc := c.Query("q")
	//go getIgProfileFromName(querySrc, profileCh)
	//go getIgProfileFromId(userId, profileCh)
	ProfileResponse2(profileCh, c, timer)
}





func GetIGPosts(c *gin.Context) {
	userId := c.Param("userId")
	postCh := make(chan result.Posts)
	querySrc := c.Query("q")
	go getIgPostFromName(querySrc, postCh)
	go getIgPostFromId(userId, postCh)
	PostsResponse(postCh, c)
}

func getIgPostFromName(userName string, ch chan result.Posts) {
	url := "https://www.instagram.com/" + userName + "/media/"
	body := GetIGAPI(url)
	var posts result.Posts
	var data result.IGNameRawPosts
	posts.RawData = body
	err := json.Unmarshal([]byte(posts.RawData), &data)
	if err != nil {
		posts.Status = false
	} else {
		posts.MergeIGNamePosts(data)
	}
	posts.Date = time.Now().Unix()
	ch <- posts

}

func getIgPostFromId(userId string, ch chan result.Posts) {
	url := "https://i.instagram.com/api/v1/users/" + userId + "/info/"
	body := GetIGAPI(url)
	var data result.IGIDRawProfile
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		var result result.Posts
		result.Status = false
		ch <- result
	} else {
		userName := data.User.Username

		getIgPostFromRegex(userName, ch)
	}

}

func getIgPostFromRegex(userId string, ch chan result.Posts) {
	url := "https://www.instagram.com/" + userId + "/"
	body := GetIGAPI(url)
	postsMat := util.Matcher(REGEX_INSTAGRAM_POSTS, body)

	var posts result.Posts
	var data result.IGIDRawPosts
	if len(postsMat) > 2 {

		posts.RawData = `{ "nodes": ` + postsMat[2] + "]}"

		err := json.Unmarshal([]byte(posts.RawData), &data)
		if err != nil {
			posts.Status = false
		} else {
			posts.MergeIGIdPosts(data)
		}
	} else {
		posts.Status = false
	}

	ch <- posts

}
