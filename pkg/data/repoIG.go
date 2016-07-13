package data

import (
	"github.com/llitfkitfk/cirkol/pkg/client"
	"github.com/llitfkitfk/cirkol/pkg/models"
)

type IGRepo struct {
	UserId string
	RawUrl string
}

func NewIGRepoWithUid(userId string) *IGRepo {
	return &IGRepo{
		UserId: userId,
	}
}

func NewIGRepoWithUrl(rawUrl string) *IGRepo {
	return &IGRepo{
		RawUrl: rawUrl,
	}
}

func (r *IGRepo) FetchUIDApi() client.Result {
	return GR().GetIGUIDResult(r.RawUrl)
}

func (r *IGRepo) FetchProfileApi() client.Result {
	return GR().GetIGProfileResult(r.UserId)
}

func (r *IGRepo) FetchPostsApi() client.Result {
	return GR().GetIGPostsResult(r.UserId)
}

func (r *IGRepo) FetchPostInfo() client.Result {
	return GR().GetIGPostResult(r.RawUrl)
}

func (r *IGRepo) ParseRawUID(result client.Result) models.UID {
	data, err := result.GetIGUID()
	if err != nil {
		return FetchUIDErr(err)
	}
	return data

}

func (r *IGRepo) ParseRawProfile(result client.Result) models.Profile {
	data, err := result.GetIGProfile()
	if err != nil {
		return FetchProfileErr(err)
	}
	return data
}

func (r *IGRepo) ParseRawPosts(result client.Result) models.Posts {
	data, err := result.GetIGPosts()
	if err != nil {
		return FetchPostsErr(err)
	}
	return data
}

func (r *IGRepo) ParsePostInfo(result client.Result) models.Post {
	data, err := result.GetIGPost()
	if err != nil {
		return FetchPostErr(err)
	}
	return data
}
