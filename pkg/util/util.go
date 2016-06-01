package util

import (
	"regexp"
	"fmt"
)

func Matcher(expr string, s string) []string {
	r, _ := regexp.Compile(expr)
	matcher := r.FindStringSubmatch(s)
	return matcher
}

func ToString(v interface{}) string {
	return fmt.Sprintf("%+v", v)
}
