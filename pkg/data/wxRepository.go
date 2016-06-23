package data

import (
	"github.com/parnurzeal/gorequest"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"errors"
	"encoding/json"
)

type WXRepo struct {
	Agent *gorequest.SuperAgent
}

func (r *WXRepo) GetApi(url string) *gorequest.SuperAgent {
	return r.Agent.Timeout(8 * time.Second).Set("accept-language", "en-US").Get(url)
}

func (r *WXRepo) GetRawProfile(c *gin.Context) (models.WXRawProfile, error) {
	userId := c.Param("userId")
	url := "https://graph.facebook.com/v2.6/" + userId + "?fields=" + PAGE_PROFILE_FIELDS_ENABLE + "&access_token=" + common.GetFBToken()
	_, body, errs := r.GetApi(url).End()
	var rawProfile models.WXRawProfile

	if errs != nil {
		return rawProfile, errors.New(ERROR_MSG_API_TIMEOUT)
	}
	err := json.Unmarshal([]byte(body), &rawProfile)
	if err != nil {
		return rawProfile, err
	}
	return rawProfile, nil
}

func (r *WXRepo) GetRawPosts(c *gin.Context) (models.WXRawPosts, error) {
	userId := c.Param("userId")
	limit := c.DefaultQuery("limit", "10")
	url := "https://graph.facebook.com/v2.6/" + userId + "/feed?fields=" + PAGE_FEED_FIELDS_ENABLE + "," + PAGE_FEED_CONNECTIONS + "&limit=" + limit + "&access_token=" + common.GetFBToken()
	_, body, errs := r.GetApi(url).End()
	var rawPosts models.WXRawPosts

	if errs != nil {
		return rawPosts, errors.New(ERROR_MSG_API_TIMEOUT)
	}
	err := json.Unmarshal([]byte(body), &rawPosts)
	if err != nil {
		return rawPosts, err
	}
	return rawPosts, nil
}