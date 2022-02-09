package main

import "fmt"

type base struct {
	num int
}

type container struct {
	base
	str string
}

type describer interface {
	describe() int
}

func (b base) describe() int {
	fmt.Printf("%v", b.num)
	return b.num
}

// 嵌入结构可以给被嵌入接口带来接口实现、属性
func main() {
	// 当创建含有嵌入的结构体，必须显示的初始化
	co := container{
		base: base{
			num: 10,
		},
		str: "hello",
	}

	fmt.Println(co.num)
	fmt.Println(co.base.num)
	//因为container嵌入了base,直接从co上调用了一个从base嵌入的方法
	fmt.Println(co.describe())
	var d describer = co
	fmt.Println(d.describe())
}
