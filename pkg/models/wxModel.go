package models

type WXRawProfile struct  {



}
type WXRawPosts struct {
	List []struct {
		AppMsgExtInfo struct {
			Author              string `json:"author"`
			Content             string `json:"content"`
			ContentURL          string `json:"content_url"`
			CopyrightStat       int    `json:"copyright_stat"`
			Cover               string `json:"cover"`
			Digest              string `json:"digest"`
			Fileid              int    `json:"fileid"`
			IsMulti             int    `json:"is_multi"`
			MultiAppMsgItemList []struct {
				Author        string `json:"author"`
				Content       string `json:"content"`
				ContentURL    string `json:"content_url"`
				CopyrightStat int    `json:"copyright_stat"`
				Cover         string `json:"cover"`
				Digest        string `json:"digest"`
				Fileid        int    `json:"fileid"`
				SourceURL     string `json:"source_url"`
				Title         string `json:"title"`
			} `json:"multi_app_msg_item_list"`
			SourceURL string `json:"source_url"`
			Subtype   int    `json:"subtype"`
			Title     string `json:"title"`
		} `json:"app_msg_ext_info"`
		CommMsgInfo struct {
			Content  string `json:"content"`
			Datetime int    `json:"datetime"`
			Fakeid   string `json:"fakeid"`
			ID       int    `json:"id"`
			Status   int    `json:"status"`
			Type     int    `json:"type"`
		} `json:"comm_msg_info"`
	} `json:"list"`
}
