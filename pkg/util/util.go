package util

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
	"html"
)

const (
	REGEX_URL_TYPE = `(facebook|instagram|weixin|weibo)`
	TIMESTAMP_LAYOUT = "2006-01-02T15:04:05+0000"
	ERROR_MSG_URL_TYPE_NOT_FOUND = "URL Type Not Found"
)

func MatchString(length int, expr string, src string) string {
	r, _ := regexp.Compile(expr)
	matcher := r.FindStringSubmatch(src)
	if len(matcher) > length {
		return matcher[length + 1]
	}
	return ""
}

func Matcher(expr string, s string) []string {
	r, _ := regexp.Compile(expr)
	matcher := r.FindStringSubmatch(s)
	return matcher
}

func DecodeString(src string) string {
	return html.UnescapeString(src)
}

func ToString(v interface{}) string {
	return fmt.Sprintf("%+v", v)
}

func JsonToString(data []byte, err error) string {
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func Int2Str(src int) string {
	return strconv.Itoa(src)
}

func Str2Int(src string) int64 {
	i, err := strconv.Atoi(src)
	if err != nil {
		return 0
	}
	return int64(i)
}

func DateFormat(dateStr string) string {
	time, err := time.Parse(TIMESTAMP_LAYOUT, dateStr)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%d", time.Unix())

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
	return "", errors.New(ERROR_MSG_URL_TYPE_NOT_FOUND)
}
