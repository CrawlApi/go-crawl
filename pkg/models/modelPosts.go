package models

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
)

type Post struct {
	ID                 string `json:"id"`
	CreatedAt          int64  `json:"created_at"`
	UpdatedAt          int64  `json:"updated_at"`
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
	Status             bool   `json:"status"`
	ErrMessage         string `json:"error_message"`
}

func (p *Post) FetchErr(err error) {
	p.ErrMessage = err.Error()
	p.Date = common.Now()
	p.Status = false
}

func (p *Post) ParseFBRawPost(data FBRawPost) {
	p.ID = data.ID
	p.CreatedAt = data.CreatedAt
	p.UpdatedAt = data.UpdatedAt
	p.ShareCount = data.ShareCount
	p.LikeCount = data.LikeCount
	p.CommentCount = data.CommentCount
	p.ContentType = data.ContentType
	p.ContentBody = data.ContentBody
	p.ContentFullPicture = data.ContentFullPicture
	p.PermalinkUrl = data.PermalinkUrl
	p.HasComment = data.HasComment
	p.RawData = common.Interface2String(data)
	p.Date = common.Now()
	p.Status = true
}

func (p *Post) ParseWBRawPost(data WBRawPost) {
	p.ID = data.Mblog.Bid
	p.CreatedAt = common.ParseWBCreatedAt(data.Mblog.CreatedAt)
	//p.UpdatedAt = data.Mblog
	p.ShareCount = data.Mblog.RepostsCount
	p.LikeCount = data.Mblog.LikeCount
	p.CommentCount = data.Mblog.CommentsCount
	p.ViewCount = data.Mblog.ReadsCount
	p.ContentType = data.Mblog.PageInfo.ObjectType
	p.ContentBody = data.Mblog.RawText
	p.ContentFullPicture = data.Mblog.PageInfo.PagePic
	p.PermalinkUrl = data.Mblog.PageInfo.PageURL
	//p.HasComment = data
	p.RawData = common.Interface2String(data)
	p.Date = common.Now()
	p.Status = true
}

type Posts struct {
	Items      []Post `json:"data"`
	Date       int64  `json:"date"`
	Status     bool   `json:"status"`
	RawData    string `json:"raw_data"`
	ErrMessage string `json:"error_message"`
}

func (p *Posts) FetchErr(err error) {
	p.ErrMessage = err.Error()
	p.Date = common.Now()
	p.Status = false
}

func (p *Posts) ParseFBRawPosts(data FBRawPosts) {
	for _, item := range data.Data {
		var data Post
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
		data.RawData = common.Interface2String(item)
		data.Date = common.Now()

		p.Items = append(p.Items, data)
	}

	p.Status = true
}

func (p *Posts) ParseWBRawPosts(data WBRawPosts) {
	for _, node := range data.CardGroup {
		var data Post
		data.ID = node.Mblog.Bid

		data.CreatedAt = common.ParseWBCreatedAt(node.Mblog.CreatedAt)
		//data.UpdatedAt = node.Mblog
		data.ShareCount = node.Mblog.RepostsCount
		data.LikeCount = node.Mblog.AttitudesCount
		data.CommentCount = node.Mblog.CommentsCount
		//data.ViewCount
		//data.ContentType        =
		//data.ContentCaption	=
		data.ContentBody = node.Mblog.Text
		data.ContentFullPicture = node.Mblog.ThumbnailPic
		data.PermalinkUrl = "http://m.weibo.cn/" + common.Int2Str(node.Mblog.User.ID) + "/" + node.Mblog.Bid
		data.HasComment = true
		data.RawData = common.Interface2String(node)
		data.Date = common.Now()
		p.Items = append(p.Items, data)
	}
	p.Status = true
}

func (p *Posts) ParseWXRawPosts(data WXRawPosts) {
	for _, item := range data.List {
		var data Post
		data.ID = common.Int2Str(item.AppMsgExtInfo.Fileid)
		data.CreatedAt = int64(item.CommMsgInfo.Datetime)

		data.ContentFullPicture = item.AppMsgExtInfo.Cover

		data.ContentCaption = item.AppMsgExtInfo.Title
		data.ContentBody = item.AppMsgExtInfo.Digest
		data.ContentFullPicture = item.AppMsgExtInfo.Cover
		data.PermalinkUrl = "http://mp.weixin.qq.com" + item.AppMsgExtInfo.ContentURL
		data.RawData = common.Interface2String(item)
		data.Date = common.Now()

		p.Items = append(p.Items, data)

		if item.AppMsgExtInfo.IsMulti == 1 {
			for _, item2 := range item.AppMsgExtInfo.MultiAppMsgItemList {
				var data2 Post
				data2.ID = common.Int2Str(item2.Fileid)
				data2.CreatedAt = int64(item.CommMsgInfo.Datetime)

				data2.ContentFullPicture = item2.Cover

				data2.ContentCaption = item2.Title
				data2.ContentBody = item2.Digest
				data2.ContentFullPicture = item2.Cover
				data2.PermalinkUrl = "http://mp.weixin.qq.com" + item2.ContentURL
				data2.RawData = common.Interface2String(item2)
				data2.Date = common.Now()

				p.Items = append(p.Items, data2)
			}
		}
	}
	p.Status = true
}

func (p *Posts) ParseIGRawPosts(rawPost IGRawPosts) {

	for _, item := range rawPost.Items {
		var data Post
		data.ID = item.ID
		data.CreatedAt = common.Str2Int64(item.CreatedTime)

		data.LikeCount = item.Likes.Count
		data.CommentCount = item.Comments.Count
		data.ViewCount = item.VideoViews
		data.ContentType = item.Type
		data.ContentCaption = item.Caption.Text

		data.ContentFullPicture = item.Images.StandardResolution.URL
		data.PermalinkUrl = item.Link
		data.HasComment = item.CanViewComments
		data.RawData = common.Interface2String(item)
		data.Date = common.Now()

		p.Items = append(p.Items, data)
	}

	p.Status = true
}

func (p *Posts) ParseIGV2RawPosts(rawPost IGV2RawPosts) {
	for _, node := range rawPost.Nodes {
		var data Post
		data.ID = node.ID
		data.CreatedAt = int64(node.Date)

		data.LikeCount = node.Likes.Count
		data.CommentCount = node.Comments.Count
		data.ViewCount = node.VideoViews
		if node.IsVideo {
			data.ContentType = "video"
		} else {
			data.ContentType = "image"
		}

		data.ContentCaption = node.Caption
		data.ContentFullPicture = node.DisplaySrc
		data.PermalinkUrl = "https://www.instagram.com/p/" + node.Code + "/"
		data.HasComment = true
		data.RawData = common.Interface2String(node)
		data.Date = common.Now()

		p.Items = append(p.Items, data)
	}

	p.Status = true
}
