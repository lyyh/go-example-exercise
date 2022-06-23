package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (b byLength) Len() int {
	return len(b)
}
func (b byLength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// 将控制实际的自定义排序逻辑
func (b byLength) Less(i, j int) bool {
	return len(b[i]) < len(b[j])
}
func sortByFunction() {
	// 原地排序
	s := []string{"ab", "a", "abc"}
	sort.Sort(byLength(s))
	fmt.Println(s)
}

func main() {
	// list := []string{"b", "a"}
	// sort.Strings(list)
	// fmt.Println(list)

	// ints := []int{1, 2, 4, 3, 1}
	// sort.Ints(ints)
	// fmt.Println(ints)

	// ret := sort.IntsAreSorted(ints)
	// fmt.Println(ret)
	sortByFunction()
}
