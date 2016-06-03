package main

import (
	"fmt"
	"log"
	"regexp"
	"time"
	"net/url"
)

var tokenCh chan string
var popchan chan string

func main() {
	//StartTokenGen()
	//URLTest()
	RegexpDemo()
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
