package main

import (
	"fmt"
	"time"
)

func f1(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(name, i)
	}
}
func main() {
	go f1("f1")
	go func(name string) {
		for i := 0; i < 3; i++ {
			fmt.Println(name, i)
		}
	}("f2")

	time.Sleep(time.Second)
}
