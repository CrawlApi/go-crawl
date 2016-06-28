package data

import (
	"errors"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/parnurzeal/gorequest"
	"time"
)

type Profile interface {
	FetchProfileApi() (string, error)
	ParseRawProfile(string) models.Profile
}

type Posts interface {
	FetchPostsApi() (string, error)
	ParseRawPosts(string) models.Posts
}

type UID interface {
	FetchUIDApi() (string, error)
	ParseRawUID(string) models.UID
}

func getApi(agent *gorequest.SuperAgent, url string) (string, error) {
	_, body, errs := agent.Timeout(10 * time.Second).Set("accept-language", "en-US").Get(url).End()
	if errs != nil {
		return body, errors.New(common.ERROR_MSG_API_TIMEOUT)
	}
	return body, nil
}
