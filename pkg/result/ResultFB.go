package result

type FBRawProfile struct {
	About    string `json:"about"`
	AppLinks struct {
		Android []struct {
			AppName string `json:"app_name"`
			Package string `json:"package"`
			URL     string `json:"url"`
		} `json:"android"`
		Ios []struct {
			AppName    string `json:"app_name"`
			AppStoreID string `json:"app_store_id"`
			URL        string `json:"url"`
		} `json:"ios"`
	} `json:"app_links"`
	Birthday   string `json:"birthday"`
	CanCheckin bool   `json:"can_checkin"`
	CanPost    bool   `json:"can_post"`
	Category   string `json:"category"`
	Checkins   int    `json:"checkins"`
	Cover      struct {
		CoverID string `json:"cover_id"`
		ID      string `json:"id"`
		OffsetX int    `json:"offset_x"`
		OffsetY int    `json:"offset_y"`
		Source  string `json:"source"`
	} `json:"cover"`
	Engagement struct {
		Count          int    `json:"count"`
		SocialSentence string `json:"social_sentence"`
	} `json:"engagement"`
	FanCount                   int    `json:"fan_count"`
	GlobalBrandPageName        string `json:"global_brand_page_name"`
	HasAddedApp                bool   `json:"has_added_app"`
	Hometown                   string `json:"hometown"`
	ID                         string `json:"id"`
	IsAlwaysOpen               bool   `json:"is_always_open"`
	IsCommunityPage            bool   `json:"is_community_page"`
	IsPermanentlyClosed        bool   `json:"is_permanently_closed"`
	IsPublished                bool   `json:"is_published"`
	IsUnclaimed                bool   `json:"is_unclaimed"`
	IsVerified                 bool   `json:"is_verified"`
	LeadgenTosAccepted         bool   `json:"leadgen_tos_accepted"`
	Link                       string `json:"link"`
	Name                       string `json:"name"`
	NameWithLocationDescriptor string `json:"name_with_location_descriptor"`
	Parking                    struct {
		Lot    int `json:"lot"`
		Street int `json:"street"`
		Valet  int `json:"valet"`
	} `json:"parking"`
	StartInfo struct {
		Date struct {
			Day   int `json:"day"`
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"date"`
		Type string `json:"type"`
	} `json:"start_info"`
	TalkingAboutCount  int    `json:"talking_about_count"`
	Username           string `json:"username"`
	VerificationStatus string `json:"verification_status"`
	VoipInfo           struct {
		HasMobileApp      bool   `json:"has_mobile_app"`
		HasPermission     bool   `json:"has_permission"`
		IsCallable        bool   `json:"is_callable"`
		IsCallableWebrtc  bool   `json:"is_callable_webrtc"`
		IsPushable        bool   `json:"is_pushable"`
		ReasonCode        int    `json:"reason_code"`
		ReasonDescription string `json:"reason_description"`
	} `json:"voip_info"`
	Website       string `json:"website"`
	WereHereCount int    `json:"were_here_count"`
}

type FBRawPosts struct {
	Data []struct {
		Comments struct {
			Data []struct {
				CreatedTime string `json:"created_time"`
				From        struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"from"`
				ID      string `json:"id"`
				Message string `json:"message"`
			} `json:"data"`
			Paging struct {
				Cursors struct {
					After  string `json:"after"`
					Before string `json:"before"`
				} `json:"cursors"`
				Next string `json:"next"`
			} `json:"paging"`
			Summary struct {
				CanComment bool   `json:"can_comment"`
				Order      string `json:"order"`
				TotalCount int    `json:"total_count"`
			} `json:"summary"`
		} `json:"comments"`
		CreatedTime string `json:"created_time"`
		From        struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"from"`
		FullPicture          string `json:"full_picture"`
		Icon                 string `json:"icon"`
		ID                   string `json:"id"`
		InstagramEligibility string `json:"instagram_eligibility"`
		IsExpired            bool   `json:"is_expired"`
		IsHidden             bool   `json:"is_hidden"`
		IsInstagramEligible  bool   `json:"is_instagram_eligible"`
		IsPublished          bool   `json:"is_published"`
		IsSpherical          bool   `json:"is_spherical"`
		Likes                struct {
			Data []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"data"`
			Paging struct {
				Cursors struct {
					After  string `json:"after"`
					Before string `json:"before"`
				} `json:"cursors"`
				Next string `json:"next"`
			} `json:"paging"`
			Summary struct {
				CanLike    bool `json:"can_like"`
				HasLiked   bool `json:"has_liked"`
				TotalCount int  `json:"total_count"`
			} `json:"summary"`
		} `json:"likes"`
		Link         string `json:"link"`
		Message      string `json:"message"`
		Name         string `json:"name"`
		ObjectID     string `json:"object_id"`
		PermalinkURL string `json:"permalink_url"`
		Picture      string `json:"picture"`
		Privacy      struct {
			Allow       string `json:"allow"`
			Deny        string `json:"deny"`
			Description string `json:"description"`
			Friends     string `json:"friends"`
			Value       string `json:"value"`
		} `json:"privacy"`
		PromotionStatus string `json:"promotion_status"`
		Shares          struct {
			Count int `json:"count"`
		} `json:"shares"`
		StatusType         string `json:"status_type"`
		TimelineVisibility string `json:"timeline_visibility"`
		Type               string `json:"type"`
		UpdatedTime        string `json:"updated_time"`
	} `json:"data"`
	Paging struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
	} `json:"paging"`
}

type FBRawUid struct {
}
