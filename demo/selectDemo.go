package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)
	timer := time.Now()
	go func() {
		log.Println("4 ", time.Since(timer))
		messages <- "TEST"
	}()
	// Here's a non-blocking receive. If a value is
	// available on `messages` then `select` will take
	// the `<-messages` `case` with that value. If not
	// it will immediately take the `default` case.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	log.Println("1 ",time.Since(timer))
	// A non-blocking send works similarly.
	msg := "hi"
	time.Sleep(time.Second)

	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	log.Println("2 ", time.Since(timer))
	// We can use multiple `case`s above the `default`
	// clause to implement a multi-way non-blocking
	// select. Here we attempt non-blocking receives
	// on both `messages` and `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	log.Println("3 ",time.Since(timer))

	time.Sleep(time.Second)
}
