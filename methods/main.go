package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	// 调用方法时会自动处理值和指针之间的转换
	rt := &rect{width: 10, height: 20}
	fmt.Println(rt.area())
	rt.width = 20
	fmt.Println(rt.perim())
}
