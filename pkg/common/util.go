package common

import (
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
	"fmt"
)

const TIMESTAMP_LAYOUT = "2006-01-02T15:04:05+0000"

func Timer(c *gin.Context) <-chan time.Time {
	//timeout := Str2Int(c.DefaultQuery("timeout", "5"))
	return time.After(5 * time.Second)
}

func Str2Int(src string) int {
	i, err := strconv.Atoi(src)
	if err != nil {
		return 5
	}
	return i
}

func DateFormat(dateStr string) string {
	time, err := time.Parse(TIMESTAMP_LAYOUT, dateStr)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%d", time.Unix())

}

func JsonToString(data []byte, err error) string {
	if err != nil {
		return err.Error()
	}
	return string(data)
}