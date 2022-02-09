package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 只读，只写通道
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	ch1 := make(chan string)
	go func() {
		ch1 <- "message"
	}()

	// 阻塞版的通道
	a := <-ch1
	fmt.Println(a)

	ch2 := make(chan string, 2)
	ch2 <- "bufferd"
	ch2 <- "close"
	for {
		select {
		case val := <-ch2:
			if val == "close" {
				close(ch2)
				return
			}
			fmt.Println(val)
		}
	}

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "aaa")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
