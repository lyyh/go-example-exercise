package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2
	fmt.Println(m)

	delete(m, "k1")
	_, prs := m["k2"] // 键不存在或者键为零值的歧义
	fmt.Println(prs)
}
