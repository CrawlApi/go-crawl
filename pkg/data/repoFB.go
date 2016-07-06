package data

import (
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/parnurzeal/gorequest"
	"strings"
)

const (
	PAGE_PROFILE_FIELDS_ENABLE  = "name,about,affiliation,app_links,artists_we_like,attire,awards,band_interests,band_members,best_page,birthday,bio,booking_agent,built,business,can_checkin,category,category_list,can_post,checkins,company_overview,contact_address,country_page_likes,cover,culinary_team,access_token,app_id,current_location,description,description_html,directed_by,display_subtext,emails,engagement,fan_count,featured_video,food_styles,features,founded,general_info,general_manager,genre,global_brand_page_name,has_added_app,global_brand_root_id,hometown,hours,id,impressum,influences,is_always_open,is_community_page,is_permanently_closed,is_published,is_unclaimed,is_verified,last_used_time,leadgen_tos_accepted,link,location,members,mission,mpg,name_with_location_descriptor,network,new_like_count,offer_eligible,parent_page,parking,payment_options,personal_info,personal_interests,pharma_safety_info,phone,place_type,press_contact,plot_outline,price_range,produced_by,products,promotion_ineligible_reason,public_transit,record_label,release_date,restaurant_services,restaurant_specialties,schedule,screenplay_by,season,single_line_address,starring,start_info,store_location_descriptor,store_number,studio,talking_about_count,username,verification_status,unread_message_count,unread_notif_count,unseen_message_count,voip_info,website,were_here_count,written_by"
	PAGE_PROFILE_FIELDS_DISABLE = "ad_campaign,context,instant_articles_review_status,owner_business,promotion_eligible,supports_instant_articles"

	PAGE_FEED_FIELDS_ENABLE = "actions,admin_creator,application,call_to_action,child_attachments,caption,comments_mirroring_domain,coordinates,created_time,description,event,expanded_height,expanded_width,feed_targeting,from,full_picture,height,icon,id,is_expired,is_crossposting_eligible,instagram_eligibility,is_hidden,is_instagram_eligible,is_popular,is_published,is_spherical,link,message,message_tags,name,object_id,parent_id,permalink_url,picture,place,privacy,promotion_status,properties,scheduled_publish_time,shares,source,status_type,story,story_tags,subscribed,target,targeting,timeline_visibility,type,updated_time,via,width,with_tags"
	PAGE_FEED_CONNECTIONS   = "comments.limit(1).summary(true),likes.limit(1).summary(true)"

	PAGE_REACTIONS_FIELDS_ENABLE = "reactions.type(NONE).limit(0).summary(total_count).as(reactions_none),reactions.type(LIKE).limit(0).summary(total_count).as(reactions_like),reactions.type(LOVE).limit(0).summary(total_count).as(reactions_love),reactions.type(WOW).limit(0).summary(total_count).as(reactions_wow),reactions.type(HAHA).limit(0).summary(total_count).as(reactions_haha),reactions.type(SAD).limit(0).summary(total_count).as(reactions_sad),reactions.type(ANGRY).limit(0).summary(total_count).as(reactions_angry),reactions.type(THANKFUL).limit(0).summary(total_count).as(reactions_thankful)"
)

const (
	URL_FACEBOOK_PROFILE        = "https://graph.facebook.com/v2.6/%s?fields=%s&access_token=%s"
	URL_FACEBOOK_POSTS          = "https://graph.facebook.com/v2.6/%s/posts?fields=%s,%s&limit=%s&access_token=%s"
	URL_FACEBOOK_POST_REACTIONS = "https://graph.facebook.com/v2.6/%s?fields=%s&access_token=%s"
)

const (
	REGEXP_FACEBOOK_PROFILE_ID         = `fb://(page|profile|group)/(\d+)`
	REGEXP_FACEBOOK_POST_ID            = `posts/(\d+)`
	REGEXP_FACEBOOK_POST_CREATED_AT    = `data-utime="(\d+)"`
	REGEXP_FACEBOOK_POST_COMMENT_COUNT = `"commentcount":(\d+)`
	REGEXP_FACEBOOK_POST_LIKE_COUNT    = `"likecount":(\d+)`
	REGEXP_FACEBOOK_POST_SHARE_COUNT   = `"sharecount":(\d+)`
	REGEXP_FACEBOOK_POST_CONTENT       = `<meta name="description" content="(...+)" /><meta name="robots"`
	REGEXP_FACEBOOK_POST_PIC           = `<img class="scaledImageFitWidth img" src="(...+)" alt="`
	REGEXP_FACEBOOK_POST_PERMALINK     = `"permalink":"(...+)","permalinkcommentid"`
)

type FBRepo struct {
	Agent  *gorequest.SuperAgent
	UserId string
	RawUrl string
	Limit  string
	PostId string
}

func NewFBRepoWithLimit(userId string, limit string) *FBRepo {
	return &FBRepo{
		Agent:  common.GetAgent(),
		UserId: userId,
	}
}

func NewFBRepoWithUid(userId string) *FBRepo {
	return &FBRepo{
		Agent:  common.GetAgent(),
		UserId: userId,
	}
}

func NewFBRepoWithPid(postId string) *FBRepo {
	return &FBRepo{
		Agent:  common.GetAgent(),
		PostId: postId,
	}
}

func NewFBRepoWithUrl(rawUrl string) *FBRepo {
	return &FBRepo{
		Agent:  common.GetAgent(),
		RawUrl: rawUrl,
	}
}

func (r *FBRepo) FetchUIDApi() (string, error) {
	return getApi(r.Agent, r.RawUrl)
}

func (r *FBRepo) FetchProfileApi() (string, error) {
	return getApi(r.Agent, common.UrlString(URL_FACEBOOK_PROFILE, r.UserId, PAGE_PROFILE_FIELDS_ENABLE, common.GetFBToken()))
}

func (r *FBRepo) FetchPostsApi() (string, error) {
	return getApi(r.Agent, common.UrlString(URL_FACEBOOK_POSTS, r.UserId, PAGE_FEED_FIELDS_ENABLE, PAGE_FEED_CONNECTIONS, r.Limit, common.GetFBToken()))
}

func (r *FBRepo) FetchReactionsApi() (string, error) {
	return getApi(r.Agent, common.UrlString(URL_FACEBOOK_POST_REACTIONS, r.PostId, PAGE_REACTIONS_FIELDS_ENABLE, common.GetFBToken()))
}

func (r *FBRepo) FetchPostInfo() (string, error) {
	return getApi(r.Agent, r.RawUrl)
}

func (r *FBRepo) ParseRawUID(body string) models.UID {

	matcher := common.Matcher(REGEXP_FACEBOOK_PROFILE_ID, body)

	var uid models.UID
	uid.Media = "fb"
	if len(matcher) > 2 {
		uid.Type = matcher[1]
		uid.UserId = matcher[2]
		uid.Status = true
	}
	uid.Date = common.Now()
	return uid
}

func (r *FBRepo) ParseRawProfile(body string) models.Profile {
	var rawProfile models.FBRawProfile
	err := common.ParseJson(body, &rawProfile)

	var profile models.Profile
	if err != nil {
		profile.FetchErr(err)
		return profile
	}
	profile.ParseFBProfile(rawProfile)

	return profile
}

func (r *FBRepo) ParseRawPosts(body string) models.Posts {
	var rawPosts models.FBRawPosts
	err := common.ParseJson(body, &rawPosts)
	var posts models.Posts
	if err != nil {
		posts.FetchErr(err)
		return posts
	}
	posts.ParseFBRawPosts(rawPosts)

	return posts
}

func (r *FBRepo) ParsePostInfo(body string) models.Post {

	data, _ := r.parseRawPost(body)

	var post models.Post
	post.ParseFBRawPost(data)
	return post
}

func (r *FBRepo) parseRawPost(body string) (models.FBRawPost, error) {
	var data models.FBRawPost
	data.ID = common.GetMatcherValue(1, REGEXP_FACEBOOK_POST_ID, r.RawUrl)
	data.CreatedAt = common.Str2Int64(common.GetMatcherValue(1, REGEXP_FACEBOOK_POST_CREATED_AT, body))
	//data.UpdatedAt = common.GetMatcherValue(1, REGEXP_FACEBOOK_POST_ID, body)
	data.ShareCount = common.Str2Int(common.GetMatcherValue(1, REGEXP_FACEBOOK_POST_SHARE_COUNT, body))
	data.LikeCount = common.Str2Int(common.GetMatcherValue(1, REGEXP_FACEBOOK_POST_LIKE_COUNT, body))
	data.CommentCount = common.Str2Int(common.GetMatcherValue(1, REGEXP_FACEBOOK_POST_COMMENT_COUNT, body))
	//data.ViewCount = common.GetMatcherValue(1, REGEXP_FACEBOOK_POST_ID, body)
	//data.ContentType = common.GetMatcherValue(1, REGEXP_FACEBOOK_POST_ID, body)
	//data.ContentCaption = common.GetMatcherValue(1, REGEXP_FACEBOOK_POST_ID, body)
	data.ContentBody = common.GetMatcherValue(1, REGEXP_FACEBOOK_POST_CONTENT, body)
	data.ContentFullPicture = common.DecodeString(common.GetMatcherValue(1, REGEXP_FACEBOOK_POST_PIC, body))
	data.PermalinkUrl = r.getPermalink(common.GetMatcherValue(1, REGEXP_FACEBOOK_POST_PERMALINK, body))
	//data.HasComment = common.GetMatcherValue(1, REGEXP_FACEBOOK_POST_ID, body)
	return data, nil
}

func (r *FBRepo) getPermalink(src string) string {
	src = strings.Replace(src, `\`, "", -1)
	src = "https://www.facebook.com" + src
	return src
}

func (r *FBRepo) ParseRawReactions(body string) models.FBReactions {
	var data models.FBRawReactions
	err := common.ParseJson(body, &data)

	var reactions models.FBReactions
	if err != nil {
		reactions.FetchErr(err)
		return reactions
	}
	reactions.ParseFBReactions(data)

	return reactions
}
