package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
		time.Sleep(time.Second * 1)
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
		time.Sleep(time.Second * 2)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go ping(c1)
	go pong(c2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c2:
			fmt.Println(msg1)
		case msg2 := <-c1:
			fmt.Println(msg2)
		}
	}

}
