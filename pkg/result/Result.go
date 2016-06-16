package result

import (
	"encoding/json"
	"github.com/llitfkitfk/cirkol/pkg/util"
	"time"
)

type UID struct {
	Type       string `json:"type"`
	UserId     string `json:"user_id"`
	Date       int64  `json:"date"`
	Url        string `json:"url"`
	Status     bool   `json:"status"`
	Media      string `json:"media"`
	Message    string `json:"message"`
	ErrCode    int    `json:"error_code"`
	ErrMessage string `json:"error_message"`
}

type Profile struct {
	UserId     string `json:"user_id"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	PostNum    int    `json:"post_num"`
	FollowNum  int    `json:"follow_num"`
	Fans       int    `json:"fans"`
	Birthday   string `json:"birthday"`
	Website    string `json:"website"`
	About      string `json:"about"`
	RawData    string `json:"raw_data"`
	Status     bool   `json:"status"`
	Date       int64  `json:"date"`
	ErrCode    int    `json:"error_code"`
	ErrMessage string `json:"error_message"`
}

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

func (p *Profile) MergeIGNameProfile(data IGNameRawProfile) {
	p.UserId = data.User.Username
	p.Name = data.User.FullName
	p.Avatar = data.User.ProfilePicURL
	p.PostNum = data.User.Media.Count
	p.FollowNum = data.User.Follows.Count
	p.Fans = data.User.FollowedBy.Count

	p.Website = "https://www.instagram.com/" + data.User.Username + "/"
	p.Status = true
}

func (p *Profile) MergeIGIDProfile(data IGIDRawProfile) {
	p.UserId = util.Int2Str(data.User.Pk)
	p.Name = data.User.Username
	p.Avatar = data.User.ProfilePicURL
	p.PostNum = data.User.MediaCount
	p.FollowNum = data.User.FollowingCount
	p.Fans = data.User.FollowerCount
	p.About = data.User.Biography
	p.Website = "https://www.instagram.com/" + data.User.Username + "/"
	p.Status = true
}

func (p *Posts) MergeIGIdPosts(rawPost IGIDRawPosts) {
	for _, node := range rawPost.Nodes {
		var data post
		data.ID = node.ID
		data.CreatedAt = util.Int2Str(node.Date)

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
		data.RawData = util.JsonToString(json.Marshal(node))
		data.Date = time.Now().Unix()

		p.Items = append(p.Items, data)
	}

	p.Status = true
}

func (p *Posts) MergeIGNamePosts(rawPost IGNameRawPosts) {

	for _, item := range rawPost.Items {
		var data post
		data.ID = item.ID
		data.CreatedAt = item.CreatedTime

		data.LikeCount = item.Likes.Count
		data.CommentCount = item.Comments.Count
		data.ViewCount = item.VideoViews
		data.ContentType = item.Type
		data.ContentCaption = item.Caption.Text

		data.ContentFullPicture = item.Images.StandardResolution.URL
		data.PermalinkUrl = item.Link
		data.HasComment = item.CanViewComments
		data.RawData = util.JsonToString(json.Marshal(item))
		data.Date = time.Now().Unix()

		p.Items = append(p.Items, data)
	}

	p.Status = true
}

func (p *Profile) MergeFBRawProfile(data FBRawProfile) {
	p.UserId = data.ID
	p.Name = data.Name
	p.Avatar = data.Cover.Source

	p.Fans = data.Engagement.Count
	p.About = data.About
	p.Website = data.Link
	p.Status = true
}

func (p *Posts) MergeFBRawPosts(rawPosts FBRawPosts) {
	for _, item := range rawPosts.Data {
		var data post
		data.ID = item.ID
		data.CreatedAt = util.DateFormat(item.CreatedTime)
		data.UpdatedAt = util.DateFormat(item.UpdatedTime)
		data.ShareCount = item.Shares.Count
		data.LikeCount = item.Likes.Summary.TotalCount
		data.CommentCount = item.Comments.Summary.TotalCount

		data.ContentType = item.Type

		data.ContentBody = item.Message
		data.ContentFullPicture = item.FullPicture
		data.PermalinkUrl = item.PermalinkURL
		data.HasComment = item.Comments.Summary.CanComment
		data.RawData = util.JsonToString(json.Marshal(item))
		data.Date = time.Now().Unix()

		p.Items = append(p.Items, data)
	}

	p.Status = true
}

func (p *Profile) MergeWBRawProfile(data WBRawProfile) {
	p.UserId = util.Int2Str(data.UserInfo.ID)
	p.Name = data.UserInfo.Name
	p.Avatar = data.UserInfo.ProfileImageURL
	p.PostNum = data.UserInfo.StatusesCount
	p.FollowNum = data.UserInfo.FriendsCount
	p.Fans = data.UserInfo.FollowersCount
	p.About = data.UserInfo.Description

	p.Status = true
}

func (p *Posts) MergeWXRawPosts(rawPosts WXRawPosts) {
	for _, item := range rawPosts.List {
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
