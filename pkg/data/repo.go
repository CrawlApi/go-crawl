package data

import (
	"github.com/llitfkitfk/cirkol/pkg/client"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
)

var Agent *client.Client

func GR() *client.Client {
	return Agent
}

func FetchProfileErr(err error) models.Profile {
	return models.Profile{
		ErrMessage: err.Error(),
		Date:       common.Now(),
		Status:     false,
	}
}

func FetchPostsErr(err error) models.Posts {
	return models.Posts{
		ErrMessage: err.Error(),
		Date:       common.Now(),
		Status:     false,
	}
}

func FetchPostErr(err error) models.Post {
	return models.Post{
		ErrMessage: err.Error(),
		Date:       common.Now(),
		Status:     false,
	}
}

func FetchUIDErr(err error) models.UID {
	return models.UID{
		ErrMessage: err.Error(),
		Date:       common.Now(),
		Status:     false,
	}
}

type Profile interface {
	FetchProfileApi() client.Result
	ParseRawProfile(client.Result) models.Profile
}

type Posts interface {
	FetchPostsApi() client.Result
	ParseRawPosts(client.Result) models.Posts
}

type Post interface {
	FetchPostInfo() client.Result
	ParsePostInfo(client.Result) models.Post
}

type UID interface {
	FetchUIDApi() client.Result
	ParseRawUID(client.Result) models.UID
}
