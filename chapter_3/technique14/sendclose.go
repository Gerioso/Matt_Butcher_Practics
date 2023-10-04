package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	msg := make(chan string)
	until := time.After(5 * time.Second)

	go send(msg, done)
	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true
			time.Sleep(500 * time.Millisecond)
			return
		}
	}

}

func send(msg chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			println("Done")
			close(msg)
			return
		default:
			msg <- "Hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
