package client

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/llitfkitfk/cirkol/pkg/parser"
)

func (r *Result) GetFBPosts() (models.Posts, error) {
	var posts models.Posts
	if r.err != nil {
		return posts, r.err
	}

	var rawPosts models.FBRawPosts
	err := common.ParseJson(r.Body, &rawPosts)
	if err != nil {
		return posts, err
	}

	posts.ParseFBRawPosts(rawPosts)
	return posts, nil
}

func (r *Result) GetWBPosts() (models.Posts, error) {
	var posts models.Posts
	if r.err != nil {
		return posts, r.err
	}

	var rawPosts models.WBRawPosts
	err := common.ParseJson(parser.ParseWBPostsStr(r.Body), &rawPosts)
	if err != nil {
		return posts, err
	}

	posts.ParseWBRawPosts(rawPosts)
	return posts, nil
}

func (r *Result) GetIGPosts() (models.Posts, error) {
	var posts models.Posts
	if r.err != nil {
		return posts, r.err
	}

	var data models.IGRawPosts
	err := common.ParseJson(r.Body, &data)
	if err != nil {
		return posts, err
	}
	posts.ParseIGRawPosts(data)

	return posts, nil
}

func (r *Result) GetIGV2Posts() (models.Posts, error) {
	var posts models.Posts
	if r.err != nil {
		return posts, r.err
	}

	var data models.IGV2RawPosts
	err := common.ParseJson(parser.ParseIGV2PostsStr(r.Body), &data)
	if err != nil {
		return posts, err
	}

	posts.ParseIGV2RawPosts(data)

	return posts, nil
}

func (r *Result) GetWXPosts() (models.Posts, error) {
	var posts models.Posts
	if r.err != nil {
		return posts, r.err
	}

	var data models.WXRawPosts
	err := common.ParseJson(parser.ParseWXPostsStr(r.Body), &data)
	if err != nil {
		return posts, err
	}

	posts.ParseWXRawPosts(data)

	return posts, nil
}

func (r *Result) GetYTBPosts() (models.Posts, error) {
	//var rawPosts models.WBRawPosts
	//
	//err := common.ParseJson(body, &rawPosts)
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

func (r *Result) GetFBPost() (models.Post, error) {
	var post models.Post
	if r.err != nil {
		return post, r.err
	}

	var data models.FBRawPost
	err := common.ParseJson(r.Body, &data)
	if err != nil {
		return post, err
	}

	post.ParseFBRawPost(data)
	return post, nil
}

func (r *Result) parseRawPost(body string) (models.WBRawPost, error) {
	str := parser.ParseWBPostStr(body)
	var result models.WBRawPost
	err := common.ParseJson(str, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *Result) GetWBPost() (models.Post, error) {
	var post models.Post
	if r.err != nil {
		return post, r.err
	}

	data, err := r.parseRawPost(r.Body)
	if err != nil {
		return post, err
	}

	post.ParseWBRawPost(data)
	return post, nil
}

func (r *Result) GetIGPost() (models.Post, error) {
	var post models.Post
	if r.err != nil {
		return post, r.err
	}

	var data models.IGRawPost
	err := common.ParseJson(parser.ParseIGPostStr(r.Body), &data)
	if err != nil {
		return post, err
	}

	post.ParseIGRawPost(data)

	return post, nil
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
	var reactions models.FBReactions
	if r.err != nil {
		return reactions, r.err
	}

	var data models.FBRawReactions
	err := common.ParseJson(r.Body, &data)
	if err != nil {
		return reactions, err
	}

	reactions.ParseFBReactions(data)
	return reactions, nil
}
