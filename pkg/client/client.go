package client

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/parser"
	"github.com/parnurzeal/gorequest"
	"time"
)

const (
	PAGE_PROFILE_FIELDS_ENABLE = "name,about,affiliation,app_links,artists_we_like,attire,awards,band_interests,band_members,best_page,birthday,bio,booking_agent,built,business,can_checkin,category,category_list,can_post,checkins,company_overview,contact_address,country_page_likes,cover,culinary_team,access_token,app_id,current_location,description,description_html,directed_by,display_subtext,emails,engagement,fan_count,featured_video,food_styles,features,founded,general_info,general_manager,genre,global_brand_page_name,has_added_app,global_brand_root_id,hometown,hours,id,impressum,influences,is_always_open,is_community_page,is_permanently_closed,is_published,is_unclaimed,is_verified,last_used_time,leadgen_tos_accepted,link,location,members,mission,mpg,name_with_location_descriptor,network,new_like_count,offer_eligible,parent_page,parking,payment_options,personal_info,personal_interests,pharma_safety_info,phone,place_type,press_contact,plot_outline,price_range,produced_by,products,promotion_ineligible_reason,public_transit,record_label,release_date,restaurant_services,restaurant_specialties,schedule,screenplay_by,season,single_line_address,starring,start_info,store_location_descriptor,store_number,studio,talking_about_count,username,verification_status,unread_message_count,unread_notif_count,unseen_message_count,voip_info,website,were_here_count,written_by"
	PAGE_PROFILE_FIELDS_DISABLE = "ad_campaign,context,instant_articles_review_status,owner_business,promotion_eligible,supports_instant_articles"

	PAGE_FEED_FIELDS_ENABLE = "actions,admin_creator,application,call_to_action,child_attachments,caption,comments_mirroring_domain,coordinates,created_time,description,event,expanded_height,expanded_width,feed_targeting,from,full_picture,height,icon,id,is_expired,is_crossposting_eligible,instagram_eligibility,is_hidden,is_instagram_eligible,is_popular,is_published,is_spherical,link,message,message_tags,name,object_id,parent_id,permalink_url,picture,place,privacy,promotion_status,properties,scheduled_publish_time,shares,source,status_type,story,story_tags,subscribed,target,targeting,timeline_visibility,type,updated_time,via,width,with_tags"
	PAGE_FEED_CONNECTIONS = "comments.limit(1).summary(true),likes.limit(1).summary(true)"

	PAGE_REACTIONS_FIELDS_ENABLE = "reactions.type(NONE).limit(0).summary(total_count).as(reactions_none),reactions.type(LIKE).limit(0).summary(total_count).as(reactions_like),reactions.type(LOVE).limit(0).summary(total_count).as(reactions_love),reactions.type(WOW).limit(0).summary(total_count).as(reactions_wow),reactions.type(HAHA).limit(0).summary(total_count).as(reactions_haha),reactions.type(SAD).limit(0).summary(total_count).as(reactions_sad),reactions.type(ANGRY).limit(0).summary(total_count).as(reactions_angry),reactions.type(THANKFUL).limit(0).summary(total_count).as(reactions_thankful)"
	PAGE_POST_INFO_FIELDS = "comments.limit(1).summary(true),likes.limit(1).summary(true),shares,application,actions,caption,admin_creator,call_to_action,child_attachments,comments_mirroring_domain,coordinates,created_time,description,event,expanded_height,expanded_width,feed_targeting,from,full_picture,height,icon,id,is_expired,is_hidden,is_instagram_eligible,is_popular,is_published,is_spherical,link,message,message_tags,name,object_id,parent_id,picture,place,privacy,promotion_status,properties,scheduled_publish_time,source,status_type,story,story_tags,subscribed,target,timeline_visibility,targeting,type,updated_time,via,width,attachments,insights,permalink_url,dynamic_posts,to,with_tags,sponsor_tags"
)

const (
	URL_FACEBOOK_PROFILE = "https://graph.facebook.com/v2.6/%s?fields=%s&access_token=%s"
	URL_FACEBOOK_POSTS = "https://graph.facebook.com/v2.6/%s/posts?fields=%s,%s&limit=%s&access_token=%s"
	URL_FACEBOOK_POST_REACTIONS = "https://graph.facebook.com/v2.6/%s?fields=%s&access_token=%s"
	URL_FACEBOOK_POST_INFO = `https://graph.facebook.com/v2.6/%s%s%s?fields=%s&access_token=%s`
)

const (
	URL_WEIBO_PROFILE = "http://mapi.weibo.com/2/profile?gsid=_&c=&s=&user_domain=%s"
	URL_WEIBO_POSTS = "http://m.weibo.cn/%s"
	URL_WEIBO_API_POSTS = "http://m.weibo.cn/page/tpl?containerid=%s_-_WEIBO_SECOND_PROFILE_WEIBO&itemid=&title=全部微博"
	WEIBO_POST_LINK_PREF = "http://m.weibo.cn/%s"
)

const (
	URL_INSTAGRAM_PROFILE = "https://www.instagram.com/%s/"
	URL_INSTAGRAM_POSTS = "https://www.instagram.com/%s/media/"

	URL_INSTAGRAM_PROFILE_V2 = "https://i.instagram.com/api/v1/users/%s/info/"
	URL_INSTAGRAM_POSTS_V2 = "https://i.instagram.com/api/v1/users/%s/info/"

	REGEX_INSTAGRAM_PROFILE = `ProfilePage": \[([\s\S]+), "nodes": ([\s\S]+)]([\s\S]+)]},`
	REGEX_INSTAGRAM_POST_INFO = `_sharedData =(...+);</script>`

	URL_INSTAGRAM_API_POSTS = "https://www.instagram.com/%s/"
	REGEX_INSTAGRAM_POSTS = `ProfilePage": \[([\s\S]+), "nodes": ([\s\S]+)]([\s\S]+)]},`
	REGEX_INSTAGRAM_PROFILE_ID = `"owner": {"id": "(\d+)`
)

const (
	URL_WECHAT_PROFILE = "http://weixin.sogou.com/weixin?type=1&query=%s&ie=utf8&_sug_=n&_sug_type_="
	URL_WECHAT_POSTS = "http://weixin.sogou.com/weixin?type=1&query=%s&ie=utf8&_sug_=n&_sug_type_="

	REGEXP_WEIXIN_PROFILE_ID = `微信号: (\S+)</p>`
	REGEXP_WEIXIN_LOGO = `src="((http://img01.sogoucdn.com/app/a)\S+)"`
	REGEXP_WEIXIN_NAME = `<h3>(\S+)</h3>`
	REGEXP_WEIXIN_FEATURE = `功能介绍(...+)class="sp-txt">(...+)</span>`
	REGEXP_WEIXIN_URL = `href="((http://mp.weixin.qq.com/profile)\S+)"`
	REGEXP_WEIXIN_POSTS = `var msgList = '(\S+)';`
)

const (
	URL_YOUTUBE_PROFILE = "https://www.youtube.com/user/%s/about"
	URL_YOUTUBE_POSTS   = "https://www.youtube.com/user/%s/videos"
)


type Client struct {
	agent *gorequest.SuperAgent
}

func New() *Client {
	client := Client{
		agent: gorequest.New(),
	}
	return &client
}

func (c *Client) sendRequest(url string) Result {
	_, body, errs := c.agent.Timeout(10 * time.Second).Set("accept-language", "en-US").Get(url).End()
	return Result{Body: body, err: common.Errs2Error(errs)}
}

// facebook
func (c *Client) GetFBUIDResult(url string) Result {
	return c.sendRequest(url)
}

func (c *Client) GetFBProfileResult(uid string) Result {
	common.Info("Fetch Facebook Profile: ", uid)
	return c.sendRequest(common.UrlString(URL_FACEBOOK_PROFILE, uid, PAGE_PROFILE_FIELDS_ENABLE, common.GetFBToken()))
}

func (c *Client) GetFBPostsResult(uid, limit string) Result {
	return c.sendRequest(common.UrlString(URL_FACEBOOK_POSTS, uid, PAGE_FEED_FIELDS_ENABLE, PAGE_FEED_CONNECTIONS, limit, common.GetFBToken()))
}

func (c *Client) GetFBPostResult(url string) Result {

	uidStr := parser.ParseFBUIDFromUrl(url)
	postSuffId := parser.ParseFBPostSuffId(url)

	uidResult := c.GetFBUIDResult(common.UrlString(`https://www.facebook.com/%s`, uidStr))
	uid, err := uidResult.GetFBUID()
	if err != nil {
		return Result{err: common.ParseUIDError()}
	}
	return c.sendRequest(common.UrlString(URL_FACEBOOK_POST_INFO, uid, "_", postSuffId, PAGE_POST_INFO_FIELDS, common.GetFBToken()))
}

func (c *Client) GetFBReactionsResult(postId string) Result {
	return c.sendRequest(common.UrlString(URL_FACEBOOK_POST_REACTIONS, postId, PAGE_REACTIONS_FIELDS_ENABLE, common.GetFBToken()))
}

// weibo
func (c *Client) GetWBUIDResult(url string) Result {
	return c.sendRequest(url)
}

func (c *Client) GetWBProfileResult(userId string) Result {
	return c.sendRequest(common.UrlString(URL_WEIBO_PROFILE, userId))
}

func (c *Client) GetWBPostsResult(userId string) Result {
	postsResult := c.sendRequest(common.UrlString(URL_WEIBO_POSTS, userId))
	postId := common.DecodeString(parser.ParseWBPostsUrl(postsResult.Body))
	return c.sendRequest(common.UrlString(URL_WEIBO_API_POSTS, postId))
}

func (c *Client) GetWBPostResult(url string) Result {
	postSuffUrl := parser.ParseWBPostUrl(url)
	return c.sendRequest(common.UrlString(WEIBO_POST_LINK_PREF, postSuffUrl))
}

// instagram
func (c *Client) GetIGUIDResult(url string) Result {
	return c.sendRequest(url)
}

func (c *Client) GetIGProfileResult(userId string) Result {
	return c.sendRequest(common.UrlString(URL_INSTAGRAM_PROFILE, userId))
}

func (c *Client) GetIGPostsResult(userId string) Result {
	return c.sendRequest(common.UrlString(URL_INSTAGRAM_POSTS, userId))
}

func (c *Client) GetIGPostResult(url string) Result {
	return c.sendRequest(url)

}

// instagram v2
func (c *Client) GetIGV2UIDResult(url string) Result {
	return c.sendRequest(url)
}

func (c *Client) GetIGV2ProfileResult(userId string) Result {
	return c.sendRequest(common.UrlString(URL_INSTAGRAM_PROFILE_V2, userId))
}

func (c *Client) GetIGV2PostsResult(userId string) Result {
	//body, err := getApi(r.Agent, common.UrlString(URL_INSTAGRAM_POSTS_V2, r.UserId))
	//if err != nil {
	//	return body, err
	//}
	//urlStr, err := r.getPostsUrl(body)
	//if err != nil {
	//	return urlStr, err
	//}
	//postsBody, err := getApi(r.Agent, urlStr)
	//if err != nil {
	//	return postsBody, err
	//}
	//return postsBody, nil
	return c.sendRequest(userId)

}

//func (r *IGV2Repo) getPostsUrl(body string) (string, error) {
//	var data models.IGV2RawProfile
//	err := common.ParseJson(body, &data)
//	postsUrl := common.UrlString(URL_INSTAGRAM_API_POSTS, data.User.Username)
//	if err != nil {
//		return postsUrl, err
//	}
//	return postsUrl, nil
//}

// wechat
func (c *Client) GetWXUIDResult(url string) Result {
	return c.sendRequest(url)
}

func (c *Client) GetWXProfileResult(userId string) Result {
	return c.sendRequest(common.UrlString(URL_WECHAT_PROFILE, userId))
}

func (c *Client) GetWXPostsResult(userId string) Result {
	//body, err := getApi(r.Agent, common.UrlString(URL_WECHAT_POSTS, r.UserId))
	//if err != nil {
	//	return body, err
	//}
	//
	//urlStr, err := r.getPostsUrl(body)
	//if err != nil {
	//	return urlStr, err
	//}
	//
	//postsBody, err := getApi(r.Agent, urlStr)
	//if err != nil {
	//	return postsBody, err
	//}
	//
	//return postsBody, nil
	return c.sendRequest(userId)
}

//func (r *WXRepo) getPostsUrl(body string) (string, error) {
//	matcher := common.Matcher(REGEXP_WEIXIN_URL, body)
//	if len(matcher) > 1 {
//		return matcher[1], nil
//	}
//	return "", errors.New(common.ERROR_MSG_REGEX_MISS_MATCHED)
//}

func (c *Client) GetWXPostResult(url string) Result {
	return c.sendRequest(url)

}

// youtube
func (c *Client) GetYTBUIDResult(url string) Result {
	return c.sendRequest(url)
}

func (c *Client) GetYTBProfileResult(userId string) Result {
	return c.sendRequest(common.UrlString(URL_YOUTUBE_PROFILE, userId))
}

func (c *Client) GetYTBPostsResult(userId string) Result {
	//body, err := getApi(r.Agent, common.UrlString(URL_YOUTUBE_POSTS, r.UserId))
	//if err != nil {
	//	return body, err
	//}
	//
	//return body, nil
	return c.sendRequest(userId)
}

func (c *Client) GetYTBPostResult(url string) Result {
	return c.sendRequest(url)
}

//func (c *Client) Get1Result(uid, limit string) Result {
//}