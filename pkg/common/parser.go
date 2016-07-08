package common

import (
	"strings"
)

type Parser struct {
	RawUrl string
}

func NewParser(rawUrl string) *Parser {
	return &Parser{RawUrl: strings.TrimSpace(rawUrl)}
}

func (p *Parser) parseUIDLink() string {
	return UrlString(`https://www.facebook.com/%s`, GetMatcherValue(1, `facebook.com/(\S+)/posts`, p.RawUrl))
}

func (p *Parser) parsePostSuffId() string {
	return GetMatcherValue(2, `facebook.com/(\S+)/posts/(\d+)`, p.RawUrl)
}

func ParseUIDLink(rawUrl string) string {
	p := NewParser(rawUrl)
	return p.parseUIDLink()
}

func ParsePostSuffId(rawUrl string) string {
	p := NewParser(rawUrl)
	return p.parsePostSuffId()
}
