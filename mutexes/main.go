package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu    sync.Mutex
	count map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count[name]++
}

func main() {
	c := Container{
		count: map[string]int{
			"a": 0,
			"b": 0,
		},
	}

	var wg sync.WaitGroup

	doInc := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(3)
	go doInc("a", 100)
	go doInc("b", 100)
	go doInc("c", 100)
	wg.Wait()
	fmt.Println(c.count)
}
