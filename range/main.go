package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	var sum int
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)

	kvs := map[string]string{"a": "hello", "b": "world"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
