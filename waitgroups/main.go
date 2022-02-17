package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(1000)
	fmt.Printf("worker %d done\n", id)
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// 避免协程闭包中重复利用相同的值
		j := i
		go func() {
			fmt.Printf("goroutine %d starting\n", j)
			worker(j)
			defer wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("main finished")
}
