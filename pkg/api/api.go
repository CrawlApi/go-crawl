package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"github.com/llitfkitfk/cirkol/pkg/util"
	"log"
	"net/http"
	"time"
)

const (
	REGEXP_URI = `(http|https)://(\S+).(\S+)`

	// FACEBOOK CONST
	REGEXP_FACEBOOK_PROFILE_ID = `fb://(page|profile|group)/(\d+)`

	PAGE_PROFILE_FIELDS_ENABLE = "name,about,affiliation,app_links,artists_we_like,attire,awards,band_interests,band_members,best_page,birthday,bio,booking_agent,built,business,can_checkin,category,category_list,can_post,checkins,company_overview,contact_address,country_page_likes,cover,culinary_team,access_token,app_id,current_location,description,description_html,directed_by,display_subtext,emails,engagement,fan_count,featured_video,food_styles,features,founded,general_info,general_manager,genre,global_brand_page_name,has_added_app,global_brand_root_id,hometown,hours,id,impressum,influences,is_always_open,is_community_page,is_permanently_closed,is_published,is_unclaimed,is_verified,last_used_time,leadgen_tos_accepted,link,location,members,mission,mpg,name_with_location_descriptor,network,new_like_count,offer_eligible,parent_page,parking,payment_options,personal_info,personal_interests,pharma_safety_info,phone,place_type,press_contact,plot_outline,price_range,produced_by,products,promotion_ineligible_reason,public_transit,record_label,release_date,restaurant_services,restaurant_specialties,schedule,screenplay_by,season,single_line_address,starring,start_info,store_location_descriptor,store_number,studio,talking_about_count,username,verification_status,unread_message_count,unread_notif_count,unseen_message_count,voip_info,website,were_here_count,written_by"
	PAGE_PROFILE_FIELDS_DISABLE = "ad_campaign,context,instant_articles_review_status,owner_business,promotion_eligible,supports_instant_articles"

	USER_PROFILE_FIELDS_ENABLE = "about,admin_notes,age_range,bio,birthday,cover,currency,devices,education,email,favorite_athletes,favorite_teams,first_name,gender,hometown,inspirational_people,install_type,installed,interested_in,id,is_verified,labels,languages,last_name,link,location,locale,meeting_for,middle_name,name,name_format,payment_pricepoints,public_key,political,quotes,relationship_status,religion,security_settings,significant_other,sports,test_group,timezone,third_party_id,updated_time,verified,video_upload_limits,viewer_can_send_gift,website,work"
	USER_PROFILE_FIELDS_DISABLE = "context,is_shared_login,shared_login_upgrade_required_by,token_for_business"

	PAGE_FEED_FIELDS_ENABLE = "actions,admin_creator,application,call_to_action,child_attachments,caption,comments_mirroring_domain,coordinates,created_time,description,event,expanded_height,expanded_width,feed_targeting,from,full_picture,height,icon,id,is_expired,is_crossposting_eligible,instagram_eligibility,is_hidden,is_instagram_eligible,is_popular,is_published,is_spherical,link,message,message_tags,name,object_id,parent_id,permalink_url,picture,place,privacy,promotion_status,properties,scheduled_publish_time,shares,source,status_type,story,story_tags,subscribed,target,targeting,timeline_visibility,type,updated_time,via,width,with_tags"
	PAGE_FEED_CONNECTIONS = "comments.limit(1).summary(true),likes.limit(1).summary(true)"
	PAGE_FEED_FIELDS_DISABLE = "allowed_advertising_objectives,entities,implicit_place,is_app_share"

	// INSTAGRAM CONST
	REGEX_INSTAGRAM_PROFILE_ID = `"owner": {"id": "(\d+)`
	REGEX_INSTAGRAM_PROFILE_NAME = `"user": {"username": "(\S+)",`
	REGEX_INSTAGRAM_PROFILE = `ProfilePage": \[([\s\S]+), "nodes": ([\s\S]+)]([\s\S]+)]},`
	REGEX_INSTAGRAM_POSTS = `ProfilePage": \[([\s\S]+), "nodes": ([\s\S]+)]([\s\S]+)]},`

	// WEIXIN CONST
	REGEXP_WEIXIN_PROFILE_ID = `微信号: (\S+)</p>`
	REGEXP_WEIXIN_LOGO = `src="((http://img01.sogoucdn.com/app/a)\S+)"`
	REGEXP_WEIXIN_FEATURE = `功能介绍(...+)class="sp-txt">(...+)</span>`
	REGEXP_WEIXIN_URL = `href="((http://mp.weixin.qq.com/profile)\S+)"`
	REGEXP_WEIXIN_POSTS = `var msgList = '(\S+)';`

	REGEXP_WEIBO_PROFILE_ID = `uid=(\d+)`
)

const (
	ERROR_CODE_API_MISS_MATCHED = 4001
	ERROR_CODE_API_TIMEOUT = 4002
	ERROR_CODE_JSON_ERROR = 4003
	ERROR_CODE_TIMEOUT = 4004
	ERROR_CODE_REGEX_MISS_MATCHED = 4005
	ERROR_CODE_URL_TYPE_NOT_FOUND = 4006

	ERROR_MSG_API_MISS_MATCHED = "no api matched"
	ERROR_MSG_API_TIMEOUT = "request api timeout"
	ERROR_MSG_JSON_ERROR = "json parse error"
	ERROR_MSG_TIMEOUT = "request timeout"
	ERROR_MSG_REGEX_MISS_MATCHED = "regex miss matched"
	ERROR_MSG_URL_MISS_MATCHED = "url miss matched"
)

var (
	FACEBOOK_TOKEN = "490895874437565|3ce74d840577a6d598af56cd46fd0450"
	INSTAGRAM_TOKEN = "28177225.e67f6b8.1a30e1aa29d44d4eb34d76dd128c7788"
	WEIBO_TOKEN = "2.00m9AuWD0IVHcF858d98077e0YDshC"
)

func ReqApi(url string) (string, error) {
	_, body, errs := reqClient.Set("accept-language", "en-US").Get(url).End()
	if errs != nil {
		return "", errors.New(ERROR_MSG_API_TIMEOUT)
	}
	return body, nil
}

func SearchOtherProfile(c *gin.Context, ch chan <- result.Profile) {
	var profile result.Profile
	profile.ErrCode = ERROR_CODE_API_MISS_MATCHED
	profile.ErrMessage = ERROR_MSG_API_MISS_MATCHED
	profile.Date = time.Now().Unix()
	ch <- profile
}
func SearchOtherPosts(c *gin.Context, ch chan <- result.Posts) {
	var posts result.Posts
	posts.ErrCode = ERROR_CODE_API_MISS_MATCHED
	posts.ErrMessage = ERROR_MSG_API_MISS_MATCHED
	posts.Date = time.Now().Unix()
	ch <- posts
}

func SearchOtherUID(c *gin.Context, ch chan <- result.UID) {
	var uid result.UID
	uid.ErrCode = ERROR_CODE_API_MISS_MATCHED
	uid.ErrMessage = ERROR_MSG_API_MISS_MATCHED
	uid.Date = time.Now().Unix()
	ch <- uid
}

func GetProfile(c *gin.Context) {
	timer := time.After(5 * time.Second)
	apiType := c.Param("type")
	pCh := make(chan result.Profile)

	switch apiType {
	case "fb":
		go SearchFBProfile(c, pCh)
	case "ig":
		go SearchIGProfile(c, pCh)
	case "wb":
		go SearchWBProfile(c, pCh)
	case "wx":
		go SearchWXProfile(c, pCh)
	default:
		go SearchOtherProfile(c, pCh)
	}

	select {
	case profile := <-pCh:
		c.JSON(http.StatusOK, gin.H{
			"profile": profile,
		})
	case <-timer:
		c.JSON(http.StatusOK, gin.H{
			"profile": result.Profile{
				ErrCode:    ERROR_CODE_TIMEOUT,
				ErrMessage: ERROR_MSG_TIMEOUT,
				Date:       time.Now().Unix(),
			},
		})
	}
}

func GetPosts(c *gin.Context) {
	timer := time.After(8 * time.Second)
	apiType := c.Param("type")
	pCh := make(chan result.Posts)

	switch apiType {
	case "fb":
		go SearchFBPosts(c, pCh)
	case "ig":
		go SearchIGPosts(c, pCh)
	case "wb":
	case "wx":
		go SearchWXPosts(c, pCh)
	default:
		go SearchOtherPosts(c, pCh)
	}

	select {
	case posts := <-pCh:
		posts.Date = time.Now().Unix()
		c.JSON(http.StatusOK, gin.H{
			"posts": posts,
		})
	case <-timer:
		c.JSON(http.StatusOK, gin.H{
			"posts": result.Posts{
				ErrCode:    ERROR_CODE_TIMEOUT,
				ErrMessage: ERROR_MSG_TIMEOUT,
				Date:       time.Now().Unix(),
			},
		})
	}

}

func GetUid(c *gin.Context) {
	timer := time.After(5 * time.Second)
	rawurl := c.PostForm("url")
	realUrl := util.Matcher(REGEXP_URI, rawurl)

	var uid result.UID
	if len(realUrl) == 0 {
		uid.ErrCode = ERROR_CODE_REGEX_MISS_MATCHED
		uid.ErrMessage = ERROR_MSG_URL_MISS_MATCHED
		uid.Date = time.Now().Unix()

		c.JSON(http.StatusOK, gin.H{
			"uid": uid,
		})
		return
	}

	apiType, err := util.CheckUrl(rawurl)
	if err != nil {
		uid.ErrCode = ERROR_CODE_URL_TYPE_NOT_FOUND
		uid.ErrMessage = err.Error()
		uid.Date = time.Now().Unix()
		c.JSON(http.StatusOK, gin.H{
			"uid": uid,
		})
		return
	}

	uidCh := make(chan result.UID)
	switch apiType {
	case "fb":
	//go SearchFBUID(c, uidCh)
	case "ig":
		go SearchIGUID(c, uidCh)
	case "wb":
	//go SearchWBUID(c, uidCh)
	case "wx":
		go SearchWXUID(c, uidCh)
	default:
		go SearchOtherUID(c, uidCh)
	}

	select {
	case uid = <-uidCh:
		uid.Date = time.Now().Unix()
		c.JSON(http.StatusOK, gin.H{
			"uid": uid,
		})
	case <-timer:
		c.JSON(http.StatusOK, gin.H{
			"uid": result.UID{
				ErrCode:    ERROR_CODE_TIMEOUT,
				ErrMessage: ERROR_MSG_TIMEOUT,
				Date:       time.Now().Unix(),
			},
		})
	}
}

func UpdateToken(c *gin.Context) {
	token := c.PostForm("token")
	log.Println(token)
	c.String(http.StatusOK, "success")
}
