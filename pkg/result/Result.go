package result

import (
	"strconv"
	"time"
)

type UID struct {
	Type    string `json:"type"`
	UserId  string `json:"user_id"`
	Date    int64  `json:"date"`
	Url     string `json:"url"`
	Status  bool   `json:"status"`
	Media   string `json:"media"`
	Message string `json:"message"`
}

type Profile struct {
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	PostNum   int    `json:"post_num"`
	FollowNum int    `json:"follow_num"`
	Fans      int    `json:"fans"`
	Birthday  string `json:"birthday"`
	Website   string `json:"website"`
	About     string `json:"about"`
	RawData   string `json:"raw_data"`
	Status    bool   `json:"status"`
	Date      int64  `json:"date"`
}

type post struct {
	CreatedAt          string `json:"created_at"`
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
}

type Posts struct {
	Items   []post `json:"data"`
	Date    int64  `json:"date"`
	Status  bool   `json:"status"`
	RawData string `json:"raw_data"`
}

func (p *Posts) MergeIGIdPosts(rawPost IGIDRawPosts) {
	var data post
	for _, node := range rawPost.Nodes {
		data.CreatedAt = strconv.Itoa(node.Date)

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

		p.Items = append(p.Items, data)
	}

	p.Status = true
}

func (p *Posts) MergeIGNamePosts(rawPost IGNameRawPosts) {

	var data post
	for _, item := range rawPost.Items {
		data.CreatedAt = item.CreatedTime

		data.LikeCount = item.Likes.Count
		data.CommentCount = item.Comments.Count
		data.ViewCount = item.VideoViews
		data.ContentType = item.Type

		data.ContentCaption = item.Caption.Text
		data.ContentFullPicture = item.Images.StandardResolution.URL
		data.PermalinkUrl = item.Link
		data.HasComment = item.CanViewComments

		p.Items = append(p.Items, data)
	}

	p.Status = true
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
	p.Date = time.Now().Unix()

}

func (p *Profile) MergeIGIDProfile(data IGIDRawProfile) {
	p.UserId = strconv.Itoa(data.User.Pk)
	p.Name = data.User.Username
	p.Avatar = data.User.ProfilePicURL
	p.PostNum = data.User.MediaCount
	p.FollowNum = data.User.FollowingCount
	p.Fans = data.User.FollowerCount
	p.About = data.User.Biography
	p.Website = "https://www.instagram.com/" + data.User.Username + "/"
	p.Status = true
	p.Date = time.Now().Unix()

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

func (p *Profile) MergeWBRawProfile(data WBRawProfile) {
	p.UserId = strconv.Itoa(data.UserInfo.ID)
	p.Name = data.UserInfo.Name
	p.Avatar = data.UserInfo.ProfileImageURL
	p.PostNum = data.UserInfo.StatusesCount
	p.FollowNum = data.UserInfo.FriendsCount
	p.Fans = data.UserInfo.FollowersCount
	p.About = data.UserInfo.Description

	p.Status = true
}



