package api

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/llitfkitfk/cirkol/pkg/util"
	"time"
)

const (
	REGEXP_URI = `(http|https)://(\S+).(\S+)`
	// FACEBOOK CONST
	REGEXP_FACEBOOK_PROFILE_ID = `fb://(page|profile|group)/(\d+)`
	FACEBOOK_TOKEN = "490895874437565|3ce74d840577a6d598af56cd46fd0450"
	PAGE_PROFILE_FIELDS_ENABLE = "name,about,affiliation,app_links,artists_we_like,attire,awards,band_interests,band_members,best_page,birthday,bio,booking_agent,built,business,can_checkin,category,category_list,can_post,checkins,company_overview,contact_address,country_page_likes,cover,culinary_team,access_token,app_id,current_location,description,description_html,directed_by,display_subtext,emails,engagement,fan_count,featured_video,food_styles,features,founded,general_info,general_manager,genre,global_brand_page_name,has_added_app,global_brand_root_id,hometown,hours,id,impressum,influences,is_always_open,is_community_page,is_permanently_closed,is_published,is_unclaimed,is_verified,last_used_time,leadgen_tos_accepted,link,location,members,mission,mpg,name_with_location_descriptor,network,new_like_count,offer_eligible,parent_page,parking,payment_options,personal_info,personal_interests,pharma_safety_info,phone,place_type,press_contact,plot_outline,price_range,produced_by,products,promotion_ineligible_reason,public_transit,record_label,release_date,restaurant_services,restaurant_specialties,schedule,screenplay_by,season,single_line_address,starring,start_info,store_location_descriptor,store_number,studio,talking_about_count,username,verification_status,unread_message_count,unread_notif_count,unseen_message_count,voip_info,website,were_here_count,written_by"
	PAGE_PROFILE_FIELDS_DISABLE = "ad_campaign,context,instant_articles_review_status,owner_business,promotion_eligible,supports_instant_articles"

	USER_PROFILE_FIELDS_ENABLE = "about,admin_notes,age_range,bio,birthday,cover,currency,devices,education,email,favorite_athletes,favorite_teams,first_name,gender,hometown,inspirational_people,install_type,installed,interested_in,id,is_verified,labels,languages,last_name,link,location,locale,meeting_for,middle_name,name,name_format,payment_pricepoints,public_key,political,quotes,relationship_status,religion,security_settings,significant_other,sports,test_group,timezone,third_party_id,updated_time,verified,video_upload_limits,viewer_can_send_gift,website,work"
	USER_PROFILE_FIELDS_DISABLE = "context,is_shared_login,shared_login_upgrade_required_by,token_for_business"

	PAGE_FEED_FIELDS_ENABLE = "actions,admin_creator,application,call_to_action,child_attachments,caption,comments_mirroring_domain,coordinates,created_time,description,event,expanded_height,expanded_width,feed_targeting,from,full_picture,height,icon,id,is_expired,is_crossposting_eligible,instagram_eligibility,is_hidden,is_instagram_eligible,is_popular,is_published,is_spherical,link,message,message_tags,name,object_id,parent_id,permalink_url,picture,place,privacy,promotion_status,properties,scheduled_publish_time,shares,source,status_type,story,story_tags,subscribed,target,targeting,timeline_visibility,type,updated_time,via,width,with_tags"
	PAGE_FEED_CONNECTIONS = "comments.limit(1).summary(true),likes.limit(1).summary(true)"
	PAGE_FEED_FIELDS_DISABLE = "allowed_advertising_objectives,entities,implicit_place,is_app_share"

	// INSTAGRAM CONST
	REGEX_INSTAGRAM_PROFILE_ID = `"owner": {"id": "(\d+)`
	INSTAGRAM_TOKEN = "28177225.c8a40c7.f120241729664093a4eba6a6bd079a94"

	REGEXP_WEIXIN_PROFILE_ID = `微信号: (\S+)</p>`
	REGEXP_WEIXIN_LOGO = `src="((http://img01.sogoucdn.com/app/a)\S+)"`
	REGEXP_WEIXIN_FEATURE = `功能介绍(\S+)\sclass="sp-txt">(\S+)</span>`
	REGEXP_WEIXIN_URL = `href="((http://mp.weixin.qq.com/profile)\S+)"`
	REGEXP_WEIXIN_POSTS = `var msgList = '(\S+)';`
)

func StringResponse(body string, c *gin.Context) {
	if len(body) > 0 {
		c.String(http.StatusOK, body)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Request API Failed",
			"status":  false,
			"date": time.Now().Unix(),
		})
	}
}

func ProfileResponse(ch chan Profile, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"profile": <-ch,
	})
}

//func PostResponse(ch chan Post, c *gin.Context)  {
//	c.JSON(http.StatusOK, gin.H{
//		"profile": <-ch,
//	})
//}

func GetWhichUid(c *gin.Context) {
	log.Println(c.PostForm("url"))
	rawurl := c.PostForm("url")
	realUrl := util.Matcher(REGEXP_URI, rawurl)
	if len(realUrl) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"uid":  &UID{
				Message: "not real url",
				Status: false,
				Date: time.Now().Unix(),
			},
		})
	} else {
		urlType, err := util.CheckUrl(rawurl)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"uid": &UID{
					Message: "URL Type Not Found",
					Status: false,
					Date: time.Now().Unix(),
				},
			})
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/api/" + urlType + "/uid")
		}
	}
}

func Response(ch chan UID, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"uid": <-ch,
	})
}

