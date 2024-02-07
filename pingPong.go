package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		time.Sleep(time.Second)
		c <- "ping"
	}
}
func pong(c chan string) {
	for {
		time.Sleep(time.Second)
		c <- "pong"
	}
}
func main() {
	c := make(chan string)

	go ping(c)
	go pong(c)

	for {
		fmt.Println(<-c)
	}
}
