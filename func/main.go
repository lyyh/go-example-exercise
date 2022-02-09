package main

import "fmt"

func plus(a, b, c int) int {
	return a + b + c
}

func sum(nums ...int) (total int) {
	for _, n := range nums {
		total += n
	}
	return
}

func main() {
	ret := plus(1, 2, 3)

	fmt.Println(ret)

	sli := []int{1, 2, 3}
	fmt.Println(sum(sli...))
}
