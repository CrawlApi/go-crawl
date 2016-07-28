package common

import (
	"encoding/json"
	"fmt"
	"html"
	"strconv"
	"strings"
	"time"
)

const (
	TIMESTAMP_LAYOUT    = "2006-01-02T15:04:05+0000"
	TIMESTAMP_LAYOUT_WB = "2006-1-2 15:4"
	TIMESTAMP_LAYOUT_YTB = "Jan _2, 2006"
)

func Replace(src, old, new string) string {
	return strings.Replace(src, old, new, -1)
}

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

func DateFormatYTB(dateStr string) int64  {
	time, err := time.Parse(TIMESTAMP_LAYOUT_YTB, dateStr)
	if err != nil {
		return 0
	}
	return time.Unix()
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
	if strings.Contains(dateStr, "分钟前") {
		mAgo := Str2Int(dateStr[:len(dateStr)-9])
		var h int
		var m int
		if today.Minute() >= mAgo {
			h = today.Hour()
			m = today.Minute() - mAgo
		} else {
			m = mAgo - today.Minute()
			h = today.Hour() - 1
		}
		resultStr = fmt.Sprint(today.Year(), "-", int(today.Month()), "-", today.Day(), " ", h, ":", m)
	} else if strings.Contains(dateStr, "今天") {
		resultStr = fmt.Sprint(today.Year(), "-", int(today.Month()), "-", today.Day(), " ", dateStr[len(dateStr)-5:len(dateStr)])
	} else {
		resultStr = fmt.Sprint(today.Year(), "-", dateStr)
	}
	date, err := time.Parse(TIMESTAMP_LAYOUT_WB, resultStr)
	if err != nil {
		return 0
	}
	return date.Unix()
}
