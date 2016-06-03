package util

import (
	"regexp"
	"fmt"
	"errors"
)

const REGEX_URL_TYPE = `(facebook|instagram|weixin|weibo)`

func Matcher(expr string, s string) []string {
	r, _ := regexp.Compile(expr)
	matcher := r.FindStringSubmatch(s)
	return matcher
}

func ToString(v interface{}) string {
	return fmt.Sprintf("%+v", v)
}

func CheckUrl(url string) (string, error) {
	r, _ := regexp.Compile(REGEX_URL_TYPE)
	matcher := r.FindStringSubmatch(url)
	if len(matcher) > 0 {
		switch matcher[0] {
		case "facebook":
			return "fb", nil
		case "instagram":
			return "ig", nil
		case "weixin":
			return "wx", nil
		case "weibo":
			return "wb", nil
		}
	}
	return "", errors.New("Not Found")
}