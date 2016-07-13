package common

import (
	"errors"
	"fmt"
)

const (
	ERROR_MSG_API_TIMEOUT = "request api timeout"
	ERROR_MSG_EMPTY_REPO = "repo empty"
	ERROR_MSG_EMPTY_API = "request api empty"
	ERROR_MSG_PARSE_MODEL = "model parse error"
	ERROR_MSG_PARSE_UID = "parse uid error"

	ERROR_MSG_REGEX_MISS_MATCHED = "regex miss matched"
	ERROR_MSG_URL_NOT_MATCHED = "Url not matched"
)

func TimeOutError() error {
	return errors.New(ERROR_MSG_API_TIMEOUT)
}

func EmptyAPIError() error {
	return errors.New(ERROR_MSG_EMPTY_API)
}

func EmptyRepoError() error {
	return errors.New(ERROR_MSG_EMPTY_REPO)
}

func MissMatchError() error {
	return errors.New(ERROR_MSG_REGEX_MISS_MATCHED)
}

func ParseModelError() error {
	return errors.New(ERROR_MSG_PARSE_MODEL)
}

func Errs2Error(errs []error) error {
	return fmt.Errorf(`%s`, errs)
}

func ParseUIDError()  error {
	return errors.New(ERROR_MSG_PARSE_UID )
}