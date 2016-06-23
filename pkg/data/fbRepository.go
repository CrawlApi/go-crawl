package data

import (
	"github.com/parnurzeal/gorequest"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"time"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	ERROR_MSG_API_TIMEOUT = "request api timeout"

	PAGE_PROFILE_FIELDS_ENABLE = "name,about,affiliation,app_links,artists_we_like,attire,awards,band_interests,band_members,best_page,birthday,bio,booking_agent,built,business,can_checkin,category,category_list,can_post,checkins,company_overview,contact_address,country_page_likes,cover,culinary_team,access_token,app_id,current_location,description,description_html,directed_by,display_subtext,emails,engagement,fan_count,featured_video,food_styles,features,founded,general_info,general_manager,genre,global_brand_page_name,has_added_app,global_brand_root_id,hometown,hours,id,impressum,influences,is_always_open,is_community_page,is_permanently_closed,is_published,is_unclaimed,is_verified,last_used_time,leadgen_tos_accepted,link,location,members,mission,mpg,name_with_location_descriptor,network,new_like_count,offer_eligible,parent_page,parking,payment_options,personal_info,personal_interests,pharma_safety_info,phone,place_type,press_contact,plot_outline,price_range,produced_by,products,promotion_ineligible_reason,public_transit,record_label,release_date,restaurant_services,restaurant_specialties,schedule,screenplay_by,season,single_line_address,starring,start_info,store_location_descriptor,store_number,studio,talking_about_count,username,verification_status,unread_message_count,unread_notif_count,unseen_message_count,voip_info,website,were_here_count,written_by"
	PAGE_PROFILE_FIELDS_DISABLE = "ad_campaign,context,instant_articles_review_status,owner_business,promotion_eligible,supports_instant_articles"

	PAGE_FEED_FIELDS_ENABLE = "actions,admin_creator,application,call_to_action,child_attachments,caption,comments_mirroring_domain,coordinates,created_time,description,event,expanded_height,expanded_width,feed_targeting,from,full_picture,height,icon,id,is_expired,is_crossposting_eligible,instagram_eligibility,is_hidden,is_instagram_eligible,is_popular,is_published,is_spherical,link,message,message_tags,name,object_id,parent_id,permalink_url,picture,place,privacy,promotion_status,properties,scheduled_publish_time,shares,source,status_type,story,story_tags,subscribed,target,targeting,timeline_visibility,type,updated_time,via,width,with_tags"
	PAGE_FEED_CONNECTIONS = "comments.limit(1).summary(true),likes.limit(1).summary(true)"
)

type FBRepo struct {
	Agent *gorequest.SuperAgent
}

func (r *FBRepo) GetAgent()  {
}

func (r *FBRepo) GetApi(url string) *gorequest.SuperAgent {
	return r.Agent.Timeout(8 * time.Second).Set("accept-language", "en-US").Get(url)
}

func (r *FBRepo) GetRawProfile(c *gin.Context) (models.FBRawProfile, error) {
	userId := c.Param("userId")
	url := "https://graph.facebook.com/v2.6/" + userId + "?fields=" + PAGE_PROFILE_FIELDS_ENABLE + "&access_token=" + common.GetFBToken()
	_, body, errs := r.GetApi(url).End()
	var rawProfile models.FBRawProfile

	if errs != nil {
		return rawProfile, errors.New(ERROR_MSG_API_TIMEOUT)
	}
	err := json.Unmarshal([]byte(body), &rawProfile)
	if err != nil {
		return rawProfile, err
	}
	return rawProfile, nil
}

func (r *FBRepo) GetRawPosts(c *gin.Context) (models.FBRawPosts, error) {
	userId := c.Param("userId")
	limit := c.DefaultQuery("limit", "10")
	url := "https://graph.facebook.com/v2.6/" + userId + "/feed?fields=" + PAGE_FEED_FIELDS_ENABLE + "," + PAGE_FEED_CONNECTIONS + "&limit=" + limit + "&access_token=" + common.GetFBToken()
	_, body, errs := r.GetApi(url).End()
	var rawPosts models.FBRawPosts

	if errs != nil {
		return rawPosts, errors.New(ERROR_MSG_API_TIMEOUT)
	}
	err := json.Unmarshal([]byte(body), &rawPosts)
	if err != nil {
		return rawPosts, err
	}
	return rawPosts, nil
}