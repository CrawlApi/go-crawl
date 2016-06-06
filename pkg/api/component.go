package api

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/parnurzeal/gorequest"
	"log"
	"github.com/gin-gonic/gin"
)

var (
	logCh chan interface{}
	reqClient  *gorequest.SuperAgent
)

func SetupComponent(router *gin.Engine) {
	setSession(router)
	setupRequestClient()
	setupLogger()
}

func setSession(router *gin.Engine) {
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("sessionManage", store))
	//router.Use(SessionManage())
}

func SessionManage() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get("token")
		if v == nil {
			v = FACEBOOK_TOKEN
		}
		session.Set("token", v)
		session.Save()
		log.Println("Session Save")
		c.Next()
	}
}

func setupRequestClient() {

	reqClient = gorequest.New()
}

func setupLogger() {
	logCh = make(chan interface{}, 10)
	go logging()
}

func logging() {
	for {
		log.Println(<-logCh)
	}
}