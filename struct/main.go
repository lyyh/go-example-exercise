package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPersion(name string, age int) *person {
	return &person{name, age}
}
func main() {
	p1 := newPersion("lyh", 26)
	fmt.Println(p1)
	// 自动解引用
	fmt.Println(p1.name)
	s := person{name: "a", age: 10}
	fmt.Println(&s)
	fmt.Printf("%#v", s)
}
