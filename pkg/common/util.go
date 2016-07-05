package common

import (
	"encoding/json"
	"fmt"
	"html"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	//ERROR_MSG_API_FETCH = "request api timeout"
	//ERROR_MSG_API_MISS_MATCHED = "no api matched"
	ERROR_MSG_API_TIMEOUT        = "request api timeout"
	ERROR_MSG_JSON_ERROR         = "json parse error"
	ERROR_MSG_WX_POSTS_API_FETCH = "weixin fetch api error"
	ERROR_MSG_REGEX_MISS_MATCHED = "regex miss matched"
	ERROR_MSG_URL_NOT_MATCHED    = "Url not matched"
	//ERROR_MSG_TIMEOUT = "request timeout"
	//ERROR_MSG_URL_MISS_MATCHED = "url miss matched"
)

const (
	TIMESTAMP_LAYOUT    = "2006-01-02T15:04:05+0000"
	TIMESTAMP_LAYOUT_WB = "2006-1-2 15:04"
)

func UrlString(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func Int2Str(src int) string {
	return strconv.Itoa(src)
}

func Str2Int(src string) int {
	i, err := strconv.Atoi(src)
	if err != nil {
		return 0
	}
	return i
}

func Str2Int64(src string) int64 {
	i, err := strconv.Atoi(src)
	if err != nil {
		return 0
	}
	return int64(i)
}

func DateFormat(dateStr string) int64 {
	time, err := time.Parse(TIMESTAMP_LAYOUT, dateStr)
	if err != nil {
		return 0
	}
	return time.Unix()

}

func ParseJson(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}

func DecodeString(src string) string {
	return html.UnescapeString(src)
}

func Interface2String(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func Now() int64 {
	return time.Now().Unix()
}

func Matcher(expr string, s string) []string {
	r, _ := regexp.Compile(expr)
	return r.FindStringSubmatch(s)

}

func Timeout(d string) <-chan time.Time {
	i := time.Duration(query2Int(d))
	return time.After(i * time.Second)
}

func query2Int(src string) int {
	i, err := strconv.Atoi(src)
	if err != nil {
		return 5
	}
	return i
}

func ParseWBCreatedAt(dateStr string) int64 {
	today := time.Now()
	var resultStr string
	if strings.Contains(dateStr, "今天") {
		resultStr = fmt.Sprint(today.Year(), "-", int(today.Month()), "-", today.Day(), " ", dateStr[len(dateStr)-5:len(dateStr)])
	} else {
		resultStr = fmt.Sprint(today.Year(), "-", dateStr)
	}
	time, err := time.Parse(TIMESTAMP_LAYOUT_WB, resultStr)
	if err != nil {
		return 0
	}
	return time.Unix()
}

func GetMatcherValue(length int, expr, body string) string {
	matcher := Matcher(expr, body)
	if len(matcher) > length {
		return matcher[length]
	}
	return ""
}
