package data

import (
	"github.com/llitfkitfk/cirkol/pkg/client"
	"github.com/llitfkitfk/cirkol/pkg/models"
)

type WXRepo struct {
	UserId string
	RawUrl string
}

func NewWXRepoWithUid(userId string) *WXRepo {
	return &WXRepo{
		UserId: userId,
	}
}

func NewWXRepoWithUrl(rawUrl string) *WXRepo {
	return &WXRepo{
		RawUrl: rawUrl,
	}
}

func (r *WXRepo) FetchUIDApi() client.Result {
	return GR().GetWXUIDResult(r.RawUrl)
}

func (r *WXRepo) FetchProfileApi() client.Result {
	return GR().GetWXProfileResult(r.UserId)
}

func (r *WXRepo) FetchPostsApi() client.Result {
	return GR().GetWXPostsResult(r.UserId)
}

func (r *WXRepo) ParseRawUID(result client.Result) models.UID {
	data, err := result.GetWXUID()
	if err != nil {
		return FetchUIDErr(err)
	}
	return data
}

func (r *WXRepo) ParseRawProfile(result client.Result) models.Profile {
	data, err := result.GetWXProfile()
	if err != nil {
		return FetchProfileErr(err)
	}
	return data
}

func (r *WXRepo) ParseRawPosts(result client.Result) models.Posts {
	data, err := result.GetWXPosts()
	if err != nil {
		return FetchPostsErr(err)
	}
	return data

}
