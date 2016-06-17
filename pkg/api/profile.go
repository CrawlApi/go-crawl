package api

import (
	"github.com/llitfkitfk/cirkol/pkg/result"
	"encoding/json"
)

func GetProfileApi(url string, ch chan <- result.Profile) string {
	body, errs := GetApi(url)
	if errs != nil {
		ch <- result.Profile{
			Website:    url,
			ErrCode:    ERROR_CODE_API_TIMEOUT,
			ErrMessage: ERROR_MSG_API_TIMEOUT,
		}
		return ""
	}
	return body
}

func ParseProfileJson(src string, v interface{}, ch chan <- result.Profile) {
	err := json.Unmarshal([]byte(src), v)
	if err != nil {
		ch <- result.Profile{
			RawData:    src,
			ErrCode:    ERROR_CODE_JSON_ERROR,
			ErrMessage: err.Error(),
		}
	}
}

func MatchStrProfileCh(length int, expr string, src string, ch chan <- result.Profile) string {

	des := MatchStrNoCh(length, expr, src)
	if des != "" {
		return des
	} else {
		ch <- result.Profile{
			ErrCode:    ERROR_CODE_REGEX_MISS_MATCHED,
			ErrMessage: ERROR_MSG_REGEX_MISS_MATCHED,
		}
		return ""
	}
}
