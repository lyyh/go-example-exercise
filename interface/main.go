package main

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}
func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 10, height: 20}
	measure(r)
}
