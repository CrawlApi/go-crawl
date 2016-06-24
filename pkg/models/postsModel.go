package models

import (
	"encoding/json"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/util"
	"time"
)

type post struct {
	ID                 string `json:"id"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	ShareCount         int    `json:"share_count"`
	LikeCount          int    `json:"like_count"`
	CommentCount       int    `json:"comment_count"`
	ViewCount          int    `json:"view_count"`
	ContentType        string `json:"content_type"`
	ContentCaption     string `json:"content_caption"`
	ContentBody        string `json:"content_body"`
	ContentFullPicture string `json:"content_full_picture"`
	PermalinkUrl       string `json:"permalink_url"`
	HasComment         bool   `json:"has_comment"`
	RawData            string `json:"raw_data"`
	Date               int64  `json:"date"`
}

type Posts struct {
	Items      []post `json:"data"`
	Date       int64  `json:"date"`
	Status     bool   `json:"status"`
	RawData    string `json:"raw_data"`
	ErrCode    int    `json:"error_code"`
	ErrMessage string `json:"error_message"`
}

func (p *Posts) FetchErr(err error) {
	p.ErrCode = common.ERROR_CODE_API_FETCH
	p.ErrMessage = err.Error()
	p.Date = time.Now().Unix()
	p.Status = false
}

func (p *Posts) ParseFBRawPosts(data FBRawPosts) {
	for _, item := range data.Data {
		var data post
		data.ID = item.ID
		data.CreatedAt = common.DateFormat(item.CreatedTime)
		data.UpdatedAt = common.DateFormat(item.UpdatedTime)
		data.ShareCount = item.Shares.Count
		data.LikeCount = item.Likes.Summary.TotalCount
		data.CommentCount = item.Comments.Summary.TotalCount

		data.ContentType = item.Type

		data.ContentBody = item.Message
		data.ContentFullPicture = item.FullPicture
		data.PermalinkUrl = item.PermalinkURL
		data.HasComment = item.Comments.Summary.CanComment
		data.RawData = common.JsonToString(json.Marshal(item))
		data.Date = time.Now().Unix()

		p.Items = append(p.Items, data)
	}

	p.Status = true
}

func (p *Posts) ParseWXRawPosts(data WXRawPosts) {
	for _, item := range data.List {
		var data post
		data.ID = util.Int2Str(item.AppMsgExtInfo.Fileid)
		data.CreatedAt = util.Int2Str(item.CommMsgInfo.Datetime)

		data.ContentFullPicture = item.AppMsgExtInfo.Cover

		data.ContentCaption = item.AppMsgExtInfo.Title
		data.ContentBody = item.AppMsgExtInfo.Digest
		data.ContentFullPicture = item.AppMsgExtInfo.Cover
		data.PermalinkUrl = "http://mp.weixin.qq.com" + item.AppMsgExtInfo.ContentURL
		data.RawData = util.JsonToString(json.Marshal(item))
		data.Date = time.Now().Unix()

		p.Items = append(p.Items, data)

		if item.AppMsgExtInfo.IsMulti == 1 {
			for _, item2 := range item.AppMsgExtInfo.MultiAppMsgItemList {
				var data2 post
				data2.ID = util.Int2Str(item2.Fileid)
				data2.CreatedAt = util.Int2Str(item.CommMsgInfo.Datetime)

				data2.ContentFullPicture = item2.Cover

				data2.ContentCaption = item2.Title
				data2.ContentBody = item2.Digest
				data2.ContentFullPicture = item2.Cover
				data2.PermalinkUrl = "http://mp.weixin.qq.com" + item2.ContentURL
				data2.RawData = util.JsonToString(json.Marshal(item2))
				data2.Date = time.Now().Unix()

				p.Items = append(p.Items, data2)
			}
		}
	}
	p.Status = true
}
