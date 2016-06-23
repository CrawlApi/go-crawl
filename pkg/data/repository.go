package data

import "github.com/parnurzeal/gorequest"

type Repo interface {
	GetAgent() *gorequest.SuperAgent

}
