package util

import (
	"regexp"
	"fmt"
	"errors"
	"strconv"
	"time"
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

func DateFormat(dateStr string) int64 {
	time, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return 0
	}
	return time.Unix()

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