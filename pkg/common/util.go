package common

import (
	"encoding/json"
	"fmt"
	"html"
	"regexp"
	"strconv"
	"time"
)

const (
	ERROR_CODE_API_FETCH = 4001
	ERROR_CODE_API_TIMEOUT = 4002
	ERROR_CODE_JSON_ERROR = 4003
	//ERROR_CODE_API_MISS_MATCHED = 4001
	//ERROR_CODE_TIMEOUT = 4004
	//ERROR_CODE_REGEX_MISS_MATCHED = 4005
	//ERROR_CODE_URL_TYPE_NOT_FOUND = 4006

	//ERROR_MSG_API_FETCH = "request api timeout"
	//ERROR_MSG_API_MISS_MATCHED = "no api matched"
	ERROR_MSG_API_TIMEOUT = "request api timeout"
	ERROR_MSG_JSON_ERROR = "json parse error"
	ERROR_MSG_WX_POSTS_API_FETCH = "weixin fetch api error"
	ERROR_MSG_REGEX_MISS_MATCHED = "regex miss matched"
	//ERROR_MSG_TIMEOUT = "request timeout"
	//ERROR_MSG_URL_MISS_MATCHED = "url miss matched"
)

const (
	TIMESTAMP_LAYOUT = "2006-01-02T15:04:05+0000"
)

func UrlString(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func Str2Int(src string) int {
	i, err := strconv.Atoi(src)
	if err != nil {
		return 5
	}
	return i
}

func Int2Str(src int) string {
	return strconv.Itoa(src)
}

func DateFormat(dateStr string) string {
	time, err := time.Parse(TIMESTAMP_LAYOUT, dateStr)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%d", time.Unix())

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
	i := time.Duration(Str2Int(d))
	return time.After(i * time.Second)
}
