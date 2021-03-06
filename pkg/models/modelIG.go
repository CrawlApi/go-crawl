package models

type IGRawProfile struct {
	User struct {
		Username string `json:"username"`
		Follows  struct {
			Count int `json:"count"`
		} `json:"follows"`
		RequestedByViewer bool `json:"requested_by_viewer"`
		FollowedBy        struct {
			Count int `json:"count"`
		} `json:"followed_by"`
		CountryBlock           interface{} `json:"country_block"`
		HasRequestedViewer     bool        `json:"has_requested_viewer"`
		ExternalURLLinkshimmed interface{} `json:"external_url_linkshimmed"`
		FollowsViewer          bool        `json:"follows_viewer"`
		ProfilePicURL          string      `json:"profile_pic_url"`
		ID                     string      `json:"id"`
		Biography              string      `json:"biography"`
		FullName               string      `json:"full_name"`
		Media                  struct {
			Count    int `json:"count"`
			PageInfo struct {
				HasPreviousPage bool   `json:"has_previous_page"`
				StartCursor     string `json:"start_cursor"`
				EndCursor       string `json:"end_cursor"`
				HasNextPage     bool   `json:"has_next_page"`
			} `json:"page_info"`
		} `json:"media"`
		BlockedByViewer  bool        `json:"blocked_by_viewer"`
		FollowedByViewer bool        `json:"followed_by_viewer"`
		IsVerified       bool        `json:"is_verified"`
		HasBlockedViewer bool        `json:"has_blocked_viewer"`
		IsPrivate        bool        `json:"is_private"`
		ExternalURL      interface{} `json:"external_url"`
	} `json:"user"`
}

type IGRawPosts struct {
	Items []struct {
		AltMediaURL       string `json:"alt_media_url"`
		CanDeleteComments bool   `json:"can_delete_comments"`
		CanViewComments   bool   `json:"can_view_comments"`
		Caption           struct {
			CreatedTime string `json:"created_time"`
			From        struct {
				FullName       string `json:"full_name"`
				ID             string `json:"id"`
				ProfilePicture string `json:"profile_picture"`
				Username       string `json:"username"`
			} `json:"from"`
			ID   string `json:"id"`
			Text string `json:"text"`
		} `json:"caption"`
		Code     string `json:"code"`
		Comments struct {
			Count int `json:"count"`
			Data  []struct {
				CreatedTime string `json:"created_time"`
				From        struct {
					FullName       string `json:"full_name"`
					ID             string `json:"id"`
					ProfilePicture string `json:"profile_picture"`
					Username       string `json:"username"`
				} `json:"from"`
				ID   string `json:"id"`
				Text string `json:"text"`
			} `json:"data"`
		} `json:"comments"`
		CreatedTime string `json:"created_time"`
		ID          string `json:"id"`
		Images      struct {
			LowResolution struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"low_resolution"`
			StandardResolution struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"standard_resolution"`
			Thumbnail struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"thumbnail"`
		} `json:"images"`
		Likes struct {
			Count int `json:"count"`
			Data  []struct {
				FullName       string `json:"full_name"`
				ID             string `json:"id"`
				ProfilePicture string `json:"profile_picture"`
				Username       string `json:"username"`
			} `json:"data"`
		} `json:"likes"`
		Link     string      `json:"link"`
		Location interface{} `json:"location"`
		Type     string      `json:"type"`
		User     struct {
			FullName       string `json:"full_name"`
			ID             string `json:"id"`
			ProfilePicture string `json:"profile_picture"`
			Username       string `json:"username"`
		} `json:"user"`
		UserHasLiked bool `json:"user_has_liked"`
		VideoViews   int  `json:"video_views"`
		Videos       struct {
			LowBandwidth struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"low_bandwidth"`
			LowResolution struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"low_resolution"`
			StandardResolution struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"standard_resolution"`
		} `json:"videos"`
	} `json:"items"`
	MoreAvailable bool   `json:"more_available"`
	Status        string `json:"status"`
}

type IGV2RawProfile struct {
	Status string `json:"status"`
	User   struct {
		Biography                  string `json:"biography"`
		ExternalURL                string `json:"external_url"`
		FollowerCount              int    `json:"follower_count"`
		FollowingCount             int    `json:"following_count"`
		FullName                   string `json:"full_name"`
		HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
		HdProfilePicURLInfo        struct {
			Height int    `json:"height"`
			URL    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"hd_profile_pic_url_info"`
		IsBusiness    bool   `json:"is_business"`
		IsPrivate     bool   `json:"is_private"`
		MediaCount    int    `json:"media_count"`
		Pk            int    `json:"pk"`
		ProfilePicURL string `json:"profile_pic_url"`
		Username      string `json:"username"`
		UsertagsCount int    `json:"usertags_count"`
	} `json:"user"`
}

type IGV2RawPosts struct {
	Nodes []struct {
		Code       string `json:"code"`
		Date       int    `json:"date"`
		Dimensions struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"dimensions"`
		Comments struct {
			Count int `json:"count"`
		} `json:"comments"`
		Caption string `json:"caption"`
		Likes   struct {
			Count int `json:"count"`
		} `json:"likes"`
		Owner struct {
			ID string `json:"id"`
		} `json:"owner"`
		ThumbnailSrc string `json:"thumbnail_src"`
		IsVideo      bool   `json:"is_video"`
		ID           string `json:"id"`
		DisplaySrc   string `json:"display_src"`
		VideoViews   int    `json:"video_views,omitempty"`
	} `json:"nodes"`
}

type IGRawPost struct {
	Config struct {
		CsrfToken string      `json:"csrf_token"`
		Viewer    interface{} `json:"viewer"`
	} `json:"config"`
	CountryCode                  string `json:"country_code"`
	DisplayPropertiesServerGuess struct {
		PixelRatio    float64 `json:"pixel_ratio"`
		ViewportWidth int     `json:"viewport_width"`
	} `json:"display_properties_server_guess"`
	EntryData struct {
		PostPage []struct {
			Media struct {
				Caption         string `json:"caption"`
				CaptionIsEdited bool   `json:"caption_is_edited"`
				Code            string `json:"code"`
				Comments        struct {
					Count int `json:"count"`
					Nodes []struct {
						ID   string `json:"id"`
						Text string `json:"text"`
						User struct {
							ID            string `json:"id"`
							ProfilePicURL string `json:"profile_pic_url"`
							Username      string `json:"username"`
						} `json:"user"`
					} `json:"nodes"`
					PageInfo struct {
						EndCursor       interface{} `json:"end_cursor"`
						HasNextPage     bool        `json:"has_next_page"`
						HasPreviousPage bool        `json:"has_previous_page"`
						StartCursor     string      `json:"start_cursor"`
					} `json:"page_info"`
				} `json:"comments"`
				Date       int `json:"date"`
				Dimensions struct {
					Height int `json:"height"`
					Width  int `json:"width"`
				} `json:"dimensions"`
				DisplaySrc string `json:"display_src"`
				ID         string `json:"id"`
				IsAd       bool   `json:"is_ad"`
				IsVideo    bool   `json:"is_video"`
				Likes      struct {
					Count int `json:"count"`
					Nodes []struct {
						User struct {
							ID            string `json:"id"`
							ProfilePicURL string `json:"profile_pic_url"`
							Username      string `json:"username"`
						} `json:"user"`
					} `json:"nodes"`
					ViewerHasLiked bool `json:"viewer_has_liked"`
				} `json:"likes"`
				Location interface{} `json:"location"`
				Owner    struct {
					BlockedByViewer   bool   `json:"blocked_by_viewer"`
					FollowedByViewer  bool   `json:"followed_by_viewer"`
					FullName          string `json:"full_name"`
					HasBlockedViewer  bool   `json:"has_blocked_viewer"`
					ID                string `json:"id"`
					IsPrivate         bool   `json:"is_private"`
					IsUnpublished     bool   `json:"is_unpublished"`
					ProfilePicURL     string `json:"profile_pic_url"`
					RequestedByViewer bool   `json:"requested_by_viewer"`
					Username          string `json:"username"`
				} `json:"owner"`
				Usertags struct {
					Nodes []interface{} `json:"nodes"`
				} `json:"usertags"`
				VideoURL   string `json:"video_url"`
				VideoViews int    `json:"video_views"`
			} `json:"media"`
		} `json:"PostPage"`
	} `json:"entry_data"`
	EnvironmentSwitcherVisibleServerGuess bool `json:"environment_switcher_visible_server_guess"`
	Gatekeepers                           struct {
		Sulgin bool `json:"sulgin"`
		Vvc    bool `json:"vvc"`
	} `json:"gatekeepers"`
	Hostname     string `json:"hostname"`
	LanguageCode string `json:"language_code"`
	Platform     string `json:"platform"`
	Qe           struct {
		Discovery struct {
			G string   `json:"g"`
			P struct{} `json:"p"`
		} `json:"discovery"`
		Su struct {
			G string `json:"g"`
			P struct {
				Enabled string `json:"enabled"`
			} `json:"p"`
		} `json:"su"`
		SuUniverse struct {
			G string   `json:"g"`
			P struct{} `json:"p"`
		} `json:"su_universe"`
	} `json:"qe"`
	Qs             string `json:"qs"`
	ShowAppInstall bool   `json:"show_app_install"`
	StaticRoot     string `json:"static_root"`
}
