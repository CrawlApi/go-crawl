package api

import (
	"github.com/llitfkitfk/cirkol/pkg/result"
	"encoding/json"
)

func GetPostsApi(url string, ch chan <- result.Posts) string {

	body, errs := GetApi(url)
	if errs != nil {
		ch <- result.Posts{
			ErrCode:    ERROR_CODE_API_TIMEOUT,
			ErrMessage: ERROR_MSG_API_TIMEOUT,
		}
		return ""
	}
	return body
}

func ParsePostsJson(src string, v interface{}, ch chan <- result.Posts) {
	err := json.Unmarshal([]byte(src), v)
	if err != nil {
		ch <- result.Posts{
			RawData:    src,
			ErrCode:    ERROR_CODE_JSON_ERROR,
			ErrMessage: err.Error(),
		}
	}
}

func MatchStrPostCh(length int, expr string, src string, ch chan<- result.Posts) string {

	des := MatchStrNoCh(length, expr, src)
	if des != "" {
		return des
	} else {
		ch <- result.Posts{
			ErrCode:    ERROR_CODE_REGEX_MISS_MATCHED,
			ErrMessage: ERROR_MSG_REGEX_MISS_MATCHED,
		}
		return ""
	}
}