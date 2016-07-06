package models

type WBRawProfile struct {
	FansScheme   string `json:"fans_scheme"`
	FollowScheme string `json:"follow_scheme"`
	TabsInfo     struct {
		SelectedTab int `json:"selectedTab"`
		Tabs        []struct {
			Containerid string `json:"containerid"`
			TabType     string `json:"tab_type"`
			Title       string `json:"title"`
		} `json:"tabs"`
	} `json:"tabsInfo"`
	UserInfo struct {
		AllowAllActMsg  bool   `json:"allow_all_act_msg"`
		AllowAllComment bool   `json:"allow_all_comment"`
		AllowMsg        int    `json:"allow_msg"`
		AvatarHd        string `json:"avatar_hd"`
		AvatarLarge     string `json:"avatar_large"`
		Badge           struct {
			Ali1688        int `json:"ali_1688"`
			Anniversary    int `json:"anniversary"`
			BindTaobao     int `json:"bind_taobao"`
			Dailv          int `json:"dailv"`
			Daiyan         int `json:"daiyan"`
			Enterprise     int `json:"enterprise"`
			FoolsDay2016   int `json:"fools_day_2016"`
			Gongyi         int `json:"gongyi"`
			GongyiLevel    int `json:"gongyi_level"`
			Hongbao2014    int `json:"hongbao_2014"`
			Suishoupai2014 int `json:"suishoupai_2014"`
			Suishoupai2016 int `json:"suishoupai_2016"`
			SuperStar2016  int `json:"super_star_2016"`
			Taobao         int `json:"taobao"`
			Travel2013     int `json:"travel2013"`
			UcDomain       int `json:"uc_domain"`
			UefaEuro2016   int `json:"uefa_euro_2016"`
			UnreadPool     int `json:"unread_pool"`
			VipActivity1   int `json:"vip_activity1"`
			VipActivity2   int `json:"vip_activity2"`
			Zongyiji       int `json:"zongyiji"`
		} `json:"badge"`
		BadgeTop             string `json:"badge_top"`
		BiFollowersCount     int    `json:"bi_followers_count"`
		Birthday             string `json:"birthday"`
		BlockApp             int    `json:"block_app"`
		BlockWord            int    `json:"block_word"`
		City                 string `json:"city"`
		Class                int    `json:"class"`
		CloseFriendsType     int    `json:"close_friends_type"`
		CoverImagePhone      string `json:"cover_image_phone"`
		CoverImagePhoneLevel int    `json:"cover_image_phone_level"`
		CreatedAt            string `json:"created_at"`
		CreditScore          int    `json:"credit_score"`
		Description          string `json:"description"`
		Domain               string `json:"domain"`
		Email                string `json:"email"`
		Extend               struct {
			Mbprivilege string `json:"mbprivilege"`
			Privacy     struct {
				Mobile int `json:"mobile"`
			} `json:"privacy"`
		} `json:"extend"`
		FavouritesCount     int    `json:"favourites_count"`
		FollowMe            bool   `json:"follow_me"`
		FollowersCount      int    `json:"followers_count"`
		Following           bool   `json:"following"`
		FriendsCount        int    `json:"friends_count"`
		FriendshipsRelation int    `json:"friendships_relation"`
		Gender              string `json:"gender"`
		GeoEnabled          bool   `json:"geo_enabled"`
		HasAbilityTag       int    `json:"has_ability_tag"`
		ID                  int    `json:"id"`
		Idstr               string `json:"idstr"`
		Lang                string `json:"lang"`
		Level               int    `json:"level"`
		Location            string `json:"location"`
		Mbrank              int    `json:"mbrank"`
		Mbtype              int    `json:"mbtype"`
		Msn                 string `json:"msn"`
		Name                string `json:"name"`
		OnlineStatus        int    `json:"online_status"`
		PagefriendsCount    int    `json:"pagefriends_count"`
		PicBg               string `json:"pic_bg"`
		ProfileImageURL     string `json:"profile_image_url"`
		ProfileURL          string `json:"profile_url"`
		Province            string `json:"province"`
		Ptype               int    `json:"ptype"`
		Qq                  string `json:"qq"`
		Remark              string `json:"remark"`
		ScreenName          string `json:"screen_name"`
		Star                int    `json:"star"`
		StatusID            int    `json:"status_id"`
		StatusesCount       int    `json:"statuses_count"`
		Type                int    `json:"type"`
		Ulevel              int    `json:"ulevel"`
		Urank               int    `json:"urank"`
		URL                 string `json:"url"`
		UserAbility         int    `json:"user_ability"`
		Verified            bool   `json:"verified"`
		VerifiedReason      string `json:"verified_reason"`
		VerifiedReasonURL   string `json:"verified_reason_url"`
		VerifiedSource      string `json:"verified_source"`
		VerifiedSourceURL   string `json:"verified_source_url"`
		VerifiedTrade       string `json:"verified_trade"`
		VerifiedType        int    `json:"verified_type"`
		Weihao              string `json:"weihao"`
	} `json:"userInfo"`
}

type WBRawPosts struct {
	CardGroup []struct {
		CardType int `json:"card_type"`
		Mblog    struct {
			AttitudesCount   int    `json:"attitudes_count"`
			AttitudesStatus  int    `json:"attitudes_status"`
			Bid              string `json:"bid"`
			BizFeature       int    `json:"biz_feature"`
			BmiddlePic       string `json:"bmiddle_pic"`
			CommentsCount    int    `json:"comments_count"`
			CreatedAt        string `json:"created_at"`
			CreatedTimestamp int    `json:"created_timestamp"`
			ExtendInfo       struct {
				WeiboCamera struct {
					C []string `json:"c"`
				} `json:"weibo_camera"`
			} `json:"extend_info"`
			Favorited    bool          `json:"favorited"`
			HotWeiboTags []interface{} `json:"hot_weibo_tags"`
			ID           int           `json:"id"`
			Idstr        string        `json:"idstr"`
			IsLongText   bool          `json:"isLongText"`
			LikeCount    int           `json:"like_count"`
			Mblogtype    int           `json:"mblogtype"`
			Mid          string        `json:"mid"`
			Mlevel       int           `json:"mlevel"`
			OriginalPic  string        `json:"original_pic"`
			PageType     int           `json:"page_type"`
			PicStatus    string        `json:"picStatus"`
			PicIds       []string      `json:"pic_ids"`
			Pics         []struct {
				Geo struct {
					Byte   int  `json:"byte"`
					Croped bool `json:"croped"`
					Height int  `json:"height"`
					Width  int  `json:"width"`
				} `json:"geo"`
				Pid  string `json:"pid"`
				Size string `json:"size"`
				URL  string `json:"url"`
			} `json:"pics"`
			PositiveRecomFlag int           `json:"positive_recom_flag"`
			RepostsCount      int           `json:"reposts_count"`
			Source            string        `json:"source"`
			SourceAllowclick  int           `json:"source_allowclick"`
			SourceType        int           `json:"source_type"`
			Text              string        `json:"text"`
			TextLength        int           `json:"textLength"`
			TextTagTips       []interface{} `json:"text_tag_tips"`
			ThumbnailPic      string        `json:"thumbnail_pic"`
			TopicStruct       []struct {
				TopicTitle string `json:"topic_title"`
				TopicURL   string `json:"topic_url"`
			} `json:"topic_struct"`
			User struct {
				Description     string      `json:"description"`
				FollowMe        bool        `json:"follow_me"`
				Following       bool        `json:"following"`
				Gender          string      `json:"gender"`
				ID              int         `json:"id"`
				Ismember        int         `json:"ismember"`
				Mbtype          int         `json:"mbtype"`
				ProfileImageURL string      `json:"profile_image_url"`
				ProfileURL      string      `json:"profile_url"`
				Remark          string      `json:"remark"`
				ScreenName      string      `json:"screen_name"`
				StatusesCount   int         `json:"statuses_count"`
				Valid           interface{} `json:"valid"`
				Verified        bool        `json:"verified"`
				VerifiedReason  string      `json:"verified_reason"`
				VerifiedType    int         `json:"verified_type"`
			} `json:"user"`
			UserType int `json:"userType"`
			Visible  struct {
				ListID int `json:"list_id"`
				Type   int `json:"type"`
			} `json:"visible"`
		} `json:"mblog"`
	} `json:"card_group"`
	LoadMore       bool   `json:"loadMore"`
	MaxPage        int    `json:"maxPage"`
	ModType        string `json:"mod_type"`
	NextCursor     string `json:"next_cursor"`
	Page           int    `json:"page"`
	PreviousCursor string `json:"previous_cursor"`
	URL            string `json:"url"`
}

type WBRawPost struct {
	CardType    int `json:"card_type"`
	Hideoperate int `json:"hideoperate"`
	Mblog       struct {
		Appid            int    `json:"appid"`
		AttitudesCount   int    `json:"attitudes_count"`
		AttitudesStatus  int    `json:"attitudes_status"`
		Bid              string `json:"bid"`
		BizFeature       int    `json:"biz_feature"`
		BizIds           []int  `json:"biz_ids"`
		CommentsCount    int    `json:"comments_count"`
		CreatedAt        string `json:"created_at"`
		CreatedTimestamp int    `json:"created_timestamp"`
		Favorited        bool   `json:"favorited"`
		HotWeiboTags     []struct {
			Containerid string  `json:"containerid"`
			FromCateid  string  `json:"from_cateid"`
			TagHidden   int     `json:"tag_hidden"`
			TagName     string  `json:"tag_name"`
			TagScheme   string  `json:"tag_scheme"`
			TagType     int     `json:"tag_type"`
			TagWeight   float64 `json:"tag_weight"`
			URLTypePic  string  `json:"url_type_pic"`
		} `json:"hot_weibo_tags"`
		ID         int    `json:"id"`
		Idstr      string `json:"idstr"`
		IsLongText bool   `json:"isLongText"`
		LikeCount  int    `json:"like_count"`
		Mid        string `json:"mid"`
		Mlevel     int    `json:"mlevel"`
		PageInfo   struct {
			Actionlog struct {
				ActCode int    `json:"act_code"`
				ActType int    `json:"act_type"`
				Ext     string `json:"ext"`
				Oid     string `json:"oid"`
			} `json:"actionlog"`
			Buttons []struct {
				Actionlog struct {
					ActCode int         `json:"act_code"`
					Cardid  interface{} `json:"cardid"`
					Ext     string      `json:"ext"`
					Fid     interface{} `json:"fid"`
					Oid     string      `json:"oid"`
				} `json:"actionlog"`
				Name   string `json:"name"`
				Params struct {
					Scheme string `json:"scheme"`
				} `json:"params"`
				Pic    string `json:"pic"`
				Scheme string `json:"scheme"`
				Type   string `json:"type"`
			} `json:"buttons"`
			Content1   string      `json:"content1"`
			Content2   string      `json:"content2"`
			ObjectID   string      `json:"object_id"`
			ObjectType string      `json:"object_type"`
			PageDesc   string      `json:"page_desc"`
			PageID     string      `json:"page_id"`
			PagePic    string      `json:"page_pic"`
			PageTitle  string      `json:"page_title"`
			PageURL    string      `json:"page_url"`
			Tips       string      `json:"tips"`
			Type       int         `json:"type"`
			TypeIcon   interface{} `json:"type_icon"`
		} `json:"page_info"`
		PageType          int           `json:"page_type"`
		PicIds            []interface{} `json:"pic_ids"`
		PositiveRecomFlag int           `json:"positive_recom_flag"`
		RawText           string        `json:"raw_text"`
		ReadsCount        int           `json:"reads_count"`
		RepostsCount      int           `json:"reposts_count"`
		Source            string        `json:"source"`
		SourceAllowclick  int           `json:"source_allowclick"`
		SourceType        int           `json:"source_type"`
		TagStruct         []struct {
			Oid        string `json:"oid"`
			TagHidden  int    `json:"tag_hidden"`
			TagName    string `json:"tag_name"`
			TagScheme  string `json:"tag_scheme"`
			TagType    int    `json:"tag_type"`
			URLTypePic string `json:"url_type_pic"`
		} `json:"tag_struct"`
		Text        string        `json:"text"`
		TextLength  int           `json:"textLength"`
		TextTagTips []interface{} `json:"text_tag_tips"`
		TopicStruct []struct {
			TopicTitle string `json:"topic_title"`
			TopicURL   string `json:"topic_url"`
		} `json:"topic_struct"`
		URLStruct []struct {
			Hide       int    `json:"hide"`
			Log        string `json:"log"`
			OriURL     string `json:"ori_url"`
			PageID     string `json:"page_id"`
			Result     bool   `json:"result"`
			ShortURL   string `json:"short_url"`
			URLTitle   string `json:"url_title"`
			URLType    int    `json:"url_type"`
			URLTypePic string `json:"url_type_pic"`
		} `json:"url_struct"`
		User struct {
			Description string `json:"description"`
			FansNum     string `json:"fansNum"`
			FollowMe    bool   `json:"follow_me"`
			Following   bool   `json:"following"`
			Gender      string `json:"gender"`
			H5icon      struct {
				Main  string   `json:"main"`
				Other []string `json:"other"`
			} `json:"h5icon"`
			ID              int         `json:"id"`
			Ismember        int         `json:"ismember"`
			Mbtype          int         `json:"mbtype"`
			ProfileImageURL string      `json:"profile_image_url"`
			ProfileURL      string      `json:"profile_url"`
			Remark          string      `json:"remark"`
			ScreenName      string      `json:"screen_name"`
			StatusesCount   int         `json:"statuses_count"`
			Valid           interface{} `json:"valid"`
			Verified        bool        `json:"verified"`
			VerifiedReason  string      `json:"verified_reason"`
			VerifiedType    int         `json:"verified_type"`
		} `json:"user"`
		UserType int `json:"userType"`
		Visible  struct {
			ListID int `json:"list_id"`
			Type   int `json:"type"`
		} `json:"visible"`
	} `json:"mblog"`
	Single int `json:"single"`
}
