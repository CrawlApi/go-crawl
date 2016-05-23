package cmd

import (
	"encoding/json"
	"github.com/huandu/facebook"
	"github.com/llitfkitfk/cirkol/pkg/client"
	"github.com/parnurzeal/gorequest"
	"log"
	"regexp"
	"time"
	"github.com/llitfkitfk/cirkol/pkg/api"
)

const (
	FACEBOOK           = "facebook"
	REGEXP_FACEBOOK_ID = `fb://(page|profile)/(\d+)`
)

var (
	Client  *client.Client
	TokenCh chan string
	CommCh  chan string
)

func GETPostInfoById(postId string, token string, userId string) string {
	facebook.Version = "v2.6"

	res, err1 := facebook.Get("/"+postId, facebook.Params{
		"fields":       "id,message,picture,full_picture,shares,updated_time,created_time,name,source,type",
		"access_token": token,
	})

	likes, err2 := facebook.Get("/"+postId+"/likes?summary=true", facebook.Params{
		"fields":       "",
		"access_token": token,
	})

	comments, err3 := facebook.Get("/"+postId+"/comments?summary=true", facebook.Params{
		"fields":       "",
		"access_token": token,
	})

	var result client.FaceBook

	if err1 != nil {
		result.Message = err1.Error()
		result.Status = false
		result.Date = time.Now().Unix()
	} else if err2 != nil {
		result.Message = err2.Error()
		result.Status = false
		result.Date = time.Now().Unix()
	} else if err3 != nil {
		result.Message = err3.Error()
		result.Status = false
		result.Date = time.Now().Unix()
	} else {
		result.Message = "OK"
		result.Status = true
		result.Date = time.Now().Unix()
		result.Content = res["message"]
		result.Picture = res["picture"]
		result.Full_picture = res["full_picture"]
		result.Shares = res["shares"]
		result.Name = res["name"]
		result.UpdatedTime = res["updated_time"]
		result.CreatedTime = res["created_time"]
		result.Id = res["id"]
		result.Source = res["source"]
		result.Type = res["type"]
		result.UserId = userId

		result.Likes = likes["summary"]
		result.Comments = comments["summary"]

	}

	postResult := &client.FaceBookPost{
		Post: &result,
	}

	out, _ := json.Marshal(postResult)

	return string(out)
}

func startPush(pushCh chan string, accessToken string, userId string) {
	//pushCh <- GETPostInfoById(accessToken, userId)

}

func fetchData(data []string) {

	if len(data) > 1 {
		var token client.TokenStr
		err := json.Unmarshal([]byte(data[1]), &token)
		if err != nil {
			log.Println(err)
		}

		fbCh := make(chan string)
		fbIdCh := make(chan string)

		request := gorequest.New()

		go FetchIdFromUrl(fbIdCh, token.Url, request)

		for id := range fbIdCh {
			log.Println(id)
			//go startPush(fbCh, token.Ids, token.AccessToken, id)
		}

		for message := range fbCh {
			Client.PushData(FACEBOOK, message)
		}

		close(fbCh)
	}
}

func FetchIdFromUrl(idCh chan string, url string, request *gorequest.SuperAgent) {

	_, body, errs := request.Get(url).End()

	if errs != nil {
		log.Println(errs)
	}
	r, _ := regexp.Compile(REGEXP_FACEBOOK_ID)

	matcher := r.FindStringSubmatch(body)

	if len(matcher) > 2 {
		idCh <- matcher[2]
	}
}

func generateToken(tag string) {
	switch tag {
	case FACEBOOK:
		TokenCh <- api.FACEBOOK_TOKEN
	}
}

func StartTokenGen() {
	for {
		select {
		case t := <-CommCh:
			go generateToken(t)
		}
	}
}

func StartService() {
	go StartTokenGen()
	for {
		data, _ := Client.StartRedis()

		if len(data) > 0 {
			go fetchData(data)
		}
	}
}
