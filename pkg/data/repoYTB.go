package data

import (
	"github.com/llitfkitfk/cirkol/pkg/client"
	"github.com/llitfkitfk/cirkol/pkg/models"
)

type YTBRepo struct {
	UserId string
	RawUrl string
}

func NewYTBRepoWithUid(userId string) *YTBRepo {
	return &YTBRepo{
		UserId: userId,
	}
}

func NewYTBRepoWithUrl(rawUrl string) *YTBRepo {
	return &YTBRepo{
		RawUrl: rawUrl,
	}
}

func (r *YTBRepo) FetchUIDApi() client.Result {
	return GR().GetYTBUIDResult(r.RawUrl)
}

func (r *YTBRepo) FetchProfileApi() client.Result {
	return GR().GetYTBProfileResult(r.UserId)
}

func (r *YTBRepo) FetchPostsApi() client.Result {
	return GR().GetYTBPostsResult(r.UserId)
}

func (r *YTBRepo) ParseRawUID(result client.Result) models.UID {
	data, err := result.GetYTBUID()
	if err != nil {
		return FetchUIDErr(err)
	}
	return data
}

func (r *YTBRepo) ParseRawProfile(result client.Result) models.Profile {
	data, err := result.GetYTBProfile()
	if err != nil {
		return FetchProfileErr(err)
	}
	return data
}

func (r *YTBRepo) ParseRawPosts(result client.Result) models.Posts {
	data, err := result.GetYTBPosts()
	if err != nil {
		return FetchPostsErr(err)
	}
	return data
}
