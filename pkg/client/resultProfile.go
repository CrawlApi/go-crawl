package client

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"github.com/llitfkitfk/cirkol/pkg/parser"
	"strings"
)

func (r *Result) GetFBProfile() (models.Profile, error) {
	var profile models.Profile

	if r.err != nil {
		return profile, r.err
	}

	var rawData models.FBRawProfile
	err := common.ParseJson(r.Body, &rawData)
	if err != nil {
		return profile, err
	}

	profile.ParseFBProfile(rawData)

	return profile, nil
}

func (r *Result) GetWBProfile() (models.Profile, error) {
	var profile models.Profile
	if r.err != nil {
		return profile, r.err
	}

	var data models.WBRawProfile
	err := common.ParseJson(r.Body, &data)
	if err != nil {
		return profile, err
	}

	profile.ParseWBProfile(data)

	return profile, nil
}

func (r *Result) GetIGProfile() (models.Profile, error) {
	var profile models.Profile
	if r.err != nil {
		return profile, r.err
	}

	var data models.IGRawProfile
	dStr := parser.ParseIGProfile(r.Body)
	err := common.ParseJson(dStr, &data)
	if err != nil {
		return profile, err
	}

	profile.ParseIGProfile(data)

	return profile, nil
}

func (r *Result) GetIGV2Profile() (models.Profile, error) {
	var profile models.Profile
	if r.err != nil {
		return profile, r.err
	}

	var data models.IGV2RawProfile
	err := common.ParseJson(r.Body, &data)
	if err != nil {
		return profile, err
	}
	profile.ParseIGV2Profile(data)

	return profile, nil
}

func (r *Result) GetWXProfile() (models.Profile, error) {
	var profile models.Profile
	if r.err != nil {
		return profile, r.err
	}
	var data models.WXRawProfile
	data.Name = parser.ParseWXName(r.Body)
	data.Website = common.DecodeString(parser.ParseWXWeb(r.Body))
	data.Avatar = parser.ParseWXAvatar(r.Body)
	data.About = parser.ParseWXAbout(r.Body)

	profile.ParseWXProfile(data)
	return profile, nil
}

func (r *Result) GetYTBProfile() (models.Profile, error) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(r.Body))

	var profile models.Profile

	uid, _ := doc.Selection.Find(".appbar-content-hidable").Find("a").Attr("href")
	if len(uid) > 6 {
		profile.UserId = uid[6:]
	}
	profile.Website = common.UrlString("https://www.youtube.com%s", uid)

	profile.Name, _ = doc.Selection.Find(".primary-header-contents").Find("a").Attr("title")
	profile.Avatar, _ = doc.Selection.Find(".channel-header-profile-image").Attr("src")

	doc.Selection.Find(".about-stats").Find(".about-stat").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			profile.Fans = common.Str2Int(common.Replace(s.Find("b").Text(), ",", ""))
		case 1:
			profile.ViewCount = common.Str2Int(common.Replace(s.Find("b").Text(), ",", ""))
		case 2:
			profile.Birthday = s.Text()[7:]
		}
	})

	profile.About = doc.Selection.Find(".about-description").Text()
	profile.Status = true
	profile.Date = common.Now()

	return profile, nil
}
