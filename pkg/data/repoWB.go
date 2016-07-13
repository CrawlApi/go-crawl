package data

import (
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/llitfkitfk/cirkol/pkg/client"
)

type WBRepo struct {
	UserId string
	RawUrl string
}

func NewWBRepoWithUid(userId string) *WBRepo {
	return &WBRepo{
		UserId: userId,
	}
}

func NewWBRepoWithUrl(rawUrl string) *WBRepo {
	return &WBRepo{
		RawUrl: rawUrl,
	}
}

func (r *WBRepo) FetchUIDApi() client.Result {
	return GR().GetWBUIDResult(r.RawUrl)
}

func (r *WBRepo) FetchProfileApi() client.Result {
	return GR().GetWBProfileResult(r.UserId)
}

func (r *WBRepo) FetchPostsApi() client.Result {
	return GR().GetWBPostsResult(r.UserId)
}

func (r *WBRepo) FetchPostInfo() client.Result {
	return GR().GetWBPostResult(r.UserId)
}

func (r *WBRepo) ParseRawUID(result client.Result) models.UID {
	data, err := result.GetWBUID()
	if err != nil {
		return FetchUIDErr(err)
	}
	return data
}

func (r *WBRepo) ParseRawProfile(result client.Result) models.Profile {
	data, err := result.GetWBProfile()
	if err != nil {
		return FetchProfileErr(err)
	}
	return data
}

func (r *WBRepo) ParseRawPosts(result client.Result) models.Posts {
	data, err := result.GetWBPosts()
	if err != nil {
		return FetchPostsErr(err)
	}
	return data
}

func (r *WBRepo) ParsePostInfo(result client.Result) models.Post {
	data, err := result.GetWBPost()
	if err != nil {
		return FetchPostErr(err)
	}
	return data

}
