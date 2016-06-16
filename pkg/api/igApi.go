package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"github.com/llitfkitfk/cirkol/pkg/util"
	"encoding/json"
)

func SearchIGPosts(userId string, c *gin.Context, ch chan <- result.Posts) {

}

func SearchIGProfile(userId string, c *gin.Context, ch chan <- result.Profile) {
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
	var data result.IGNameRawProfile

	body, err := ReqApi(url)
	if err != nil {
		profile.ErrCode = ERROR_CODE_API_TIMEOUT
		profile.ErrMessage = err.Error()
	} else {
		profileMat := util.Matcher(REGEX_INSTAGRAM_PROFILE, body)

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
			profile.ErrMessage = "ig regex miss matched"
		}
	}
	ch <- profile

}

func SearchIGProfileForId(userId string, ch chan <-result.Profile) {
	url := "https://i.instagram.com/api/v1/users/" + userId + "/info/"
	var profile result.Profile
	var data result.IGIDRawProfile
	body, err := ReqApi(url)
	if err != nil {
		profile.ErrCode = ERROR_CODE_API_TIMEOUT
		profile.ErrMessage = err.Error()
	} else {
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