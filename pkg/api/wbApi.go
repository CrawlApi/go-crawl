package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"encoding/json"
)

func SearchWBProfile(userId string, c *gin.Context, ch chan <- result.Profile) {
	url := "http://mapi.weibo.com/2/profile?gsid=_&c=&s=&user_domain=" + userId
	var profile result.Profile
	var data result.WBRawProfile

	body, err := ReqApi(url)
	if err != nil {
		profile.ErrCode = ERROR_CODE_API_TIMEOUT
		profile.ErrMessage = err.Error()
	} else {
		profile.Website = url
		profile.RawData = body
		err = json.Unmarshal([]byte(profile.RawData), &data)
		if err != nil {
			profile.ErrCode = ERROR_CODE_JSON_ERROR
			profile.ErrMessage = err.Error()
		} else {
			profile.MergeWBRawProfile(data)
		}
	}
	ch <- profile
}