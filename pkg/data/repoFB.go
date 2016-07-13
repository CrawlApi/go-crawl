package data

import (
	"github.com/llitfkitfk/cirkol/pkg/client"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
)



type FBRepo struct {
	UserId string
	RawUrl string
	Limit  string
	PostId string
}

func NewFBRepoWithLimit(userId string, limit string) *FBRepo {
	return &FBRepo{
		UserId: userId,
		Limit: limit,
	}
}

func NewFBRepoWithUid(userId string) *FBRepo {
	return &FBRepo{
		UserId: userId,
	}
}

func NewFBRepoWithPid(postId string) *FBRepo {
	return &FBRepo{
		PostId: postId,
	}
}

func NewFBRepoWithUrl(rawUrl string) *FBRepo {
	return &FBRepo{
		RawUrl: rawUrl,
	}
}

func (r *FBRepo) FetchUIDApi() client.Result {
	return GR().GetFBUIDResult(r.RawUrl)
}

func (r *FBRepo) FetchProfileApi() client.Result {
	return GR().GetFBProfileResult(r.UserId)
}

func (r *FBRepo) FetchPostsApi() client.Result {
	return GR().GetFBPostsResult(r.UserId, r.Limit)
}

func (r *FBRepo) FetchReactionsApi() client.Result {
	return GR().GetFBReactionsResult(r.PostId)
}

func (r *FBRepo) FetchPostInfo() client.Result {
	return GR().GetFBPostResult(r.RawUrl)
}

func (r *FBRepo) ParseRawUID(result client.Result) models.UID {
	data, err := result.GetFBUID()
	if err != nil {
		return FetchUIDErr(err)
	}
	return data
}

func (r *FBRepo) ParseRawProfile(result client.Result) models.Profile {
	data, err := result.GetFBProfile()
	if err != nil {
		return FetchProfileErr(err)
	}
	return data
}

func (r *FBRepo) ParseRawPosts(result client.Result) models.Posts{
	data, err := result.GetFBPosts()
	if err != nil {
		return FetchPostsErr(err)
	}
	return data
}

func (r *FBRepo) ParsePostInfo(result client.Result) models.Post {
	data, err := result.GetFBPost()
	if err != nil {
		return FetchPostErr(err)
	}
	return data

}

func (r *FBRepo) ParseRawReactions(result client.Result) models.FBReactions{
	data, err := result.GetFBReactions()
	if err != nil {
		return FetchReactionsError(err)
	}
	return data


}

func FetchReactionsError(err error) models.FBReactions {
	return models.FBReactions{
		ErrMessage: err.Error(),
		Date:       common.Now(),
		Status:     false,
	}
}