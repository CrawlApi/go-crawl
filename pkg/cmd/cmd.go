package cmd

import (
	"encoding/json"
	"github.com/huandu/facebook"
	"github.com/llitfkitfk/cirkol/pkg/client"
	"log"
	"time"
)

const (
	FACEBOOK = "facebook"
)

var (
	DBClient *client.Client
)

func GETPostInfoById(postId string, token string, userId string) string {
	facebook.Version = "v2.6"

	res, err1 := facebook.Get("/" + postId, facebook.Params{
		"fields":       "id,message,picture,full_picture,shares,updated_time,created_time,name,source,type",
		"access_token": token,
	})

	likes, err2 := facebook.Get("/" + postId + "/likes?summary=true", facebook.Params{
		"fields":       "",
		"access_token": token,
	})

	comments, err3 := facebook.Get("/" + postId + "/comments?summary=true", facebook.Params{
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

func startPush(pushCh chan string, ids []string, accessToken string, userId string) {
	for _, value := range ids {
		pushCh <- GETPostInfoById(value, accessToken, userId)
	}
}
func fetchData(data []string) {

	if len(data) > 1 {

		var tokenStr client.TokenStr
		err := json.Unmarshal([]byte(data[1]), &tokenStr)
		if err != nil {
			log.Println(err)
		}
		fbCh := make(chan string)
		go startPush(fbCh, tokenStr.Ids, tokenStr.AccessToken, tokenStr.Id)
		message := <-fbCh
		DBClient.PushData(FACEBOOK, message)
	}
}

func StartService() {

	for true {
		data, _ := DBClient.StartService()
		if len(data) > 0 {
			go fetchData(data)
		}
	}
}
