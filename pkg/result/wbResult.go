package result

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


