package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"github.com/llitfkitfk/cirkol/pkg/util"
)

func SearchIGProfile(c *gin.Context, ch chan <- result.Profile) {
	userId := c.Param("userId")
	querySrc := c.Query("q")
	middleCh := make(chan result.Profile, 2)

	go SearchIGProfileForName(querySrc, middleCh)
	go SearchIGProfileForId(userId, middleCh)

	for i := 0; i < 2; i++ {
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

func SearchIGProfileForName(userName string, ch chan <- result.Profile) {
	url := "https://www.instagram.com/" + userName + "/"
	var profile result.Profile
	body := GetProfileApi(url, ch)
	profileMat := util.Matcher(REGEX_INSTAGRAM_PROFILE, body)
	var data result.IGNameRawProfile

	if len(profileMat) > 2 {
		profile.RawData = profileMat[1] + profileMat[3]
		ParseProfileJson(profile.RawData, &data, ch)
		profile.MergeIGNameProfile(data)

	} else {
		profile.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
		profile.ErrMessage = ERROR_MSG_REGEX_MISS_MATCHED
	}

	ch <- profile

}

func SearchIGProfileForId(userId string, ch chan <- result.Profile) {
	url := "https://i.instagram.com/api/v1/users/" + userId + "/info/"
	var profile result.Profile
	body := GetProfileApi(url, ch)

	var data result.IGIDRawProfile
	profile.UserId = userId
	profile.Website = url
	profile.RawData = body
	ParseProfileJson(profile.RawData, &data, ch)
	profile.MergeIGIDProfile(data)
	ch <- profile
}

func SearchIGPosts(c *gin.Context, ch chan <- result.Posts) {
	userId := c.Param("userId")
	querySrc := c.Query("q")
	middleCh := make(chan result.Posts, 2)

	go SearchIGPostsForName(querySrc, middleCh)
	go SearchIGPostsForId(userId, middleCh)

	for i := 0; i < 2; i++ {
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

func SearchIGPostsForName(userName string, ch chan result.Posts) {
	url := "https://www.instagram.com/" + userName + "/media/"

	body := GetPostsApi(url, ch)

	var data result.IGNameRawPosts
	ParsePostsJson(body, &data, ch)

	var posts result.Posts
	posts.MergeIGNamePosts(data)
	ch <- posts

}

func SearchIGPostsForId(userId string, ch chan result.Posts) {
	url := "https://i.instagram.com/api/v1/users/" + userId + "/info/"

	body := GetPostsApi(url, ch)

	var data result.IGIDRawProfile
	ParsePostsJson(body, &data, ch)
	userName := data.User.Username

	var posts result.Posts
	SearchIGPostsForRegex(userName, ch, &posts)

	ch <- posts
}

func SearchIGPostsForRegex(userName string, ch chan result.Posts, posts *result.Posts) {
	url := "https://www.instagram.com/" + userName + "/"
	body := GetPostsApi(url, ch)

	postsMat := util.Matcher(REGEX_INSTAGRAM_POSTS, body)
	var data result.IGIDRawPosts
	if len(postsMat) > 1 {

		jsonData := `{ "nodes": ` + postsMat[2] + "]}"

		ParsePostsJson(jsonData, &data, ch)

		posts.MergeIGIdPosts(data)

	} else {
		posts.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
		posts.ErrMessage = ERROR_MSG_REGEX_MISS_MATCHED
	}

}

func SearchIGUID(c *gin.Context, ch chan <- result.UID) {
	rawurl := c.PostForm("url")
	body := GetUidApi(rawurl, ch)

	var uid result.UID
	uid.Url = rawurl
	uid.Media = "ig"
	uid.UserId = MatchStrUidCh(0, REGEX_INSTAGRAM_PROFILE_ID, body, ch)
	uid.Status = true
	ch <- uid
}
