package main

import "fmt"

func main() {
	var arr1 [5]int // 初始化为类型的零值
	fmt.Println(arr1)

	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
