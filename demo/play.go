package main

import (
	"fmt"
	"log"
	"regexp"
	"time"
	"net/url"
	"sync"
	"github.com/parnurzeal/gorequest"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"os"
	"os/signal"
)

var tokenCh chan string
var popchan chan string

type Book struct {
	Title       string `json:"title"`
	CoTitle     string `json:"co_title"`
	Cover       string `json:"cover"`
	QRCode      string `json:"qr_code"`
	Description string `json:"description"`
	Publisher   string `json:"publisher"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Year        string `json:"year"`
	Pages       string `json:"pages"`
	Language    string `json:"language"`
	FileSize    string `json:"file_size"`
	FileFormat  string `json:"file_format"`
	Download    string `json:"download"`
	Link        string `json:"link"`
}

func main() {
	//StartTokenGen()
	//URLTest()
	//RegexpDemo()
	//GoRoutine()
	//WaitGroupDemo()
	RequestDemo()
	//MongoQuery()

	//t, err := time.Parse("2006-01-02T15:04:05", "2016-02-01T12:01:03")
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(t)

}
func MongoQuery() {
	session, err := mgo.Dial("192.168.20.24:27000")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("Book")
	//ids := []int{3119, 3072, 3070, 3060}
	//for _, i := range ids {
	for i := 7106; i > 0; i-- {

		url := "http://it-ebooks.info/book/" + strconv.Itoa(i) + "/"
		count, err := c.Find(bson.M{"link": url}).Count()
		if err != nil {
			log.Fatal(err)
		}
		if count != 1 {
			log.Fatal(url)
		} else {
			log.Println(i)
		}

		//time.Sleep(200 * time.Millisecond)
		//	http://it-ebooks.info/book/3119/
		//	http://it-ebooks.info/book/3072/
		//	http://it-ebooks.info/book/3070/
		//http://it-ebooks.info/book/3060/
	}


	//fmt.Println("Phone:", result.)
}
func RequestDemo() {
	agent := gorequest.New()
	go func() {
		for {
			response, _, _ := agent.Get("http://localhost:10086/api/wb/profile/1732447702").End()
			log.Println(response.Body)
		}
	}()

	handleSignals()
}
func handleSignals() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
func URLTest() {
	rawurl := "http:/www.facebook.com/"
	realUrl, err := url.Parse(rawurl)
	if err != nil {
		log.Println("error: ", err)
	} else {
		log.Println("real url : ", realUrl.RequestURI())
	}
}

func WaitGroupDemo() {
	messages := make(chan int)
	var wg sync.WaitGroup

	// you can also add these one at
	// a time if you need to

	wg.Add(3)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 5)
		messages <- 5
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		messages <- 3
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 7)
		messages <- 7
	}()
	go func() {
		for i := range messages {
			fmt.Println(i)
		}
	}()
	wg.Wait()

}

func GoRoutine() {
	respon := make(chan string)
	go func() {
		respon <- "test1"
		time.Sleep(time.Second)
	}()

	go func() {
		respon <- "test2"
		time.Sleep(time.Second)
	}()
	go func() {
		respon <- "test2"
		time.Sleep(time.Second)
	}()
	go func() {
		respon <- "test1"
		time.Sleep(time.Second)
	}()
	for {
		select {
		case t := <-respon:
			log.Println(t)
		}

		time.Sleep(time.Second)
	}
}

func StartTokenGen() {
	tokenCh = make(chan string)
	popchan = make(chan string)

	go TokenGen(tokenCh)
	//go TokenGen1(popchan)

	for {
		select {
		case t := <-tokenCh:
			log.Println(t)
		case p := <-popchan:
			log.Println(p)
		default:
			log.Println("no message")
		}
	}

}

func TokenGen1(ch chan string) {
	for {
		popchan <- "pop"
		time.Sleep(1 * time.Second)

	}
}

func TokenGen(ch chan string) {
	for {
		log.Println("goroutine")
		time.Sleep(1 * time.Second)

	}
}
func RegexpDemo() {
	r, _ := regexp.Compile(`微信号: (\S+)<`)

	matcher := r.FindStringSubmatch("12412微信号: Bit_baike1241231")

	fmt.Println(matcher)
}
