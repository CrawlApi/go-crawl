package client

import (
	"github.com/llitfkitfk/cirkol/pkg/models"
)

func (r *Result) GetFBPosts() (models.Posts, error) {
	//var rawPosts models.FBRawPosts
	//err := common.ParseJson(result.Body, &rawPosts)
	//var posts models.Posts
	//if err != nil {
	//	posts.FetchErr(err)
	//	return posts
	//}
	//posts.ParseFBRawPosts(rawPosts)
	//
	//return posts
	return models.Posts{}, nil
}

func (r *Result) GetWBPosts() (models.Posts, error) {
	//var rawPosts models.WBRawPosts
	//
	//err := common.ParseJson(r.getPostsStr(body), &rawPosts)
	//
	//var posts models.Posts
	//if err != nil {
	//	posts.FetchErr(err)
	//	return posts
	//}
	//posts.ParseWBRawPosts(rawPosts)
	//
	//return posts
	return models.Posts{}, nil
}

func (r *Result) GetIGPosts() (models.Posts, error) {
	//var data models.IGRawPosts
	//err := common.ParseJson(body, &data)
	//
	//var posts models.Posts
	//if err != nil {
	//	posts.FetchErr(err)
	//	return posts
	//}
	//posts.ParseIGRawPosts(data)
	//
	//return posts
	return models.Posts{}, nil
}

func (r *Result) GetIGV2Posts() (models.Posts, error) {
	//var data models.IGV2RawPosts
	//
	//err := common.ParseJson(r.getRawPostsStr(body), &data)
	//
	//var posts models.Posts
	//if err != nil {
	//	posts.FetchErr(err)
	//	return posts
	//}
	//posts.ParseIGV2RawPosts(data)
	//
	//return posts
	return models.Posts{}, nil
}
//
//func (r *IGV2Repo) getRawPostsStr(body string) string {
//	matcher := common.Matcher(REGEX_INSTAGRAM_POSTS, body)
//	if len(matcher) > 2 {
//		return `{ "nodes": ` + matcher[2] + "]}"
//	}
//	return ""
//}

func (r *Result) GetWXPosts() (models.Posts, error) {
	return models.Posts{}, nil
}

func (r *Result) GetYTBPosts() (models.Posts, error) {
	return models.Posts{}, nil
}

func (r *Result) GetFBPost() (models.Post, error) {
	//var data models.FBRawPost
	//err := common.ParseJson(result.Body, &data)
	//var post models.Post
	//if err != nil {
	//	post.FetchErr(err)
	//
	//} else {
	//	post.ParseFBRawPost(data)
	//}
	//return post
	return models.Post{}, nil
}


//func (r *WBRepo) getPostsStr(body string) string {
//	matcher := common.Matcher(REGEXP_WEIBO_POSTS, body)
//	if len(matcher) > 2 {
//		return "{" + strings.Replace(matcher[2], "(MISSING)", "", -1)
//	}
//	return ""
//}
//
//func (r *WBRepo) parseRawPost(body string) models.WBRawPost {
//	str := common.GetMatcherValue(1, REGEXP_WEIBO_POST_INFO, body)
//	var result models.WBRawPost
//	common.ParseJson(str, &result)
//	return result
//}
//


func (r *Result) GetWBPost() (models.Post, error) {
	//data := r.parseRawPost(body)
	//var post models.Post
	//post.ParseWBRawPost(data)
	//return post
	return models.Post{}, nil
}

func (r *Result) GetIGPost() (models.Post, error) {
	//var data models.IGRawPost
	//
	//err := common.ParseJson(common.GetMatcherValue(1, REGEX_INSTAGRAM_POST_INFO, body), &data)
	//var post models.Post
	//if err != nil {
	//	post.FetchErr(err)
	//
	//} else {
	//	post.ParseIGRawPost(data)
	//}
	//return post
	return models.Post{}, nil
}

func (r *Result) GetIGV2Post() (models.Post, error) {
	return models.Post{}, nil
}

func (r *Result) GetWXPost() (models.Post, error) {
	return models.Post{}, nil
}

func (r *Result) GetYTBPost() (models.Post, error) {
	return models.Post{}, nil
}

func (r *Result) GetFBReactions() (models.FBReactions, error) {
	//var data models.FBRawReactions
	//err := common.ParseJson(result.Body, &data)
	//
	//var reactions models.FBReactions
	//if err != nil {
	//	reactions.FetchErr(err)
	//	return reactions
	//}
	//reactions.ParseFBReactions(data)
	//
	//return reactions
	return models.FBReactions{}, nil

}
