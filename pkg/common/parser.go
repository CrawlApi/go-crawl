package common

import (
	"strings"
)

type Parser struct {
	RawUrl   string
	PostType string
}

func NewParser(rawUrl string) *Parser {
	pt := GetMatcherValue(2, `facebook.com/(\S+)/(photos|videos|posts)/`, rawUrl)
	return &Parser{RawUrl: strings.TrimSpace(rawUrl), PostType: pt}
}

func (p *Parser) ParseUIDLink() string {
	return UrlString(`https://www.facebook.com/%s`, GetMatcherValue(1, `facebook.com/(\S+)/(photos|videos|posts)`, p.RawUrl))
}

func (p *Parser) ParsePostSuffId() string {
	switch p.PostType {
	case "videos":
		return GetMatcherValue(3, `facebook.com/(\S+)/videos/(...+)/(\d+)`, p.RawUrl)
	case "photos":
		return GetMatcherValue(3, `facebook.com/(\S+)/photos/(...+)/(\d+)`, p.RawUrl)
	case "posts":
		return GetMatcherValue(2, `facebook.com/(\S+)/posts/(\d+)`, p.RawUrl)
	}
	return ""

}
