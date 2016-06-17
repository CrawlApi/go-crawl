package api

import "github.com/llitfkitfk/cirkol/pkg/result"


func MatchStrUidCh(length int, expr string, src string, ch chan<- result.UID) string {

	des := MatchStrNoCh(length, expr, src)
	if des != "" {
		return des
	} else {
		ch <- result.UID{
			ErrCode:    ERROR_CODE_REGEX_MISS_MATCHED,
			ErrMessage: ERROR_MSG_REGEX_MISS_MATCHED,
		}
		return ""
	}
}

func GetUidApi(url string, ch chan <- result.UID) string {
	body, errs := GetApi(url)
	if errs != nil {
		ch <- result.UID{
			ErrCode:    ERROR_CODE_API_TIMEOUT,
			ErrMessage: ERROR_MSG_API_TIMEOUT,
		}
		return ""
	}
	return body
}
