package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(nanos)
	fmt.Println(millis)
}
