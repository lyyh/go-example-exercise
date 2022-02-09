package main

import "fmt"

func main() {
	var s = make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	s1 := append(s, "1")
	fmt.Println(s1)

	var c = make([]string, 5)
	copy(c, s1)
	fmt.Println(c)

	towD := make([][]int, 5)
	for i := 0; i < 3; i++ {
		towD[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			towD[i][j] = i + j
		}
	}
	fmt.Println(towD)
}
