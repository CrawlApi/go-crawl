package models

import "github.com/llitfkitfk/cirkol/pkg/common"

type FBReactions struct {
	PostID            string `json:"post_id"`
	ReactionsAngry    int    `json:"reactions_angry"`
	ReactionsHaha     int    `json:"reactions_haha"`
	ReactionsLike     int    `json:"reactions_like"`
	ReactionsLove     int    `json:"reactions_love"`
	ReactionsNone     int    `json:"reactions_none"`
	ReactionsSad      int    `json:"reactions_sad"`
	ReactionsThankful int    `json:"reactions_thankful"`
	ReactionsWow      int    `json:"reactions_wow"`
	RawData           string `json:"raw_data"`
	Status            bool   `json:"status"`
	Date              int64  `json:"date"`
	ErrMessage        string `json:"error_message"`
}

func (r *FBReactions) FetchErr(err error) {
	r.ErrMessage = err.Error()
	r.Date = common.Now()
	r.Status = false
}

func (r *FBReactions) ParseFBReactions(data FBRawReactions) {
	r.PostID = data.ID
	r.ReactionsAngry = data.ReactionsAngry.Summary.TotalCount
	r.ReactionsHaha = data.ReactionsHaha.Summary.TotalCount
	r.ReactionsLike = data.ReactionsLike.Summary.TotalCount
	r.ReactionsLove = data.ReactionsLove.Summary.TotalCount
	r.ReactionsNone = data.ReactionsNone.Summary.TotalCount
	r.ReactionsSad = data.ReactionsSad.Summary.TotalCount
	r.ReactionsThankful = data.ReactionsThankful.Summary.TotalCount
	r.ReactionsWow = data.ReactionsWow.Summary.TotalCount

	r.RawData = common.Interface2String(data)
	r.Date = common.Now()
	r.Status = true
}
