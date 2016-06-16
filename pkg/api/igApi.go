package api

import (
	"encoding/json"
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
			if item.Status {
				ch <- item
			}
		}
	}
}

func SearchIGProfileForName(userName string, ch chan <- result.Profile) {
	url := "https://www.instagram.com/" + userName + "/"
	var profile result.Profile
	body, err := ReqApi(url)
	if err != nil {
		profile.ErrCode = ERROR_CODE_API_TIMEOUT
		profile.ErrMessage = err.Error()
	} else {
		profileMat := util.Matcher(REGEX_INSTAGRAM_PROFILE, body)
		var data result.IGNameRawProfile
		if len(profileMat) > 2 {
			profile.RawData = profileMat[1] + profileMat[3]
			err = json.Unmarshal([]byte(profile.RawData), &data)
			if err != nil {
				profile.ErrCode = ERROR_CODE_JSON_ERROR
				profile.ErrMessage = err.Error()
			} else {
				profile.MergeIGNameProfile(data)

			}
		} else {
			profile.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
			profile.ErrMessage = ERROR_MSG_REGEX_MISS_MATCHED
		}
	}
	ch <- profile

}

func SearchIGProfileForId(userId string, ch chan <- result.Profile) {
	url := "https://i.instagram.com/api/v1/users/" + userId + "/info/"
	var profile result.Profile
	body, err := ReqApi(url)
	if err != nil {
		profile.ErrCode = ERROR_CODE_API_TIMEOUT
		profile.ErrMessage = err.Error()
	} else {
		var data result.IGIDRawProfile
		profile.RawData = body
		err := json.Unmarshal([]byte(profile.RawData), &data)
		if err != nil {
			profile.ErrCode = ERROR_CODE_JSON_ERROR
			profile.ErrMessage = err.Error()
		} else {
			profile.MergeIGIDProfile(data)
		}
	}
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
			if item.Status {
				ch <- item
			}
		}
	}
}

func SearchIGPostsForName(userName string, ch chan result.Posts) {
	url := "https://www.instagram.com/" + userName + "/media/"
	var posts result.Posts

	body, err := ReqApi(url)
	if err != nil {
		posts.ErrCode = ERROR_CODE_API_TIMEOUT
		posts.ErrMessage = err.Error()
	} else {
		var data result.IGNameRawPosts
		err := json.Unmarshal([]byte(body), &data)
		if err != nil {
			posts.ErrCode = ERROR_CODE_JSON_ERROR
			posts.ErrMessage = err.Error()
		} else {
			posts.MergeIGNamePosts(data)
		}
	}
	ch <- posts

}

func SearchIGPostsForId(userId string, ch chan result.Posts) {
	url := "https://i.instagram.com/api/v1/users/" + userId + "/info/"
	var posts result.Posts
	body, err := ReqApi(url)
	if err != nil {
		posts.ErrCode = ERROR_CODE_API_TIMEOUT
		posts.ErrMessage = err.Error()
	} else {
		var data result.IGIDRawProfile
		err := json.Unmarshal([]byte(body), &data)
		if err != nil {
			posts.ErrCode = ERROR_CODE_JSON_ERROR
			posts.ErrMessage = err.Error()
		} else {
			userName := data.User.Username

			SearchIGPostsForRegex(userName, ch, &posts)

		}
	}
	ch <- posts
}

func SearchIGPostsForRegex(userName string, ch chan result.Posts, posts *result.Posts) {
	url := "https://www.instagram.com/" + userName + "/"
	body, err := ReqApi(url)
	if err != nil {
		posts.ErrCode = ERROR_CODE_API_TIMEOUT
		posts.ErrMessage = err.Error()
	} else {
		postsMat := util.Matcher(REGEX_INSTAGRAM_POSTS, body)
		var data result.IGIDRawPosts
		if len(postsMat) > 2 {

			jsonData := `{ "nodes": ` + postsMat[2] + "]}"

			err = json.Unmarshal([]byte(jsonData), &data)
			if err != nil {
				posts.ErrCode = ERROR_CODE_JSON_ERROR
				posts.ErrMessage = err.Error()
			} else {
				posts.MergeIGIdPosts(data)
			}
		} else {
			posts.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
			posts.ErrMessage = ERROR_MSG_REGEX_MISS_MATCHED
		}
	}
}
