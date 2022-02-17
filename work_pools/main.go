package main

// 利用协程和通道实现一个工作池 worker pool
// 开3个协程去执行5个job通道
import (
	"fmt"
	"time"
)

func worker(id int, job <-chan int, results chan<- int) {
	for j := range job {
		fmt.Println("worker", id, "start job", j)
		time.Sleep(1000)
		fmt.Println("worker", id, "finish job", j)
		results <- j * 2
	}
}

func main() {
	jobNum := 5
	job := make(chan int, jobNum)
	results := make(chan int, jobNum)
	// 启动三个协程
	for i := 0; i < 3; i++ {
		go worker(i, job, results)
	}

	// 发送五个job通道
	for j := 0; j < jobNum; j++ {
		job <- j
	}
	close(job)
	// 收集任务的返回值
	for k := 0; k < jobNum; k++ {
		<-results
	}
}
