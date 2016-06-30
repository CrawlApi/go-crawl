package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"net/http"
)

const (
	PAGE_PROFILE_FIELDS_ENABLE = "name,about,affiliation,app_links,artists_we_like,attire,awards,band_interests,band_members,best_page,birthday,bio,booking_agent,built,business,can_checkin,category,category_list,can_post,checkins,company_overview,contact_address,country_page_likes,cover,culinary_team,access_token,app_id,current_location,description,description_html,directed_by,display_subtext,emails,engagement,fan_count,featured_video,food_styles,features,founded,general_info,general_manager,genre,global_brand_page_name,has_added_app,global_brand_root_id,hometown,hours,id,impressum,influences,is_always_open,is_community_page,is_permanently_closed,is_published,is_unclaimed,is_verified,last_used_time,leadgen_tos_accepted,link,location,members,mission,mpg,name_with_location_descriptor,network,new_like_count,offer_eligible,parent_page,parking,payment_options,personal_info,personal_interests,pharma_safety_info,phone,place_type,press_contact,plot_outline,price_range,produced_by,products,promotion_ineligible_reason,public_transit,record_label,release_date,restaurant_services,restaurant_specialties,schedule,screenplay_by,season,single_line_address,starring,start_info,store_location_descriptor,store_number,studio,talking_about_count,username,verification_status,unread_message_count,unread_notif_count,unseen_message_count,voip_info,website,were_here_count,written_by"
	PAGE_PROFILE_FIELDS_DISABLE = "ad_campaign,context,instant_articles_review_status,owner_business,promotion_eligible,supports_instant_articles"

	PAGE_FEED_FIELDS_ENABLE = "actions,admin_creator,application,call_to_action,child_attachments,caption,comments_mirroring_domain,coordinates,created_time,description,event,expanded_height,expanded_width,feed_targeting,from,full_picture,height,icon,id,is_expired,is_crossposting_eligible,instagram_eligibility,is_hidden,is_instagram_eligible,is_popular,is_published,is_spherical,link,message,message_tags,name,object_id,parent_id,permalink_url,picture,place,privacy,promotion_status,properties,scheduled_publish_time,shares,source,status_type,story,story_tags,subscribed,target,targeting,timeline_visibility,type,updated_time,via,width,with_tags"
	PAGE_FEED_CONNECTIONS = "comments.limit(1).summary(true),likes.limit(1).summary(true)"

	PAGE_REACTIONS_FIELDS_ENABLE = "reactions.type(NONE).limit(0).summary(total_count).as(reactions_none),reactions.type(LIKE).limit(0).summary(total_count).as(reactions_like),reactions.type(LOVE).limit(0).summary(total_count).as(reactions_love),reactions.type(WOW).limit(0).summary(total_count).as(reactions_wow),reactions.type(HAHA).limit(0).summary(total_count).as(reactions_haha),reactions.type(SAD).limit(0).summary(total_count).as(reactions_sad),reactions.type(ANGRY).limit(0).summary(total_count).as(reactions_angry),reactions.type(THANKFUL).limit(0).summary(total_count).as(reactions_thankful)"
)

const (
	URL_FACEBOOK_PROFILE = "https://graph.facebook.com/v2.6/%s?fields=%s&access_token=%s"
	URL_FACEBOOK_POSTS = "https://graph.facebook.com/v2.6/%s/posts?fields=%s,%s&limit=%s&access_token=%s"
	URL_FACEBOOK_POST_REACTIONS = "https://graph.facebook.com/v2.6/%s?fields=%s&access_token=%s"
)

func GetFBProfile(c *gin.Context) {
	userId := c.Param("userId")

	repo := &data.FBRepo{
		Agent: common.GetAgent(),
		Url:   common.UrlString(URL_FACEBOOK_PROFILE, userId, PAGE_PROFILE_FIELDS_ENABLE, common.GetFBToken()),
	}
	getProfile(c, repo)
}

func GetFBPosts(c *gin.Context) {
	userId := c.Param("userId")
	limit := c.DefaultQuery("limit", "10")
	repo := &data.FBRepo{
		Agent: common.GetAgent(),
		Url:   common.UrlString(URL_FACEBOOK_POSTS, userId, PAGE_FEED_FIELDS_ENABLE, PAGE_FEED_CONNECTIONS, limit, common.GetFBToken()),
	}
	getPosts(c, repo)
}

func GetFBPostReactions(c *gin.Context) {
	postId := c.Param("postId")
	repo := &data.FBRepo{
		Agent: common.GetAgent(),
		Url:   common.UrlString(URL_FACEBOOK_POST_REACTIONS, postId, PAGE_REACTIONS_FIELDS_ENABLE, common.GetFBToken()),
	}
	getReactions(c, repo)
}

func getReactions(c *gin.Context, repo *data.FBRepo) {
	timeout := c.DefaultQuery("timeout", "5")
	timer := common.Timeout(timeout)
	pCh := make(chan models.FBReactions)

	go fetchReactions(repo, pCh)

	var reactions models.FBReactions
	select {
	case reactions = <-pCh:
	case <-timer:
		reactions = TimeOutReactions()
	}
	c.JSON(http.StatusOK, gin.H{
		"like_counts": reactions,
	})
}

func fetchReactions(repo *data.FBRepo, ch chan models.FBReactions) {
	defer close(ch)
	var reactions models.FBReactions

	body, err := repo.FetchReactionsApi()
	if err != nil {
		reactions.FetchErr(err)
		ch <- reactions
		return
	}
	reactions = repo.ParseRawReactions(body)

	ch <- reactions
}

func TimeOutReactions() models.FBReactions {
	var r models.FBReactions
	r.ErrMessage = common.ERROR_MSG_API_TIMEOUT
	r.Date = common.Now()
	r.Status = false
	return r
}
