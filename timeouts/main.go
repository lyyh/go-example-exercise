package main

import (
	"fmt"
	"time"
)

func main() {
	var ch1 = make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "result 1"
	}()

	select {
	case res := <-ch1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
}
