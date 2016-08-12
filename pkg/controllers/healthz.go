package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/common"
	"github.com/llitfkitfk/cirkol/pkg/data"
	"github.com/llitfkitfk/cirkol/pkg/models"
	"net/http"
	"sync"
)

func Healthz(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

type APICheck struct {
	Status map[string]bool `json:"status"`
	lock   sync.RWMutex
}

func (ac *APICheck) SetMap(key string, value bool) {
	ac.lock.Lock()
	defer ac.lock.Unlock()
	ac.Status[key] = value
}

func APIHealthz(c *gin.Context) {
	ch := &APICheck{
		Status: make(map[string]bool),
	}
	var wg sync.WaitGroup
	wg.Add(5)
	go checkingFB(ch, &wg)
	go checkingIG(ch, &wg)
	go checkingYTB(ch, &wg)
	go checkingWX(ch, &wg)
	go checkingWB(ch, &wg)

	wg.Wait()

	c.JSON(http.StatusOK, ch)
}

func checkingFB(ch *APICheck, wg *sync.WaitGroup) {
	wg.Add(2)
	go func() {
		r := data.GR().GetFBProfileResult("JustinTimberlake")
		var rawData models.FBRawProfile
		err := common.ParseJson(r.Body, &rawData)
		if err != nil {
			ch.SetMap("facebook_api_profile", false)
		}
		if rawData.ID == "" {
			ch.SetMap("facebook_api_profile", false)
		}
		ch.SetMap("facebook_api_profile", true)

		wg.Done()
	}()
	go func() {
		r := data.GR().GetFBPostsResult("JustinTimberlake", "5")
		var rawData models.FBRawPosts
		err := common.ParseJson(r.Body, &rawData)
		if err != nil {
			ch.SetMap("facebook_api_posts", false)
		}
		if len(rawData.Data) > 0 {
			ch.SetMap("facebook_api_posts", false)
		}
		ch.SetMap("facebook_api_posts", true)

		wg.Done()
	}()

	wg.Done()
}

func checkingIG(ch *APICheck, wg *sync.WaitGroup) {
	ch.SetMap("instagram_api", true)
	wg.Done()

}
func checkingYTB(ch *APICheck, wg *sync.WaitGroup) {
	ch.SetMap("youtube_api", true)
	wg.Done()

}
func checkingWX(ch *APICheck, wg *sync.WaitGroup) {
	ch.SetMap("wechat_api", true)
	wg.Done()
}
func checkingWB(ch *APICheck, wg *sync.WaitGroup) {
	ch.SetMap("weibo_api", true)
	wg.Done()
}
