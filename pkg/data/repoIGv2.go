package data

import (
	"github.com/llitfkitfk/cirkol/pkg/client"
	"github.com/llitfkitfk/cirkol/pkg/models"
)

type IGV2Repo struct {
	UserId string
	RawUrl string
}

func NewIGV2RepoWithUid(userId string) *IGV2Repo {
	return &IGV2Repo{
		UserId: userId,
	}
}

func NewIGV2RepoWithUrl(rawUrl string) *IGV2Repo {
	return &IGV2Repo{
		RawUrl: rawUrl,
	}
}

func (r *IGV2Repo) FetchUIDApi() client.Result {
	return GR().GetIGV2UIDResult(r.RawUrl)
}

func (r *IGV2Repo) FetchProfileApi() client.Result {
	return GR().GetIGV2ProfileResult(r.UserId)
}

func (r *IGV2Repo) FetchPostsApi() client.Result {
	return GR().GetIGV2PostsResult(r.UserId)
}

func (r *IGV2Repo) ParseRawUID(result client.Result) models.UID {
	data, err := result.GetIGV2UID()
	if err != nil {
		return FetchUIDErr(err)
	}
	return data
}

func (r *IGV2Repo) ParseRawProfile(result client.Result) models.Profile {
	data, err := result.GetIGV2Profile()
	if err != nil {
		return FetchProfileErr(err)
	}
	return data
}

func (r *IGV2Repo) ParseRawPosts(result client.Result) models.Posts {
	data, err := result.GetIGV2Posts()
	if err != nil {
		return FetchPostsErr(err)
	}
	return data
}
